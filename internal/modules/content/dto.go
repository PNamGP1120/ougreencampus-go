package content

type CreateContentRequest struct {
	Title      string `json:"title" binding:"required"`
	Body       string `json:"body"`
	Image      string `json:"image"`
	CategoryID uint   `json:"category_id"`
	Tags       []uint `json:"tags"`
}

type UpdateContentRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Image string `json:"image"`
}
