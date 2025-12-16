package user

import (
	"errors"

	"github.com/PNamGP1120/ougreencampus-go/internal/utils"
	"github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
)

type AuthService interface {
	Login(email, password string) (string, error)
	Register(req RegisterRequest) (UserResponse, error)
}

type authService struct {
	repo   Repository
	jwtSvc *jwt.JWTService
}

func NewAuthService(repo Repository, jwtSvc *jwt.JWTService) AuthService {
	return &authService{repo: repo, jwtSvc: jwtSvc}
}

func (a *authService) Login(email, password string) (string, error) {
	u, err := a.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	if !u.IsActive {
		return "", errors.New("account locked")
	}
	if err := utils.CheckPassword(password, u.Password); err != nil {
		return "", errors.New("invalid credentials")
	}
	return a.jwtSvc.GenerateToken(u.ID, u.Role)
}

func (a *authService) Register(req RegisterRequest) (UserResponse, error) {
	if _, err := a.repo.FindByEmail(req.Email); err == nil {
		return UserResponse{}, errors.New("email already exists")
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return UserResponse{}, err
	}

	u := &User{
		Email:    req.Email,
		Password: hashed,
		Role:     "student", // FIX CỨNG
		Avatar:   req.Avatar,
		IsActive: true,
	}

	if err := a.repo.Create(u); err != nil {
		return UserResponse{}, err
	}

	return ToUserResponse(*u), nil
}
