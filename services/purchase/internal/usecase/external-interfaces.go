package usecase

import "context"

type PurchaseRepositoryInterface interface {
	// Purchase
	CreatePurchase(context.Context, *CreatePurchaseRequestDTO) (*CreatePurchaseReponseDTO, error)
	GetPurchase(context.Context, *GetPurchaseRequestDTO) (*GetPurchaseResponseDTO, error)

	// Participant
	AddParticipantToPurchase(context.Context, *AddParticipantToPurchaseRequestDTO) error
	RemoveParticipant(context.Context, *RemoveParticipantRequestDTO) error
	ListParticipantsByPurchaseId(context.Context, *ListParticipantsByPurchaseIdRequestDTO) (*ListParticipantsByPurchaseIdResponseDTO, error)

	// Task
	CreateTask(context.Context, *CreateTaskRequestDTO) (*CreateTaskResponseDTO, error)
	TakeTask(context.Context, *TakeTaskRequestDTO) error
	DeleteTask(context.Context, *DeleteTaskRequestDTO) error
	ListTasksByPurchaseId(context.Context, *ListTasksByPurchaseIdRequestDTO) (*ListTasksByPurchaseIdResponseDTO, error)
	MarkTaskAsDone(context.Context, *MarkTaskAsDoneRequestDTO) error

	// Status
	LockPurchase(context.Context, *LockPurchaseRequestDTO) error
	UnlockPurchase(context.Context, *UnlockPurchaseRequestDTO) error
	MarkSettlementInitiated(context.Context, *MarkSettlementInitiatedRequestDTO) error
	FinishPurchase(context.Context, *FinishPurchaseRequestDTO) error
}

type TxManagerInterface interface {
	WithinTx(context.Context, func(context.Context) error) error
}
