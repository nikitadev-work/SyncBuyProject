package grpcserver

import (
	"context"
	"time"

	"calculation/internal/infra/logger"
	uc "calculation/internal/usecase"
	pb "calculation/proto-codegen/calculation"

	"calculation/internal/infra/metrics"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CalculationServer struct {
	pb.UnimplementedCalculationServiceServer
	usecase uc.CalculationUsecase

	logger logger.LoggerInterface
}

func New(usecase uc.CalculationUsecase, logger logger.LoggerInterface) *CalculationServer {
	return &CalculationServer{
		usecase: usecase,
		logger:  logger,
	}
}

func (s *CalculationServer) CalculateDistribution(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResponse, error) {
	start := time.Now()
	defer func() {
		metrics.IncGRPCRequestsTotal("CalculateDistribution")
	}()

	s.logger.Info("calc request started", map[string]any{
		"method":       "CalculateDistribution",
		"participants": len(in.Participants),
		"expenses":     len(in.Expenses),
		"purchase_id":  in.PurchaseId,
	})

	err := ValidateDistributionRequest(in)
	if err != nil {
		s.logger.Error("validation error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(ErrInvalidInputData)
	}

	reqDTO, err := ConvertCalcReqDataToInternalDTO(in)
	if err != nil {
		s.logger.Error("convertation error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	calcResult, err := s.usecase.CalculateDistribution(ctx, *reqDTO)
	if err != nil {
		s.logger.Error("calculation error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	result, err := ConvertCalcFromInternalDTOtoResp(&calcResult)
	if err != nil {
		s.logger.Error("convertation error", map[string]any{
			"error": err.Error(),
		})
		return nil, ErrToStatus(err)
	}

	s.logger.Info("finished calculation", map[string]any{
		"intents":     len(result.Intents),
		"duration_ms": time.Since(start).Milliseconds(),
	})

	return result, nil
}

func (s *CalculationServer) Health(ctx context.Context, _ *emptypb.Empty) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{Status: "OK"}, nil
}
