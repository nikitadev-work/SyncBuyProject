package usecase

import "context"

type PurchaseRepositoryInterface interface {
	// Purchase
	CreatePurchase(CreatePurchaseRequestDTO) error
	GetPurchase(GetPurchaseRequestDTO) *GetPurchaseResponseDTO
	
	// Participant
	AddParticipantToPurchase(AddParticipantToPurchaseRequestDTO) error
	RemoveParticipant(RemoveParticipantRequestDTO) error
	ListParticipantsByPurchaseId(ListParticipantsByPurchaseIdRequestDTO) *ListParticipantsByPurchaseIdResponseDTO

	// Task
	CreateTask(CreatePurchaseRequestDTO) error
	TakeTask(TakeTaskRequestDTO) error
	DeleteTask(DeleteTaskRequestDTO) error
	ListTasksByPurchaseId(ListParticipantsByPurchaseIdRequestDTO) *ListParticipantsByPurchaseIdResponseDTO
	MarkTaskAsDone(MarkTaskAsDoneRequestDTO) error

	// Status
	LockPurchase(LockPurchaseRequestDTO) error
	UnlockPurchase(UnlockPurchaseRequestDTO) error
	GetSnapshot(GetSnapshotRequestDTO) *GetPurchaseResponseDTO
	MarkSettlementInitiated(MarkSettlementInitiatedRequestDTO) error
	FinishPurchase(FinishPurchaseRequestDTO) error
}

type TxManagerInterface interface {
	WithinTx(context.Context, func(context.Context) error) error
}
