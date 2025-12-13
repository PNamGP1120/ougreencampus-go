package content

import (
	"time"

	"gorm.io/gorm"
)

type Content struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title     string `gorm:"not null"`
	Body      string `gorm:"type:text"`
	Type      string `gorm:"not null"` // news | blog | green_news
	Status    string `gorm:"default:draft"`
	AuthorID  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
