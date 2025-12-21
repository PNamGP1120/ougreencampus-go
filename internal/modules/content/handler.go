package content

import (
	"net/http"
	"strconv"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) *Handler {
	return &Handler{s}
}

func (h *Handler) List(c *gin.Context) {
	page, limit := common.GetPagination(c)

	filter := map[string]interface{}{}
	if cid := c.Query("category_id"); cid != "" {
		filter["category_id"] = cid
	}

	items, total, _ := h.service.List(filter, page, limit)
	common.SuccessResponse(c, gin.H{
		"items": items,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

func (h *Handler) Create(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req CreateContentRequest
	c.ShouldBindJSON(&req)

	content, err := h.service.Create(userID, req)
	if err != nil {
		common.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SuccessResponse(c, gin.H{"id": content.ID})
}

func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	item, _ := h.service.Get(uint(id))
	common.SuccessResponse(c, item)
}

func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req UpdateContentRequest
	c.ShouldBindJSON(&req)
	h.service.Update(uint(id), req)
	common.SuccessResponse(c, gin.H{"message": "updated"})
}

func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.service.Delete(uint(id))
	common.SuccessResponse(c, gin.H{"message": "deleted"})
}

func (h *Handler) Categories(c *gin.Context) {
	items, _ := h.service.Categories()
	common.SuccessResponse(c, items)
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var req struct{ Name string }
	c.ShouldBindJSON(&req)
	cat, _ := h.service.CreateCategory(req.Name)
	common.SuccessResponse(c, gin.H{"id": cat.ID})
}

func (h *Handler) Tags(c *gin.Context) {
	items, _ := h.service.Tags()
	common.SuccessResponse(c, items)
}

func (h *Handler) CreateTag(c *gin.Context) {
	var req struct{ Name string }
	c.ShouldBindJSON(&req)
	tag, _ := h.service.CreateTag(req.Name)
	common.SuccessResponse(c, gin.H{"id": tag.ID})
}
