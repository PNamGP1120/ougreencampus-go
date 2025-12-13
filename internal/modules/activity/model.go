package activity

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID          string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name        string `gorm:"not null"`
	Description string `gorm:"type:text"`
	Type        string `gorm:"not null"` // program | contest | campaign
	StartDate   time.Time
	EndDate     time.Time
	CreatedBy   string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Submission struct {
	ID         string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ActivityID string `gorm:"index"`
	UserID     string `gorm:"index"`
	Content    string `gorm:"type:text"`
	Score      int
	Status     string `gorm:"default:pending"`
	CreatedAt  time.Time
}
