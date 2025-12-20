package jwt

import (
	"errors"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

/*
Claims
- UserID: dùng uint (đúng chuẩn DB / GORM)
- Role: RBAC
*/
type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwtlib.RegisteredClaims
}

/*
GenerateAccessToken
- Sinh access token
*/
func GenerateAccessToken(
	userID uint,
	role string,
	secret string,
	ttl time.Duration,
) (string, error) {

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

/*
GenerateRefreshToken
- Sinh refresh token (không cần role)
*/
func GenerateRefreshToken(
	userID uint,
	secret string,
	ttl time.Duration,
) (string, error) {

	claims := jwtlib.RegisteredClaims{
		Subject:   "refresh",
		ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(ttl)),
		IssuedAt:  jwtlib.NewNumericDate(time.Now()),
	}

	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

/*
ParseToken
- Dùng cho middleware
*/
func ParseToken(
	tokenStr string,
	secret string,
) (*Claims, error) {

	token, err := jwtlib.ParseWithClaims(
		tokenStr,
		&Claims{},
		func(token *jwtlib.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwtlib.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
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
