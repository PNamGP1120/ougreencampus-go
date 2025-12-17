package event

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

/* ========= EVENT ========= */

func (h *Handler) ListEvents(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	filter := ListEventFilter{
		Search: c.Query("search"),
		From:   c.Query("from"),
		To:     c.Query("to"),
		Page:   page,
		Limit:  limit,
	}

	items, total := h.svc.ListEvents(filter)

	c.JSON(200, gin.H{
		"items": items,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) GetEvent(c *gin.Context) {
	data, _ := h.svc.GetEvent(c.Param("id"))
	c.JSON(200, data)
}

func (h *Handler) CreateEvent(c *gin.Context) {
	var req CreateEventRequest
	_ = c.ShouldBindJSON(&req)
	id, _ := h.svc.CreateEvent(req, c.GetString("user_id"))
	c.JSON(200, gin.H{"id": id})
}

func (h *Handler) UpdateEvent(c *gin.Context) {
	var req UpdateEventRequest
	_ = c.ShouldBindJSON(&req)
	_ = h.svc.UpdateEvent(c.Param("id"), req)
	c.JSON(200, gin.H{"message": "updated"})
}

func (h *Handler) DeleteEvent(c *gin.Context) {
	_ = h.svc.DeleteEvent(c.Param("id"))
	c.JSON(200, gin.H{"message": "deleted"})
}

/* ========= REGISTRATION ========= */

func (h *Handler) Register(c *gin.Context) {
	_ = h.svc.Register(c.Param("id"), c.GetString("user_id"))
	c.JSON(200, gin.H{"message": "registered"})
}

func (h *Handler) Registrations(c *gin.Context) {
	data, _ := h.svc.Registrations(c.Param("id"))
	c.JSON(200, data)
}

func (h *Handler) CheckinQR(c *gin.Context) {
	var req CheckinQRRequest
	_ = c.ShouldBindJSON(&req)
	_ = h.svc.Checkin(c.Param("id"), req.QR)
	c.JSON(200, gin.H{"message": "checked in"})
}

func (h *Handler) CheckinManual(c *gin.Context) {
	var req ManualCheckinRequest
	_ = c.ShouldBindJSON(&req)
	_ = h.svc.Checkin(c.Param("id"), req.UserID)
	c.JSON(200, gin.H{"message": "checked in"})
}

func (h *Handler) Stats(c *gin.Context) {
	regs, _ := h.svc.Registrations(c.Param("id"))
	total := len(regs)
	checked := 0
	for _, r := range regs {
		if r.CheckedIn {
			checked++
		}
	}
	c.JSON(200, gin.H{
		"registered": total,
		"checked_in": checked,
	})
}

func (h *Handler) Export(c *gin.Context) {
	c.JSON(200, gin.H{"message": "exported"})
}
