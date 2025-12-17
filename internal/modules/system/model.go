package system

import "time"

/* ========= SYSTEM CONFIG ========= */

type SystemConfig struct {
	Key       string `gorm:"primaryKey"`
	Value     string
	UpdatedAt time.Time
}

/* ========= AUDIT LOG ========= */

type AuditLog struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    string `gorm:"index"`
	Action    string `gorm:"index"`
	Resource  string
	CreatedAt time.Time
}

/* ========= NOTIFICATION ========= */

type Notification struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    string `gorm:"index"`
	Title     string
	Content   string
	IsRead    bool `gorm:"index"`
	CreatedAt time.Time
}
