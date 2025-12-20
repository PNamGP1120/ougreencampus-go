package system

import "gorm.io/gorm"

type SystemConfig struct {
	gorm.Model
	Key   string `gorm:"uniqueIndex"`
	Value string `gorm:"type:text"`
}

type AuditLog struct {
	gorm.Model
	UserID uint
	Action string
	Target string
}

type Notification struct {
	gorm.Model
	UserID uint
	Title  string
	Body   string
	Read   bool
}
