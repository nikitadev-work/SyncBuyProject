package usecase

type IdentityRepositoryInterface interface {
	CreateNewUser(CreateNewUserRequestDTO) (CreateNewUserResponseDTO, error)
	GetUserProfileByUserId(GetUserProfileByUserIdRequestDTO) (GetUserProfileByUserIdResponseDTO, error)
	CreateNewIdentity(CreateNewIdentityRequestDTO) error
	GetIdentityByTelegramId(GetIdentityByTelegramIdRequestDTO) (GetIdentityByTelegramIdResponseDTO, error)
	GetIdentityByUserId(GetIdentityByUserIdRequestDTO) (GetIdentityByUserIdResponseDTO, error)
}
