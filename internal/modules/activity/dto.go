package activity

/* ========= ACTIVITY ========= */

type CreateActivityRequest struct {
	Title       string       `json:"title" binding:"required"`
	Type        ActivityType `json:"type" binding:"required"`
	Description string       `json:"description"`
	Image       *string      `json:"image"`
}

type UpdateActivityRequest struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Image       *string        `json:"image"`
	Status      ActivityStatus `json:"status"`
}

/* ========= SUBMISSION ========= */

type SubmitRequest struct {
	Content string  `json:"content"`
	FileURL *string `json:"file_url"`
}

type ReviewSubmissionRequest struct {
	Score   int     `json:"score"`
	Comment *string `json:"comment"`
}

/* ========= CAMPAIGN ========= */

type CreateTaskRequest struct {
	Title  string `json:"title"`
	Target int    `json:"target"`
}

type ProgressRequest struct {
	Value int `json:"value"`
}

/* ========= LIST FILTER ========= */

type ListActivityFilter struct {
	Search string
	Type   string
	Status string
	Page   int
	Limit  int
}
