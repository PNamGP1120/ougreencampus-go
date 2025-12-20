package media

import "time"

type Media struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	URL       string    `gorm:"not null" json:"url"`
	Type      string    `gorm:"not null" json:"type"` // image | video
	OwnerID   uint      `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
}
