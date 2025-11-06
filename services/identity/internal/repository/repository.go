package repository

import (
	"context"
	uc "identity/internal/usecase"
)

type Repository struct {
}

func NewRepository() *Repository {
	return &Repository{}
}

// TODO
func (repo *Repository) CreateNewUser(ctx context.Context, input *uc.CreateNewUserRequestDTO) (*uc.CreateNewUserResponseDTO, error) {
	return nil, nil
}

func (repo *Repository) GetUserProfileByUserId(ctx context.Context, input *uc.GetUserProfileByUserIdRequestDTO) (*uc.GetUserProfileByUserIdResponseDTO, error) {
	return nil, nil
}

func (repo *Repository) CreateNewIdentity(ctx context.Context, input *uc.CreateNewIdentityRequestDTO) error {
	return nil
}

func (repo *Repository) GetIdentityByTelegramId(ctx context.Context, input *uc.GetIdentityByTelegramIdRequestDTO) (*uc.GetIdentityByTelegramIdResponseDTO, error) {
	return nil, nil
}

func (repo *Repository) GetIdentityByUserId(ctx context.Context, input *uc.GetIdentityByUserIdRequestDTO) (*uc.GetIdentityByUserIdResponseDTO, error) {
	return nil, nil
}
