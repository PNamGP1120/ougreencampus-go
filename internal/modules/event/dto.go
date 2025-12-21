package event

import "time"

// Event
type CreateEventRequest struct {
	Title     string    `json:"title" binding:"required"`
	Image     string    `json:"image"`
	Location  string    `json:"location"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type UpdateEventRequest struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

// Check-in
type ManualCheckinRequest struct {
	UserID uint `json:"user_id"`
}
