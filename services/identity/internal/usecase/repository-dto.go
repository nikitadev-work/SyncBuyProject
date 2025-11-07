package usecase

import (
	"encoding/json"
	"identity/internal/domain"

	"github.com/google/uuid"
)

type CreateNewUserRequestDTO struct {
	FirstName string
	LastName  string
	Status    domain.Status
}

type CreateNewUserResponseDTO struct {
	UserId uuid.UUID
}

type GetUserProfileByUserIdRequestDTO struct {
	UserId uuid.UUID
}

type GetUserProfileByUserIdResponseDTO struct {
	User domain.User
}

type CreateNewIdentityRequestDTO struct {
	ExternalId   string
	InternalId   uuid.UUID
	ProviderType domain.ProviderType
	ChatId       string
	Meta         json.RawMessage
}

type GetIdentityByTelegramIdRequestDTO struct {
	TelegramId string
}

type GetIdentityByTelegramIdResponseDTO struct {
	Identity domain.Identity
}

type GetIdentityByUserIdRequestDTO struct {
	UserId uuid.UUID
}

type GetIdentityByUserIdResponseDTO struct {
	Identity domain.Identity
}
