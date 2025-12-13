package user

import (
	"errors"

	"github.com/PNamGP1120/ougreencampus-go/internal/utils"
)

type Service interface {
	CreateUser(req CreateUserRequest) error
	GetAllUsers() ([]UserResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(req CreateUserRequest) error {
	if !utils.IsValidEmail(req.Email) {
		return errors.New("invalid email")
	}

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := &User{
		Email:    req.Email,
		Password: hash,
		Role:     req.Role,
		IsActive: true,
	}

	return s.repo.Create(user)
}

func (s *service) GetAllUsers() ([]UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var result []UserResponse
	for _, u := range users {
		result = append(result, UserResponse{
			ID:       u.ID,
			Email:    u.Email,
			Role:     u.Role,
			IsActive: u.IsActive,
		})
	}

	return result, nil
}
