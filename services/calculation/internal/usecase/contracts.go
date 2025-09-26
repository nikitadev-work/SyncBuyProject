package usecase

import (
	"context"
)

type CalculationUsecase interface {
	CalculateDistribution(ctx context.Context, input CalculationInputDTO) (CalculationOutputDTO, error)
}
