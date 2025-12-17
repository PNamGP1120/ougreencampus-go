package event

import "time"

/* ========= EVENT ========= */

type CreateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Image       *string   `json:"image"`
	Location    string    `json:"location"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Capacity    int       `json:"capacity"`
}

type UpdateEventRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       *string   `json:"image"`
	Location    string    `json:"location"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Capacity    int       `json:"capacity"`
}

/* ========= CHECKIN ========= */

type CheckinQRRequest struct {
	QR string `json:"qr"`
}

type ManualCheckinRequest struct {
	UserID string `json:"user_id"`
}

/* ========= LIST FILTER ========= */

type ListEventFilter struct {
	Search string
	From   string
	To     string
	Page   int
	Limit  int
}
