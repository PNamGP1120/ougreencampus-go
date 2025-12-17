package user

import (
	"errors"

	"github.com/PNamGP1120/ougreencampus-go/pkg/hash"
	"github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

/* ===== AUTH ===== */

func (s *Service) Register(req RegisterRequest) (*User, error) {
	hashed, _ := hash.HashPassword(req.Password)

	u := &User{
		Email:    req.Email,
		Password: hashed,
		Name:     req.Name,
		Avatar:   req.Avatar,
		Role:     RoleStudent,
		Status:   "active",
	}
	return u, s.repo.Create(u)
}

func (s *Service) Login(email, password string) (string, *User, error) {
	u, err := s.repo.FindByEmail(email)
	if err != nil || !hash.CheckPassword(password, u.Password) {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := jwt.GenerateToken(u.ID, string(u.Role))
	return token, u, err
}

/* ===== USER ===== */

func (s *Service) ListUsers() ([]User, error) {
	return s.repo.List()
}

func (s *Service) GetUser(id string) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *Service) CreateUser(req CreateUserRequest) (*User, error) {
	hashed, _ := hash.HashPassword(req.Password)

	u := &User{
		Email:    req.Email,
		Password: hashed,
		Name:     req.Name,
		Avatar:   req.Avatar,
		Role:     req.Role,
		Status:   "active",
	}
	return u, s.repo.Create(u)
}

func (s *Service) UpdateUser(id string, req UpdateUserRequest) error {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	u.Name = req.Name
	u.Avatar = req.Avatar
	return s.repo.Update(u)
}

func (s *Service) ChangeRole(id string, role Role) error {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	u.Role = role
	return s.repo.Update(u)
}

func (s *Service) ChangeStatus(id, status string) error {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	u.Status = status
	return s.repo.Update(u)
}

func (s *Service) ChangePassword(id, oldPwd, newPwd string) error {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if !hash.CheckPassword(oldPwd, u.Password) {
		return errors.New("wrong password")
	}
	hashed, _ := hash.HashPassword(newPwd)
	u.Password = hashed
	return s.repo.Update(u)
}
