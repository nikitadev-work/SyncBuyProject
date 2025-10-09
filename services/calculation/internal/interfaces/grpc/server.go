package interfaces

import (
	"context"

	uc "calculation/internal/usecase"
	pb "calculation/proto-codegen/calculation"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CalculationServer struct {
	pb.UnimplementedCalculationServiceServer
	usecase uc.Usecase
}

func (s *CalculationServer) CalculateDistribution(ctx context.Context, in *pb.CalcRequest) (*pb.CalcResponse, error) {
	err := ValidateDistributionRequest(in)
	if err != nil {
		return nil, ErrToStatus(ErrInvalidInputData)
	}

	reqDTO, err := ConvertCalcReqDataToInternalDTO(in)
	if err != nil {
		return nil, ErrToStatus(err)
	}

	calcResult, err := s.usecase.CalculateDistribution(ctx, *reqDTO)
	if err != nil {
		return nil, ErrToStatus(err)
	}

	result, err := ConvertCalcFromInternalDTOtoResp(&calcResult)
	if err != nil {
		return nil, ErrToStatus(err)
	}

	return result, nil
}

func (s *CalculationServer) Health(ctx context.Context, _ *emptypb.Empty) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{Status: "OK"}, nil
}
