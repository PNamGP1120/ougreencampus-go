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

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

type AttachMediaRequest struct {
	Type  string `json:"type" binding:"required"`
	RefID uint   `json:"ref_id" binding:"required"`
}
