package content

type CreateContentRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
	Type  string `json:"type" binding:"required"`
}

type UpdateContentRequest struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	Type       string `json:"type"`
	IsFeatured *bool  `json:"is_featured"`
}
