package domain

import (
	"github.com/google/uuid"
)

type Intent struct {
	PayerId uuid.UUID
	PayeeId uuid.UUID
	Money   Money
}

func NewIntent(payerId uuid.UUID, payeeId uuid.UUID, money Money) (Intent, error) {
	if payeeId == payerId {
		return Intent{}, ErrSelfPayment
	}

	if money.Amount <= 0 {
		return Intent{}, ErrNotPositiveIntentAmount
	}

	return Intent{
		PayerId: payerId,
		PayeeId: payeeId,
		Money:   money,
	}, nil
}
