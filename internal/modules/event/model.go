package event

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"type:text"`
	StartTime   time.Time
	EndTime     time.Time
	Location    string
	Capacity    int
	CreatedBy   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type EventRegistration struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	EventID   string `gorm:"index"`
	UserID    string `gorm:"index"`
	CheckedIn bool   `gorm:"default:false"`
	CreatedAt time.Time
}
