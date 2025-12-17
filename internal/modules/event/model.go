package event

import "time"

type Event struct {
	ID          string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string
	Description string `gorm:"type:text"`
	Image       *string
	Location    string
	StartTime   time.Time
	EndTime     time.Time
	Capacity    int
	CreatedBy   string `gorm:"index"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Registration struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EventID   string `gorm:"index"`
	UserID    string `gorm:"index"`
	CheckedIn bool
	CreatedAt time.Time
}
