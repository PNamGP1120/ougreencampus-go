package activity

// ===== ACTIVITY =====
type CreateActivityRequest struct {
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type UpdateActivityRequest struct {
	Title string `json:"title"`
	Image string `json:"image"`
}

// ===== CONTEST =====
type SubmitContestRequest struct {
	Content string `json:"content"`
	FileURL string `json:"file_url"`
}

type ReviewSubmissionRequest struct {
	Score   int    `json:"score"`
	Comment string `json:"comment"`
}

// ===== CAMPAIGN =====
type CreateTaskRequest struct {
	Title  string `json:"title"`
	Target int    `json:"target"`
}

type ProgressRequest struct {
	Value int `json:"value"`
}

// ===== PROGRAM =====
type AddChildRequest struct {
	ChildID uint `json:"child_id"`
}
