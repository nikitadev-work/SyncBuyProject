package app

import (
	"context"
	"errors"
	"fmt"
	"identity/config"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	grpcserver "identity/internal/adapters/grpc"
	httpserver "identity/internal/adapters/http"
	txmanager "identity/internal/adapters/txmanager"
	repo "identity/internal/repository"
	uc "identity/internal/usecase"
	calcpb "identity/proto-codegen"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nikitadev-work/SyncBuyProject/common/kit/logger"
	"github.com/nikitadev-work/SyncBuyProject/common/kit/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, cfg *config.Config) error {
	// logger
	l := logger.NewLogger(
		cfg.Log.Level,
		map[string]any{
			"service": cfg.App.Name,
			"version": cfg.App.Version,
		},
	)

	l.Info("start configuration", nil)

	// metrics
	if cfg.Metrics.Enabled == true {
		metrics.InitMetrics()
	}

	// postgresql
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.PostgreSQL.User, cfg.PostgreSQL.Password, cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port, cfg.PostgreSQL.Name, strconv.FormatBool(cfg.PostgreSQL.SslEnabled))
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		l.Error("unable to create connection pool: %v\n", map[string]any{
			"error": err.Error(),
		})
		return err
	}
	defer pool.Close()

	// repository
	repository := repo.NewRepository(pool, cfg.PostgreSQL.TxMarker)

	// txmanager
	txManager := txmanager.NewTxManager(pool, l, cfg.PostgreSQL.TxMarker)

	// usecase
	usecase := uc.NewIdentityUsecase(repository, txManager)

	// grpc
	srv := grpc.NewServer()
	handler := grpcserver.New(*usecase, l)
	calcpb.RegisterIdentityServiceServer(srv, handler)

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

	// http
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

	l.Info("identity service started", map[string]any{
		"grpc.port": cfg.GRPC.Port,
		"http.port": cfg.HTTP.Port,
		"log.level": cfg.Log.Level,
		"db.name":   cfg.PostgreSQL.Name,
		"db.host":   cfg.PostgreSQL.Host,
		"db.port":   cfg.PostgreSQL.Port,
	})

	// gracefull shutdown
	select {
	case <-ctx.Done():
		l.Info("starting graceful shutdown", nil)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		done := make(chan struct{})

		go func() {
			var wg sync.WaitGroup
			wg.Add(2)

			// grpc server
			go func() {
				defer wg.Done()
				srv.GracefulStop()
			}()

			// http server
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
			// successfully finished
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
