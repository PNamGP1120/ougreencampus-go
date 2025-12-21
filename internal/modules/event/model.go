package event

import "time"

type Event struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Image     string    `json:"image"`
	Location  string    `json:"location"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedBy uint      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type EventRegistration struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	EventID   uint      `json:"event_id"`
	UserID    uint      `json:"user_id"`
	CheckedIn bool      `json:"checked_in"`
	CreatedAt time.Time `json:"created_at"`
}
