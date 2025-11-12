package domain

import (
	"time"

	"github.com/google/uuid"
)

type Status int

const (
	Active Status = iota
	Locked
	Finished
)

type Purchase struct {
	Id                    uuid.UUID
	Title                 string
	Description           string
	Currency              string
	SettlementInitiatedAt *time.Time
	Status                Status
	LockedAt              *time.Time
	FinishedAt            *time.Time
}

func NewPurchase(id uuid.UUID, title string, description string,
	currency string, SettlementInitiatedAt time.Time, status Status,
	lockedAt time.Time, finishedAt time.Time) *Purchase {
	return &Purchase{
		Id:                    id,
		Title:                 title,
		Description:           description,
		Currency:              currency,
		SettlementInitiatedAt: &SettlementInitiatedAt,
		Status:                status,
		LockedAt:              &lockedAt,
		FinishedAt:            &finishedAt,
	}
}

type Snapshot struct {
	PurchaseId         uuid.UUID
	PurchaseTitle      string
	Currency           string
	Tasks              []Task
	ParticipantUserIds []uuid.UUID
}
