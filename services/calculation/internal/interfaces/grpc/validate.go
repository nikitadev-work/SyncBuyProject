package grpcserver

import (
	"calculation/internal/domain"
	pb "calculation/proto-codegen/calculation"

	"github.com/google/uuid"
)

func ValidateDistributionRequest(input *pb.CalcRequest) error {
	if len(input.Participants) <= 0 {
		return ErrEmptyParticipantsList
	}

	if len(input.Expenses) <= 0 {
		return ErrEmptyExpensesList
	}

	_, err := uuid.Parse(input.PurchaseId)
	if err != nil {
		return err
	}

	for _, v := range input.Participants {
		_, err := uuid.Parse(v.UserId)
		if err != nil {
			return err
		}
	}

	for _, v := range input.Expenses {
		_, err := uuid.Parse(v.AuthorId)
		if err != nil {
			return err
		}

		_, err = uuid.Parse(v.Id)
		if err != nil {
			return err
		}

		if v.Cost.Amount < 0 {
			return ErrNegativeAmount
		}

		if !domain.IsSupportedCurrencyFormat(v.Cost.Currency) {
			return ErrIncorrectCurrencyFormat
		}
	}

	return nil
}
