package domain

import (
	"github.com/gofrs/uuid"
)

type Task struct {
	Id             uuid.UUID
	Title          string
	Description    string
	PurchaseId     uuid.UUID
	AuthorUserId   uuid.UUID
	ExecutorUserId uuid.UUID
	Done           bool
	Amount         int64
}

func NewTask(id uuid.UUID, title string, description string,
	purchaseId uuid.UUID, authorUserId uuid.UUID, executorUserId uuid.UUID,
	done bool, amount int64) *Task {
	return &Task{
		Id:             id,
		Title:          title,
		Description:    description,
		PurchaseId:     purchaseId,
		AuthorUserId:   authorUserId,
		ExecutorUserId: executorUserId,
		Done:           done,
		Amount:         amount,
	}
}
