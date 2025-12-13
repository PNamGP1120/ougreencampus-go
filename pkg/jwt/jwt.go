package jwt

import (
	"errors"
	"time"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret string
	expire time.Duration
	issuer string
}

// ✅ Giữ tương thích với code đang gọi jwt.New(...)
func New(secret string, expireMinutes int) *JWTService {
	return NewJWTService(secret, expireMinutes)
}

// ✅ Constructor chuẩn (bạn có thể dùng NewJWTService hoặc New đều được)
func NewJWTService(secret string, expireMinutes int) *JWTService {
	return &JWTService{
		secret: secret,
		expire: time.Duration(expireMinutes) * time.Minute,
		issuer: "ougreencampus",
	}
}

func (j *JWTService) GenerateToken(userID, role string) (string, error) {
	claims := jwtv5.MapClaims{
		"user_id": userID,
		"role":    role,
		"iss":     j.issuer,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(j.expire).Unix(),
	}

	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTService) ValidateToken(tokenStr string) (*jwtv5.Token, error) {
	return jwtv5.Parse(tokenStr, func(token *jwtv5.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtv5.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(j.secret), nil
	})
}
