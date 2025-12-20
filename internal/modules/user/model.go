package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Email     string         `gorm:"uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	Name      string         `gorm:"not null" json:"name"`
	Avatar    string         `gorm:"default:''" json:"avatar"`
	Role      string         `gorm:"not null;default:'student'" json:"role"`  // guest/student/organizer/admin
	Status    string         `gorm:"not null;default:'active'" json:"status"` // active/blocked
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type PasswordReset struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Email     string     `gorm:"index;not null" json:"email"`
	Token     string     `gorm:"uniqueIndex;not null" json:"token"`
	ExpiresAt time.Time  `gorm:"not null" json:"expires_at"`
	UsedAt    *time.Time `json:"used_at"`
	CreatedAt time.Time  `json:"created_at"`
}
