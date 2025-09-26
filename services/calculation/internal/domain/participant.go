package domain

import "github.com/google/uuid"

type Participant struct {
	UserId uuid.UUID
}

func NewParticipant(userID uuid.UUID) *Participant {
	return &Participant{UserId: userID}
}
