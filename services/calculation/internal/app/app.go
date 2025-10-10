package app

import (
	"calculation/config"
	"calculation/internal/infra/logger"
	grpcserver "calculation/internal/interfaces/grpc"
	"calculation/internal/usecase"
	calcpb "calculation/proto-codegen/calculation"
	"context"
	"errors"
	"net"
	"time"

	"google.golang.org/grpc"
)

func Run(ctx context.Context, cfg *config.Config) error {
	l := logger.NewLogger(cfg.Log.Level)

	uc := usecase.NewUsecase()

	srv := grpc.NewServer()
	handler := grpcserver.New(uc, l)
	calcpb.RegisterCalculationServiceServer(srv, handler)

	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		return err
	}

	errCh := make(chan error, 1)
	go func() {
		errCh <- srv.Serve(lis)
	}()

	select {
	case <-ctx.Done():
		//gracefull shutdown
		done := make(chan struct{})
		go func() {
			srv.GracefulStop()
			close(done)
		}()

		select {
		case <-done:
			//successfully finished
			return nil
		case <-time.After(10 * time.Second):
			srv.Stop()
			return errors.New("graceful shutdown timeout")
		}
	case err := <-errCh:
		return err
	}

}
