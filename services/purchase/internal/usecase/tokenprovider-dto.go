package usecase

import (
	"time"

	"github.com/google/uuid"
)

type GenerateTokenRequest struct {
	PurchaseId uuid.UUID
	ExpiresAt  time.Time
}

type GenerateTokenResponse struct {
	Token string
}

type ParseTokenRequest struct {
	TokenString string
}

type ParseTokenResponse struct {
	PurchaseId uuid.UUID
	ExpiresAt  time.Time
}
