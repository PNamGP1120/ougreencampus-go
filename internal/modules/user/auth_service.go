package user

import (
	"errors"
	"log"

	"github.com/PNamGP1120/ougreencampus-go/internal/utils"
	"github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
)

type AuthService interface {
	Login(email, password string) (string, error)
}

type authService struct {
	repo   Repository
	jwtSvc *jwt.JWTService
}

func NewAuthService(repo Repository, jwtSvc *jwt.JWTService) AuthService {
	return &authService{repo: repo, jwtSvc: jwtSvc}
}

func (a *authService) Login(email, password string) (string, error) {
	user, err := a.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !user.IsActive {
		return "", errors.New("account locked")
	}

	// DEBUG (tạm thời)
	if err := utils.CheckPassword(password, user.Password); err != nil {
		log.Println("❌ Password mismatch for:", email)
		return "", errors.New("invalid credentials")
	}

	log.Println("✅ Password OK for:", email)
	return a.jwtSvc.GenerateToken(user.ID, user.Role)
}
