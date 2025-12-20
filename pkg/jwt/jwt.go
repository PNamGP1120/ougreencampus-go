package jwt

import (
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwtlib.RegisteredClaims
}

// GenerateToken creates a JWT token
func GenerateToken(userID uint, role string, secret string, ttl time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwtlib.NewNumericDate(time.Now()),
		},
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken validates token and returns claims
func ParseToken(tokenStr, secret string) (*Claims, error) {
	token, err := jwtlib.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwtlib.Token) (interface{}, error) {
			return []byte(secret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
