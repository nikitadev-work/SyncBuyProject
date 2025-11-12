package domain

import (
	"time"

	"github.com/google/uuid"
)

type Participant struct {
	UserId     uuid.UUID
	PurchaseId uuid.UUID
	JoinedAt   time.Time
}

func NewParticipant(userId uuid.UUID, purchaseId uuid.UUID, joinedAt time.Time) *Participant {
	return &Participant{
		UserId:     userId,
		PurchaseId: purchaseId,
		JoinedAt:   joinedAt,
	}
}
