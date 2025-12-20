package user

import (
	"errors"
	"time"

	"github.com/PNamGP1120/ougreencampus-go/pkg/hash"
	jwtutil "github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
)

type Service interface {
	Register(req RegisterRequest) (*User, error)
	Login(req LoginRequest, jwtSecret string) (*AuthResponse, error)
	GetByID(id uint) (*User, error)
	UpdateProfile(id uint, req UpdateProfileRequest) error
	UpdatePassword(id uint, req UpdatePasswordRequest) error
	UpdateRole(id uint, role string) error
	UpdateStatus(id uint, status string) error
	List() ([]User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(req RegisterRequest) (*User, error) {
	hashed, _ := hash.HashPassword(req.Password)

	user := &User{
		Email:    req.Email,
		Password: hashed,
		Name:     req.Name,
		Avatar:   req.Avatar,
		Role:     "student",
		Status:   "active",
	}

	return user, s.repo.Create(user)
}

func (s *service) Login(req LoginRequest, secret string) (*AuthResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil || !hash.CheckPassword(user.Password, req.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, _ := jwtutil.GenerateToken(
		user.ID,
		user.Role,
		secret,
		time.Hour*24,
	)

	return &AuthResponse{
		Token: token,
		User:  *user,
	}, nil
}

func (s *service) GetByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *service) UpdateProfile(id uint, req UpdateProfileRequest) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	user.Name = req.Name
	user.Avatar = req.Avatar
	return s.repo.Update(user)
}

func (s *service) UpdatePassword(id uint, req UpdatePasswordRequest) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if !hash.CheckPassword(user.Password, req.OldPassword) {
		return errors.New("wrong password")
	}
	hashed, _ := hash.HashPassword(req.NewPassword)
	user.Password = hashed
	return s.repo.Update(user)
}

func (s *service) UpdateRole(id uint, role string) error {
	user, _ := s.repo.FindByID(id)
	user.Role = role
	return s.repo.Update(user)
}

func (s *service) UpdateStatus(id uint, status string) error {
	user, _ := s.repo.FindByID(id)
	user.Status = status
	return s.repo.Update(user)
}

func (s *service) List() ([]User, error) {
	return s.repo.List()
}
