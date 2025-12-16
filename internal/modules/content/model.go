package content

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Content struct {
	ID         string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title      string         `gorm:"not null"`
	Body       string         `gorm:"type:text"`
	Type       string         `gorm:"not null"` // news | blog | green_news
	CoverImage string         `gorm:"type:text"`
	Images     datatypes.JSON `gorm:"type:json"`
	IsFeatured bool           `gorm:"default:false"`

	CategoryID *string
	Category   *Category

	Tags []Tag `gorm:"many2many:content_tags;"`

	AuthorID  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string `gorm:"unique;not null"`
	Slug      string `gorm:"unique;not null"`
	CreatedAt time.Time
}

type Tag struct {
	ID        string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string `gorm:"unique;not null"`
	Slug      string `gorm:"unique;not null"`
	CreatedAt time.Time
}
