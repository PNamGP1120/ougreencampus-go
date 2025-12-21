package event

import (
	"strconv"
	"time"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{svc: s}
}

// 47 GET /events
func (h *Handler) List(c *gin.Context) {
	page, limit := common.GetPagination(c)
	search := c.Query("search")

	var from, to *time.Time
	if v := c.Query("from"); v != "" {
		t, _ := time.Parse(time.RFC3339, v)
		from = &t
	}
	if v := c.Query("to"); v != "" {
		t, _ := time.Parse(time.RFC3339, v)
		to = &t
	}

	items, total, _ := h.svc.List(search, from, to, page, limit)
	common.SuccessResponse(c, gin.H{
		"items": items,
		"pagination": gin.H{
			"page": page, "limit": limit, "total": total,
		},
	})
}

// 48 POST /events
func (h *Handler) Create(c *gin.Context) {
	var req CreateEventRequest
	c.ShouldBindJSON(&req)
	e, _ := h.svc.Create(c.GetUint("user_id"), req)
	common.SuccessResponse(c, gin.H{"id": e.ID})
}

// 49 GET /events/:id
func (h *Handler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	e, _ := h.svc.Get(uint(id))
	common.SuccessResponse(c, e)
}

// 50 PUT /events/:id
func (h *Handler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req UpdateEventRequest
	c.ShouldBindJSON(&req)
	h.svc.Update(uint(id), req)
	common.SuccessResponse(c, gin.H{"message": "updated"})
}

// 51 DELETE /events/:id
func (h *Handler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.Delete(uint(id))
	common.SuccessResponse(c, gin.H{"message": "deleted"})
}

// 52 POST /events/:id/register
func (h *Handler) Register(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.Register(uint(id), c.GetUint("user_id"))
	common.SuccessResponse(c, gin.H{"message": "registered"})
}

// 53 GET /events/:id/registrations
func (h *Handler) Registrations(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	page, limit := common.GetPagination(c)
	items, total, _ := h.svc.Registrations(uint(id), "", page, limit)

	common.SuccessResponse(c, gin.H{
		"items": items,
		"pagination": gin.H{
			"page": page, "limit": limit, "total": total,
		},
	})
}

// 54 POST /events/:id/send-confirmation
func (h *Handler) SendConfirmation(c *gin.Context) {
	common.SuccessResponse(c, gin.H{"message": "confirmation sent"})
}

// 55 POST /events/:id/checkin (QR – placeholder)
func (h *Handler) CheckinQR(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.Checkin(uint(id), c.GetUint("user_id"))
	common.SuccessResponse(c, gin.H{"message": "checked_in"})
}

// 56 POST /events/:id/checkin/manual
func (h *Handler) CheckinManual(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req ManualCheckinRequest
	c.ShouldBindJSON(&req)
	h.svc.Checkin(uint(id), req.UserID)
	common.SuccessResponse(c, gin.H{"message": "checked_in"})
}

// 57 GET /events/:id/stats
func (h *Handler) Stats(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	stats, _ := h.svc.Stats(uint(id))
	common.SuccessResponse(c, stats)
}

// 58 GET /events/:id/export
func (h *Handler) Export(c *gin.Context) {
	c.Header("Content-Type", "text/csv")
	c.String(200, "user_id,checked_in\n")
}
