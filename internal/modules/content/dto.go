package content

type CreateContentRequest struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body"`
	Type  string `json:"type" binding:"required"`
}

type ContentResponse struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Type   string `json:"type"`
	Status string `json:"status"`
}
