package user

import (
	"errors"

	"github.com/PNamGP1120/ougreencampus-go/internal/utils"
)

type Service interface {
	CreateUser(req CreateUserRequest) error
	GetAll() ([]User, error)
	UpdateRole(userID string, req UpdateRoleRequest) error
	SetActive(userID string, active bool) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(req CreateUserRequest) error {
	if _, err := s.repo.FindByEmail(req.Email); err == nil {
		return errors.New("email already exists")
	}

	hashed, _ := utils.HashPassword(req.Password)

	user := &User{
		Email:    req.Email,
		Password: hashed,
		Role:     req.Role,
		IsActive: true,
	}

	return s.repo.Create(user)
}

func (s *service) GetAll() ([]User, error) {
	return s.repo.FindAll()
}

func (s *service) UpdateRole(userID string, req UpdateRoleRequest) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}

	user.Role = req.Role
	return s.repo.Update(user)
}

func (s *service) SetActive(userID string, active bool) error {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}

	user.IsActive = active
	return s.repo.Update(user)
}
