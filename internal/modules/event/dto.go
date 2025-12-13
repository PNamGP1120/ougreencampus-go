package event

import "time"

type CreateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Location    string    `json:"location"`
	Capacity    int       `json:"capacity"`
}
