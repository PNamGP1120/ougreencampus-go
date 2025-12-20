package content

import (
	"net/http"
	"strconv"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/gin-gonic/gin"
)

/* ---------- CONTENT ---------- */

func (h *Handler) ListContents(c *gin.Context) {
	items, _ := h.repo.ListContents()
	c.JSON(http.StatusOK, common.Success(items))
}

func (h *Handler) CreateContent(c *gin.Context) {
	var req CreateContentRequest
	c.ShouldBindJSON(&req)

	content := Content{
		Title:      req.Title,
		Body:       req.Body,
		Image:      req.Image,
		CategoryID: req.CategoryID,
	}

	h.repo.CreateContent(&content)
	c.JSON(http.StatusOK, common.Success(gin.H{"id": content.ID}))
}

func (h *Handler) GetContent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, _ := h.repo.GetContent(uint(id))
	c.JSON(http.StatusOK, common.Success(item))
}

func (h *Handler) UpdateContent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, _ := h.repo.GetContent(uint(id))
	var req UpdateContentRequest
	c.ShouldBindJSON(&req)
	item.Title = req.Title
	item.Body = req.Body
	item.Image = req.Image
	h.repo.UpdateContent(item)
	c.JSON(http.StatusOK, common.Message("updated"))
}

func (h *Handler) DeleteContent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.repo.DeleteContent(uint(id))
	c.JSON(http.StatusOK, common.Message("deleted"))
}

/* ---------- CATEGORY ---------- */

func (h *Handler) ListCategories(c *gin.Context) {
	cats, _ := h.repo.ListCategories()
	c.JSON(http.StatusOK, common.Success(cats))
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	c.ShouldBindJSON(&req)
	h.repo.CreateCategory(&Category{Name: req.Name})
	c.JSON(http.StatusOK, common.Message("created"))
}

/* ---------- TAG ---------- */

func (h *Handler) ListTags(c *gin.Context) {
	tags, _ := h.repo.ListTags()
	c.JSON(http.StatusOK, common.Success(tags))
}

func (h *Handler) CreateTag(c *gin.Context) {
	var req CreateTagRequest
	c.ShouldBindJSON(&req)
	h.repo.CreateTag(&Tag{Name: req.Name})
	c.JSON(http.StatusOK, common.Message("created"))
}

/* ---------- MEDIA ---------- */

func (h *Handler) UploadMedia(c *gin.Context) {
	file, _ := c.FormFile("file")
	url := "/uploads/" + file.Filename

	media := Media{
		Type: "file",
		URL:  url,
	}
	h.repo.CreateMedia(&media)
	c.JSON(http.StatusOK, common.Success(gin.H{"id": media.ID, "media_url": url}))
}

func (h *Handler) ListMedia(c *gin.Context) {
	items, _ := h.repo.ListMedia()
	c.JSON(http.StatusOK, common.Success(items))
}

func (h *Handler) DeleteMedia(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.repo.DeleteMedia(uint(id))
	c.JSON(http.StatusOK, common.Message("deleted"))
}

func (h *Handler) AttachMedia(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req AttachMediaRequest
	c.ShouldBindJSON(&req)
	h.repo.AttachMedia(uint(id), req.Type, req.RefID)
	c.JSON(http.StatusOK, common.Message("attached"))
}
