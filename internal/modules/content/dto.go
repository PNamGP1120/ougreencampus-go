package content

/* ========= CONTENT ========= */

type CreateContentRequest struct {
	Title      string   `json:"title" binding:"required"`
	Body       string   `json:"body" binding:"required"`
	Image      *string  `json:"image"`
	CategoryID *string  `json:"category_id"`
	Tags       []string `json:"tags"`
}

type UpdateContentRequest struct {
	Title string  `json:"title"`
	Body  string  `json:"body"`
	Image *string `json:"image"`
}

/* ========= CATEGORY ========= */

type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

/* ========= TAG ========= */

type CreateTagRequest struct {
	Name string `json:"name" binding:"required"`
}

/* ========= MEDIA ========= */

type AttachMediaRequest struct {
	Type  string `json:"type"`
	RefID string `json:"ref_id"`
}

/* ========= LIST FILTER ========= */

type ListContentFilter struct {
	Search     string
	CategoryID string
	TagID      string
	Page       int
	Limit      int
}
