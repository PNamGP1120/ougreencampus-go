package user

import "time"

type Role string

const (
	RoleAdmin     Role = "admin"
	RoleOrganizer Role = "organizer"
	RoleStudent   Role = "student"
)

type User struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Name      string
	Avatar    string
	Role      Role   `gorm:"type:varchar(20);default:'student'"`
	Status    string `gorm:"type:varchar(20);default:'active'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
