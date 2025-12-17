package content

import "time"

type Content struct {
	ID         string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title      string
	Body       string `gorm:"type:text"`
	Image      *string
	AuthorID   string
	CategoryID *string `gorm:"index"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Category struct {
	ID   string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string `gorm:"uniqueIndex"`
}

type Tag struct {
	ID   string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name string `gorm:"uniqueIndex"`
}

type ContentTag struct {
	ContentID string `gorm:"index"`
	TagID     string `gorm:"index"`
}

type Media struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	URL       string
	Type      string
	CreatedAt time.Time
}
