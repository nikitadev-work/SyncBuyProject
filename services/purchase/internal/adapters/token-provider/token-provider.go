package tokenprovider

import (
	"time"

	uc "purchase/internal/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenProvider struct {
	secretKey string
}

var _ uc.TokenProviderInterface = (*TokenProvider)(nil)

func NewTokenProvider(secretKey string) *TokenProvider {
	return &TokenProvider{
		secretKey: secretKey,
	}
}

func (tp *TokenProvider) GenerateInviteToken(in *uc.GenerateTokenRequest) (*uc.GenerateTokenResponse, error) {
	claims := jwt.MapClaims{
		"purchaseId": in.PurchaseId.String(),
		"expiresAt":  in.ExpiresAt.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(tp.secretKey)
	if err != nil {
		return nil, err
	}

	return &uc.GenerateTokenResponse{
		Token: tokenString,
	}, nil
}

func (tp *TokenProvider) ParseAndValidateInviteToken(in *uc.ParseTokenRequest) (*uc.ParseTokenResponse, error) {
	token, err := jwt.Parse(in.TokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrParsingInviteToken
		}
		return tp.secretKey, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		purchaseId, err := uuid.Parse(claims["purchaseId"].(string))
		if err != nil {
			return nil, err
		}
		expiresAt := time.Unix(int64(claims["expiresAt"].(float64)), 0)

		return &uc.ParseTokenResponse{
			PurchaseId: purchaseId,
			ExpiresAt:  expiresAt,
		}, nil
	} else {
		return nil, err
	}
}
