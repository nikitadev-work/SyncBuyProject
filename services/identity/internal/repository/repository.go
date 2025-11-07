package repository

import (
	"context"
	"identity/internal/domain"
	uc "identity/internal/usecase"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pgPool   *pgxpool.Pool
	txMarker string
}

func NewRepository(pgPool *pgxpool.Pool, txMarker string) *Repository {
	return &Repository{
		pgPool:   pgPool,
		txMarker: txMarker,
	}
}

func (repo *Repository) CreateNewUser(ctx context.Context, input *uc.CreateNewUserRequestDTO) (*uc.CreateNewUserResponseDTO, error) {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	userId := uuid.New()
	reqStr := "INSERT INTO users (id, first_name, last_name, status) VALUES ($1, $2, $3, $4);"

	if ok {
		// Command within transaction
		_, err := tx.Exec(ctx, reqStr, userId, input.FirstName, input.LastName, input.Status)
		if err != nil {
			return nil, err
		}
	} else {
		// Command without transaction
		_, err := repo.pgPool.Exec(ctx, reqStr, userId, input.FirstName, input.LastName, input.Status)
		if err != nil {
			return nil, err
		}
	}

	return &uc.CreateNewUserResponseDTO{
		UserId: userId,
	}, nil
}

func (repo *Repository) GetUserProfileByUserId(ctx context.Context, input *uc.GetUserProfileByUserIdRequestDTO) (*uc.GetUserProfileByUserIdResponseDTO, error) {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "SELECT id, first_name, last_name, status, created_at, updated_at FROM users WHERE id = $1"
	var res pgx.Row
	var err error

	if ok {
		// Command within transaction
		res = tx.QueryRow(ctx, reqStr, input.UserId)
	} else {
		// Command without transaction
		res = repo.pgPool.QueryRow(ctx, reqStr, input.UserId)
	}

	var user domain.User
	err = res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Status, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, ErrParsingResponse
	}

	return &uc.GetUserProfileByUserIdResponseDTO{
		User: user,
	}, nil
}

func (repo *Repository) CreateNewIdentity(ctx context.Context, input *uc.CreateNewIdentityRequestDTO) error {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "INSERT INTO user_identities (external_id, internal_id, provider_type, chat_id, meta)" +
		" VALUES ($1, $2, $3, $4, $5)"

	if ok {
		// Command within transaction
		_, err := tx.Exec(ctx, reqStr, input.ExternalId, input.InternalId, input.ProviderType, input.ChatId, input.Meta)
		if err != nil {
			return err
		}
	} else {
		// Command without transaction
		_, err := repo.pgPool.Exec(ctx, reqStr, input.ExternalId, input.InternalId, input.ProviderType, input.ChatId, input.Meta)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *Repository) GetIdentityByTelegramId(ctx context.Context, input *uc.GetIdentityByTelegramIdRequestDTO) (*uc.GetIdentityByTelegramIdResponseDTO, error) {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "SELECT external_id, internal_id, provider_type, created_at, updated_at, chat_id, meta FROM user_identities WHERE external_id = $1 AND provider_type = $2"
	var res pgx.Row
	var err error

	if ok {
		// Command within transaction
		res = tx.QueryRow(ctx, reqStr, input.TelegramId, "telegram")
	} else {
		// Command without transaction
		res = repo.pgPool.QueryRow(ctx, reqStr, input.TelegramId, "telegram")
	}

	var identity domain.Identity
	err = res.Scan(&identity.ExternalId, &identity.InternalId, &identity.ProviderType,
		&identity.CreatedAt, &identity.UpdatedAt, &identity.ChatId, &identity.Meta)
	if err != nil {
		return nil, ErrParsingResponse
	}

	return &uc.GetIdentityByTelegramIdResponseDTO{
		Identity: identity,
	}, nil
}

func (repo *Repository) GetIdentityByUserId(ctx context.Context, input *uc.GetIdentityByUserIdRequestDTO) (*uc.GetIdentityByUserIdResponseDTO, error) {
	tx, ok := ctx.Value(repo.txMarker).(pgx.Tx)
	reqStr := "SELECT external_id, internal_id, provider_type, created_at, updated_at, chat_id, meta FROM user_identities WHERE internal_id = $1"
	var res pgx.Row
	var err error

	if ok {
		// Command within transaction
		res = tx.QueryRow(ctx, reqStr, input.UserId)
	} else {
		// Command without transaction
		res = repo.pgPool.QueryRow(ctx, reqStr, input.UserId)
	}

	var identity domain.Identity
	err = res.Scan(&identity.ExternalId, &identity.InternalId, &identity.ProviderType,
		&identity.CreatedAt, &identity.UpdatedAt, &identity.ChatId, &identity.Meta)
	if err != nil {
		return nil, ErrParsingResponse
	}

	return &uc.GetIdentityByUserIdResponseDTO{
		Identity: identity,
	}, nil
}
