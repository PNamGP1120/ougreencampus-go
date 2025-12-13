package system

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type SystemConfig struct {
	Key       string  `gorm:"primaryKey"`
	Value     string  `gorm:"type:text"`
	UpdatedBy *string `gorm:"type:uuid"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type AuditLog struct {
	ID        string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ActorID   *string        `gorm:"type:uuid;index"`
	Role      string         `gorm:"default:'';index"`
	Action    string         `gorm:"not null;index"`   // e.g. CREATE_EVENT, UPDATE_CONFIG
	Entity    string         `gorm:"default:'';index"` // user/content/event/activity/system
	EntityID  string         `gorm:"default:'';index"`
	Metadata  datatypes.JSON `gorm:"type:jsonb"`
	IP        string         `gorm:"default:''"`
	UserAgent string         `gorm:"type:text"`
	CreatedAt time.Time
}
