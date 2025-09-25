package domain

import "github.com/google/uuid"

type Expense struct {
	Id       uuid.UUID
	AuthorId uuid.UUID
	Money    Money
}
