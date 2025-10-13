package app

import (
	"calculation/config"
	"calculation/internal/infra/logger"
	grpcserver "calculation/internal/interfaces/grpc"
	httpserver "calculation/internal/interfaces/http"
	"calculation/internal/usecase"
	calcpb "calculation/proto-codegen/calculation"
	"context"
	"errors"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, cfg *config.Config) error {
	l := logger.NewLogger(
		cfg.Log.Level,
		map[string]any{
			"service": cfg.App.Name,
			"version": cfg.App.Version,
		},
	)

	l.Info("start configuration", nil)

	//grpc
	uc := usecase.NewUsecase()
	srv := grpc.NewServer()
	handler := grpcserver.New(uc, l)
	calcpb.RegisterCalculationServiceServer(srv, handler)

	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		return err
	}

	grpcErrCh := make(chan error, 1)
	go func() {
		l.Info("start grpc server", nil)
		err := srv.Serve(lis)
		if err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			grpcErrCh <- err
		}
	}()

	//http
	httpMux := http.NewServeMux()
	httpServer := httpserver.New(l, httpMux, cfg.HTTP.Port)
	httpMux.Handle("/metrics", promhttp.Handler())

	httpErrCh := make(chan error, 1)
	go func() {
		l.Info("start http server", nil)
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			httpErrCh <- err
		}
	}()

	l.Info("calculation service started", map[string]any{
		"grpc.port": cfg.GRPC.Port,
		"http.port": cfg.HTTP.Port,
		"log.level": cfg.Log.Level,
	})

	select {
	case <-ctx.Done():
		l.Info("starting graceful shutdown", nil)

		//gracefull shutdown
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		done := make(chan struct{})

		go func() {
			var wg sync.WaitGroup
			wg.Add(2)

			//grpc server
			go func() {
				defer wg.Done()
				srv.GracefulStop()
			}()

			//http server
			go func() {
				defer wg.Done()

				err := httpServer.Shutdown(shutdownCtx)
				if err != nil {
					l.Error("http server graceful shutdown error", map[string]any{
						"error": err.Error(),
					})
				}
			}()

			wg.Wait()
			close(done)
		}()

		select {
		case <-done:
			//successfully finished
			l.Info("gracefully finished", nil)
			return nil
		case <-shutdownCtx.Done():
			srv.Stop()
			httpServer.Close()

			err := errors.New("graceful shutdown timeout")
			l.Error("graceful shutdown error", map[string]any{
				"error": err.Error(),
			})
			return err
		}
	case err := <-grpcErrCh:
		l.Error("grpc server error", map[string]any{
			"error": err.Error(),
		})
		return err
	case err := <-httpErrCh:
		l.Error("http server error", map[string]any{
			"error": err.Error(),
		})
		return err
	}
}
