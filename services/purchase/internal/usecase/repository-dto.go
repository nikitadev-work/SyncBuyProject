package usecase

import (
	"purchase/internal/domain"

	"github.com/gofrs/uuid"
)

type CreatePurchaseRequestDTO struct {
	Title       string
	Description string
	Currency    string
}

type GetPurchaseRequestDTO struct {
	PurchaseId uuid.UUID
}

type GetPurchaseResponseDTO struct {
	Purchase domain.Purchase
}

type AddParticipantToPurchaseRequestDTO struct {
	UserId     uuid.UUID
	PurchaseId uuid.UUID
}

type RemoveParticipantRequestDTO struct {
	UserId     uuid.UUID
	PurchaseId uuid.UUID
}

type ListParticipantsByPurchaseIdRequestDTO struct {
	PurchaseId uuid.UUID
}

type ListParticipantsByPurchaseIdResponseDTO struct {
	UserIds []uuid.UUID
}

type CreateTaskRequestDTO struct {
	Title          string
	Description    string
	PurchaseId     uuid.UUID
	AuthorUserId   uuid.UUID
	ExecutorUserId uuid.UUID
	Amount         int64
}

type TakeTaskRequestDTO struct {
	TaskId uuid.UUID
}

type DeleteTaskRequestDTO struct {
	TaskId uuid.UUID
}

type ListTasksByPurchaseIdRequestDTO struct {
	PurchaseId uuid.UUID
}

type ListTasksByPurchaseIdResponseDTO struct {
	Tasks []domain.Task
}

type MarkTaskAsDoneRequestDTO struct {
	TaskId uuid.UUID
}

type LockPurchaseRequestDTO struct {
	PurchaseId uuid.UUID
}

type UnlockPurchaseRequestDTO struct {
	PurchaseId uuid.UUID
}

type GetSnapshotRequestDTO struct {
	PurchaseId uuid.UUID
}

type GetSnapshotResponseDTO struct {
	Snapshot domain.Snapstot
}

type MarkSettlementInitiatedRequestDTO struct {
	PurchaseId uuid.UUID
}

type FinishPurchaseRequestDTO struct {
	PurchaseId uuid.UUID
}
