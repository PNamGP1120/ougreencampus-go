package content

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{svc: s}
}

/* ========= CONTENT ========= */

func (h *Handler) ListContent(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	filter := ListContentFilter{
		Search:     c.Query("search"),
		CategoryID: c.Query("category_id"),
		TagID:      c.Query("tag_id"),
		Page:       page,
		Limit:      limit,
	}

	items, total := h.svc.ListContent(filter)

	c.JSON(200, gin.H{
		"items": items,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) GetContent(c *gin.Context) {
	data, _ := h.svc.GetContent(c.Param("id"))
	c.JSON(200, data)
}

func (h *Handler) CreateContent(c *gin.Context) {
	var req CreateContentRequest
	_ = c.ShouldBindJSON(&req)
	id, _ := h.svc.CreateContent(req, c.GetString("user_id"))
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) UpdateContent(c *gin.Context) {
	var req UpdateContentRequest
	_ = c.ShouldBindJSON(&req)
	_ = h.svc.UpdateContent(c.Param("id"), req)
	c.JSON(200, gin.H{"message": "updated"})
}

func (h *Handler) DeleteContent(c *gin.Context) {
	_ = h.svc.DeleteContent(c.Param("id"))
	c.JSON(200, gin.H{"message": "deleted"})
}

/* ========= CATEGORY ========= */

func (h *Handler) ListCategories(c *gin.Context) {
	data, _ := h.svc.ListCategories()
	c.JSON(200, data)
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	_ = c.ShouldBindJSON(&req)
	id, _ := h.svc.CreateCategory(req.Name)
	c.JSON(200, gin.H{"id": id})
}

/* ========= TAG ========= */

func (h *Handler) ListTags(c *gin.Context) {
	data, _ := h.svc.ListTags()
	c.JSON(200, data)
}

func (h *Handler) CreateTag(c *gin.Context) {
	var req CreateTagRequest
	_ = c.ShouldBindJSON(&req)
	id, _ := h.svc.CreateTag(req.Name)
	c.JSON(200, gin.H{"id": id})
}

/* ========= MEDIA ========= */

func (h *Handler) UploadMedia(c *gin.Context) {
	id, url := h.svc.UploadMedia()
	c.JSON(200, gin.H{"id": id, "media_url": url})
}

func (h *Handler) ListMedia(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	items, total := h.svc.ListMedia(page, limit)

	c.JSON(200, gin.H{
		"items": items,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) DeleteMedia(c *gin.Context) {
	_ = h.svc.DeleteMedia(c.Param("id"))
	c.JSON(200, gin.H{"message": "deleted"})
}

func (h *Handler) AttachMedia(c *gin.Context) {
	var req AttachMediaRequest
	_ = c.ShouldBindJSON(&req)
	_ = h.svc.AttachMedia(c.Param("id"), req.Type, req.RefID)
	c.JSON(200, gin.H{"message": "attached"})
}
