package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gotemplate/internal/auth/types"
	"gotemplate/pkg/customerror"
	"time"
)

func GeneratePair(userId, key string) (types.Tokens, error) {
	accessToken, err := Generate(userId, time.Minute*15).SignedString([]byte(key))
	if err != nil {
		return types.Tokens{}, err
	}

	refreshToken, err := Generate(userId, time.Hour*24).SignedString([]byte(key))
	if err != nil {
		return types.Tokens{}, err
	}

	return types.Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func Generate(userId string, exp time.Duration) *jwt.Token {
	issuedAt := time.Now()
	expiration := issuedAt.Add(exp)

	claims := jwt.RegisteredClaims{
		Issuer:    "gotemplate",
		Subject:   userId,
		ID:        uuid.NewString(),
		ExpiresAt: jwt.NewNumericDate(expiration),
		IssuedAt:  jwt.NewNumericDate(issuedAt),
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
}

func Verify(key, token string) (jwt.Claims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, customerror.Wrap("jwt.Parse: %w", err)
	}

	if !jwtToken.Valid {
		return nil, customerror.New(customerror.InvalidJWTTokenErrorCode, "invalid token")
	}

	return jwtToken.Claims, nil
}
