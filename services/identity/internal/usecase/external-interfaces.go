package usecase

import "context"

type IdentityRepositoryInterface interface {
	CreateNewUser(context.Context, *CreateNewUserRequestDTO) (*CreateNewUserResponseDTO, error)
	GetUserProfileByUserId(context.Context, *GetUserProfileByUserIdRequestDTO) (*GetUserProfileByUserIdResponseDTO, error)
	CreateNewIdentity(context.Context, *CreateNewIdentityRequestDTO) error
	GetIdentityByTelegramId(context.Context, *GetIdentityByTelegramIdRequestDTO) (*GetIdentityByTelegramIdResponseDTO, error)
	GetIdentityByUserId(context.Context, *GetIdentityByUserIdRequestDTO) (*GetIdentityByUserIdResponseDTO, error)
}

type TxManagerInterface interface {
	WithinTx(context.Context, func(context.Context) error) error
}
