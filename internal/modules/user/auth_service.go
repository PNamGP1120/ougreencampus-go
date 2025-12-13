package user

import (
	"errors"

	"github.com/PNamGP1120/ougreencampus-go/internal/utils"
	"github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
)

type AuthService interface {
	Login(req LoginRequest) (string, error)
}

type authService struct {
	repo   Repository
	jwtSvc *jwt.JWTService
}

func NewAuthService(repo Repository, jwtSvc *jwt.JWTService) AuthService {
	return &authService{repo: repo, jwtSvc: jwtSvc}
}

func (a *authService) Login(req LoginRequest) (string, error) {
	user, err := a.repo.FindByEmail(req.Email)
	if err != nil || !user.IsActive {
		return "", errors.New("invalid credentials")
	}

	if err := utils.CheckPassword(user.Password, req.Password); err != nil {
		return "", errors.New("invalid credentials")
	}

	return a.jwtSvc.GenerateToken(user.ID, user.Role)
}
