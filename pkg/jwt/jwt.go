package jwt

import (
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey string
	issuer    string
	expire    time.Duration
}

type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwtlib.RegisteredClaims
}

func NewJWTService(secret string, expireMinutes int) *JWTService {
	return &JWTService{
		secretKey: secret,
		issuer:    "ougreencampus",
		expire:    time.Duration(expireMinutes) * time.Minute,
	}
}

func (j *JWTService) GenerateToken(userID string, role string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwtlib.RegisteredClaims{
			Issuer:    j.issuer,
			ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(j.expire)),
			IssuedAt:  jwtlib.NewNumericDate(time.Now()),
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *JWTService) ValidateToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwtlib.ParseWithClaims(
		tokenStr,
		&CustomClaims{},
		func(token *jwtlib.Token) (interface{}, error) {
			return []byte(j.secretKey), nil
		},
	)

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}
