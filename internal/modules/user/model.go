package user

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Role      string    `gorm:"default:student" json:"role"`
	Status    string    `gorm:"default:active" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
