package usecase

import (
	"context"
	"errors"
	"identity/internal/domain"
)

type IdentityUsecaseInterface interface {
	RegisterOrGetUserByTelegram(context.Context, RegisterOrGetUserByTelegramInputDTO) (RegisterOrGetUserByTelegramOutputDTO, error)
	GetUserByUserId(context.Context, GetUserByUserIdInputDTO) (GetUserByUserIdOutputDTO, error)
	GetUserByTelegram(context.Context, GetUserByTelegramInputDTO) (GetUserByTelegramOutputDTO, error)
}

type IdentityUsecase struct {
	Repository IdentityRepositoryInterface
}

func NewIdentityUsecase(repository IdentityRepositoryInterface) *IdentityUsecase {
	return &IdentityUsecase{
		Repository: repository,
	}
}

var _ IdentityUsecaseInterface = (*IdentityUsecase)(nil)

func (iu *IdentityUsecase) RegisterOrGetUserByTelegram(ctx context.Context, input RegisterOrGetUserByTelegramInputDTO) (RegisterOrGetUserByTelegramOutputDTO, error) {
	repoIdentityRequest := GetIdentityByTelegramIdRequestDTO{
		TelegramId: input.TelegramId,
	}

	identity, err := iu.Repository.GetIdentityByTelegramId(repoIdentityRequest)
	if err == nil {
		// User exists
		repoGetRequest := GetUserProfileByUserIdRequestDTO{
			UserId: identity.Identity.InternalId,
		}

		usr, err := iu.Repository.GetUserProfileByUserId(repoGetRequest)
		if err == nil {
			return RegisterOrGetUserByTelegramOutputDTO{
				UserId: usr.User.Id,
			}, nil
		}
		return RegisterOrGetUserByTelegramOutputDTO{}, errors.New(ErrUserDoesNotExist.Error() + " while identity for this user exists")
	}

	// Register new User

	repoCreateRequest := CreateNewUserRequestDTO{
		Id:        identity.Identity.InternalId,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Status:    input.Status,
	}

	resp, err := iu.Repository.CreateNewUser(repoCreateRequest)
	if err != nil {
		return RegisterOrGetUserByTelegramOutputDTO{}, ErrCreateUser
	}

	repoCreateIdentityRequest := CreateNewIdentityRequestDTO{
		ExternalId:   input.TelegramId,
		InternalId:   resp.UserId,
		ProviderType: domain.Telegram,
		ChatId:       input.ChatId,
		Meta:         input.Meta,
	}

	err = iu.Repository.CreateNewIdentity(repoCreateIdentityRequest)
	if err != nil {
		return RegisterOrGetUserByTelegramOutputDTO{}, ErrCreateIdentity
	}

	return RegisterOrGetUserByTelegramOutputDTO{
		UserId: resp.UserId,
	}, nil
}

func (iu *IdentityUsecase) GetUserByUserId(ctx context.Context, input GetUserByUserIdInputDTO) (GetUserByUserIdOutputDTO, error) {
	repoGetRequest := GetUserProfileByUserIdRequestDTO{
		UserId: input.UserId,
	}

	resp, err := iu.Repository.GetUserProfileByUserId(repoGetRequest)
	if err != nil {
		return GetUserByUserIdOutputDTO{}, ErrUserDoesNotExist
	}

	return GetUserByUserIdOutputDTO{
		resp.User,
	}, nil

}

func (iu *IdentityUsecase) GetUserByTelegram(ctx context.Context, input GetUserByTelegramInputDTO) (GetUserByTelegramOutputDTO, error) {
	repoIdentityRequest := GetIdentityByTelegramIdRequestDTO{
		TelegramId: input.TelegramId,
	}

	identity, err := iu.Repository.GetIdentityByTelegramId(repoIdentityRequest)
	if err != nil {
		return GetUserByTelegramOutputDTO{}, ErrIdentityDoesNotExist
	}

	// User exists
	repoGetRequest := GetUserProfileByUserIdRequestDTO{
		UserId: identity.Identity.InternalId,
	}

	usr, err := iu.Repository.GetUserProfileByUserId(repoGetRequest)
	if err == nil {
		return GetUserByTelegramOutputDTO{
			User: usr.User,
		}, nil
	}
	return GetUserByTelegramOutputDTO{}, ErrUserDoesNotExist
}
