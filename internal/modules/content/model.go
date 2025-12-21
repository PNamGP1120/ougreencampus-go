package content

import "time"

type Content struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	Body       string    `gorm:"type:text" json:"body"`
	Image      string    `json:"image"`
	CategoryID uint      `json:"category_id"`
	AuthorID   uint      `json:"author_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	Category Category `gorm:"foreignKey:CategoryID"`
	Tags     []Tag    `gorm:"many2many:content_tags"`
}

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique" json:"name"`
}

type Tag struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique" json:"name"`
}

type ContentTag struct {
	ContentID uint
	TagID     uint
}
