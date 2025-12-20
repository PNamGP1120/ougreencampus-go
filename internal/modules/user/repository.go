package user

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(u *User) error {
	return r.db.Create(u).Error
}

func (r *Repository) Update(u *User) error {
	return r.db.Save(u).Error
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var u User
	if err := r.db.Where("email = ?", email).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) FindByID(id uint) (*User, error) {
	var u User
	if err := r.db.First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) List() ([]User, error) {
	var users []User
	if err := r.db.Order("id desc").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

/* -------- Password reset -------- */

func (r *Repository) CreateResetToken(email, token string, expiresAt time.Time) error {
	// Nếu email đã có token chưa dùng, tạo mới vẫn được (tuỳ chính sách)
	pr := PasswordReset{
		Email:     email,
		Token:     token,
		ExpiresAt: expiresAt,
	}
	return r.db.Create(&pr).Error
}

func (r *Repository) UseResetToken(token string) (*PasswordReset, error) {
	var pr PasswordReset
	if err := r.db.Where("token = ?", token).First(&pr).Error; err != nil {
		return nil, err
	}
	if pr.UsedAt != nil {
		return nil, errors.New("token already used")
	}
	if time.Now().After(pr.ExpiresAt) {
		return nil, errors.New("token expired")
	}
	now := time.Now()
	pr.UsedAt = &now
	if err := r.db.Save(&pr).Error; err != nil {
		return nil, err
	}
	return &pr, nil
}
