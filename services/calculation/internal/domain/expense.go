package domain

import "github.com/google/uuid"

type Expense struct {
	Id       uuid.UUID
	AuthorId uuid.UUID
	Money    Money
}

func NewExpense(id uuid.UUID, authorId uuid.UUID, money Money) *Expense {
	return &Expense{
		Id:       id,
		AuthorId: authorId,
		Money:    money,
	}
}
