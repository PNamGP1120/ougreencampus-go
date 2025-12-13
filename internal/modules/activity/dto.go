package activity

import "time"

type CreateActivityRequest struct {
	Name        string     `json:"name" binding:"required"`
	Description string     `json:"description"`
	Type        string     `json:"type" binding:"required"` // program|contest|campaign
	StartAt     *time.Time `json:"start_at"`
	EndAt       *time.Time `json:"end_at"`
}

type UpdateActivityRequest struct {
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Status      *string    `json:"status"` // draft|published|closed
	StartAt     *time.Time `json:"start_at"`
	EndAt       *time.Time `json:"end_at"`
}

type SubmitContestRequest struct {
	Content string `json:"content" binding:"required"`
}

type ReviewSubmissionRequest struct {
	Status   string `json:"status" binding:"required"` // approved|rejected
	Feedback string `json:"feedback"`
}

type ScoreSubmissionRequest struct {
	Score    int    `json:"score" binding:"required"`
	Feedback string `json:"feedback"`
}

type CreateTaskRequest struct {
	Title    string     `json:"title" binding:"required"`
	Detail   string     `json:"detail"`
	Points   int        `json:"points"`
	DueAt    *time.Time `json:"due_at"`
	IsActive *bool      `json:"is_active"`
}

type UpdateTaskRequest struct {
	Title    *string    `json:"title"`
	Detail   *string    `json:"detail"`
	Points   *int       `json:"points"`
	DueAt    *time.Time `json:"due_at"`
	IsActive *bool      `json:"is_active"`
}

type SubmitProgressRequest struct {
	Note     string                 `json:"note"`
	Evidence map[string]interface{} `json:"evidence"` // sẽ lưu jsonb
}

type ReviewProgressRequest struct {
	Status string `json:"status" binding:"required"` // approved|rejected
	Note   string `json:"note"`
}

type UpsertMetricRequest struct {
	Key   string  `json:"key" binding:"required"`
	Value float64 `json:"value" binding:"required"`
	Unit  string  `json:"unit"`
	Note  string  `json:"note"`
}
