package activity

import "time"

type CreateActivityRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Type        string    `json:"type" binding:"required"` // program|contest|campaign
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}

type SubmitRequest struct {
	Content string `json:"content" binding:"required"`
}
