package content

// ---------- REQUEST ----------

type CreateContentRequest struct {
	Title      string   `json:"title" binding:"required"`
	Body       string   `json:"body"`
	Type       string   `json:"type" binding:"required"`
	CoverImage string   `json:"cover_image"`
	Images     []string `json:"images"`
	CategoryID *string  `json:"category_id"`
	TagIDs     []string `json:"tag_ids"`
}

type UpdateContentRequest struct {
	Title      *string   `json:"title"`
	Body       *string   `json:"body"`
	Type       *string   `json:"type"`
	CoverImage *string   `json:"cover_image"`
	Images     *[]string `json:"images"`
	IsFeatured *bool     `json:"is_featured"`
	CategoryID *string   `json:"category_id"`
	TagIDs     *[]string `json:"tag_ids"`
}

// ---------- RESPONSE ----------

type ContentResponse struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Type       string    `json:"type"`
	CoverImage string    `json:"cover_image"`
	Images     []string  `json:"images"`
	IsFeatured bool      `json:"is_featured"`
	Category   *Category `json:"category"`
	Tags       []Tag     `json:"tags"`
	CreatedAt  string    `json:"created_at"`
}
