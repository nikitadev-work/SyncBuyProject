package usecase

import "calculation/internal/domain"

type CalculationInputDTO struct {
	Participants []domain.Participant
	Expenses     []domain.Expense
}

type CalculationOutputDTO struct {
	Intents []domain.Intent
}
