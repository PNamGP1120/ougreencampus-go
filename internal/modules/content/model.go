package content

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Title      string
	Body       string `gorm:"type:text"`
	Image      string
	CategoryID uint
	Category   Category
	Tags       []Tag `gorm:"many2many:content_tags"`
	AuthorID   uint
}

type Category struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

type Tag struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

type Media struct {
	gorm.Model
	Type    string // image | video | file
	URL     string
	RefID   uint
	RefType string // content | event | activity
}
