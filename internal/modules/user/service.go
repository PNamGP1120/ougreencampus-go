package user

import (
	"errors"
	"time"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/pkg/hash"
	jwtutil "github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Register(req RegisterRequest) error {
	if _, err := s.repo.FindByEmail(req.Email); err == nil {
		return errors.New("email already exists")
	}

	pw, _ := hash.HashPassword(req.Password)

	u := &User{
		Email:    req.Email,
		Password: pw,
		Name:     req.Name,
		Avatar:   req.Avatar,
		Role:     "student",
		Status:   "active",
	}
	return s.repo.Create(u)
}

func (s *Service) Login(req LoginRequest) (string, *User, error) {
	u, err := s.repo.FindByEmail(req.Email)
	if err != nil || !hash.CheckPassword(u.Password, req.Password) {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := jwtutil.GenerateAccessToken(
		u.ID,
		u.Role,
		config.Cfg.JWT.Secret,
		config.Cfg.JWT.AccessTTL,
	)
	if err != nil {
		return "", nil, err
	}

	u.Password = ""
	return token, u, nil
}

func (s *Service) RefreshToken(uid uint, role string) (string, error) {
	return jwtutil.GenerateAccessToken(
		uid,
		role,
		config.Cfg.JWT.Secret,
		config.Cfg.JWT.AccessTTL,
	)
}
