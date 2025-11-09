package usecase

import (
	"purchase/internal/domain"

	"github.com/gofrs/uuid"
)

type CreatePurchaseInputDTO struct {
	Title       string
	Description string
	Currency    string
}

type CreatePurchaseOutputDTO struct {
	PurchaseId uuid.UUID
}

type GetPurchaseInputDTO struct {
	PurchaseId uuid.UUID
}

type GetPurchaseOutputDTO struct {
	Purchase domain.Purchase
}

type CreateInviteInputDTO struct {
	PurchaseId uuid.UUID
}

type CreateInviteOutputDTO struct {
	Token       string
	Title       string
	Description string
}

type JoinByInviteInputDTO struct {
	UserId uuid.UUID
	Token  string
}

type RemoveParticipantInputDTO struct {
	UserId     uuid.UUID
	PurchaseId uuid.UUID
}

type ListParticipantsByPurchaseIdInputDTO struct {
	PurchaseId uuid.UUID
}

type ListParticipantsByPurchaseIdOutputDTO struct {
	UserIds []uuid.UUID
}

type CreateTaskInputDTO struct {
	Title        string
	Description  string
	PurchaseId   uuid.UUID
	AuthorUserId uuid.UUID
	Amount       int64
}

type CreateTaskOutputDTO struct {
	TaskId uuid.UUID
}

type TakeTaskInputDTO struct {
	TaskId         uuid.UUID
	ExecutorUserId uuid.UUID
}

type DeleteTaskInputDTO struct {
	TaskId uuid.UUID
}

type ListTasksByPurchaseIdInputDTO struct {
	PurchaseId uuid.UUID
}

type ListTasksByPurchaseIdOutputDTO struct {
	Tasks []domain.Task
}

type MarkTaskAsDoneInputDTO struct {
	TaskId uuid.UUID
}

type LockPurchaseInputDTO struct {
	PurchaseId uuid.UUID
}

type UnlockPurchaseInputDTO struct {
	PurchaseId uuid.UUID
}

type GetSnapshotInputDTO struct {
	PurchaseId uuid.UUID
}

type GetSnapshotOutputDTO struct {
	Snapshot domain.Snapstot
}

type MarkSettlementInitiatedInputDTO struct {
	PurchaseId uuid.UUID
}

type FinishPurchaseInputDTO struct {
	PurchaseId uuid.UUID
}
