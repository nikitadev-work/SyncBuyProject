package usecase

import (
	"encoding/json"
	"identity/internal/domain"

	"github.com/google/uuid"
)

type RegisterOrGetUserByTelegramInputDTO struct {
	TelegramId string
	FirstName  string
	LastName   string
	ChatId     string
	Meta       json.RawMessage
}

type RegisterOrGetUserByTelegramOutputDTO struct {
	UserId uuid.UUID
}

type GetUserByTelegramInputDTO struct {
	TelegramId string
}

type GetUserByTelegramOutputDTO struct {
	User domain.User
}

type GetUserByUserIdInputDTO struct {
	UserId uuid.UUID
}

type GetUserByUserIdOutputDTO struct {
	User domain.User
}
