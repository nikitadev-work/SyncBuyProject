package usecase

import (
	"calculation/internal/domain"
	"context"
)

type CalculationUsecase interface {
	CalculateDistribution(ctx context.Context, input CalculationInputDTO) (CalculationOutputDTO, error)
}

type Usecase struct {
}

var _ CalculationUsecase = (*Usecase)(nil)

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (uc *Usecase) CalculateDistribution(ctx context.Context, input CalculationInputDTO) (CalculationOutputDTO, error) {
	res, err := domain.CalculateIntents(input.Participants, input.Expenses)
	if err != nil {
		return CalculationOutputDTO{}, err
	}

	return CalculationOutputDTO{Intents: res}, nil
}
