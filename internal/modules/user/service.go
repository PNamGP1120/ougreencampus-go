package user

import (
	"errors"

	"github.com/PNamGP1120/ougreencampus-go/internal/utils"
)

type Service interface {
	CreateUser(req CreateUserRequest) (UserResponse, error)
	GetAll() ([]UserResponse, error)
	GetByID(id string) (UserResponse, error)
	UpdateRole(userID string, req UpdateRoleRequest) error
	SetActive(userID string, active bool) error
	UpdateAvatar(userID string, avatar string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(req CreateUserRequest) (UserResponse, error) {
	if _, err := s.repo.FindByEmail(req.Email); err == nil {
		return UserResponse{}, errors.New("email already exists")
	}

	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return UserResponse{}, err
	}

	u := &User{
		Email:    req.Email,
		Password: hashed,
		Role:     req.Role,
		Avatar:   req.Avatar, // NEW
		IsActive: true,
	}

	if err := s.repo.Create(u); err != nil {
		return UserResponse{}, err
	}

	return ToUserResponse(*u), nil
}

func (s *service) GetAll() ([]UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	out := make([]UserResponse, 0, len(users))
	for _, u := range users {
		out = append(out, ToUserResponse(u))
	}
	return out, nil
}

func (s *service) GetByID(id string) (UserResponse, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return UserResponse{}, err
	}
	return ToUserResponse(*u), nil
}

func (s *service) UpdateRole(userID string, req UpdateRoleRequest) error {
	u, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}
	u.Role = req.Role
	return s.repo.Update(u)
}

func (s *service) SetActive(userID string, active bool) error {
	u, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}
	u.IsActive = active
	return s.repo.Update(u)
}

func (s *service) UpdateAvatar(userID string, avatar string) error {
	u, err := s.repo.FindByID(userID)
	if err != nil {
		return err
	}
	u.Avatar = avatar
	return s.repo.Update(u)
}
