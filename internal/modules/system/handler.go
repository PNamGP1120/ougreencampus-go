package system

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

/* ========= CONFIG ========= */

func (h *Handler) GetConfig(c *gin.Context) {
	data, _ := h.svc.GetConfigs()
	c.JSON(200, data)
}

func (h *Handler) UpdateConfig(c *gin.Context) {
	var req UpdateConfigRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_ = h.svc.UpdateConfig(req.Key, req.Value)
	c.JSON(200, gin.H{"message": "updated"})
}

/* ========= REPORT ========= */

func (h *Handler) OverviewReport(c *gin.Context) {
	c.JSON(200, h.svc.OverviewReport())
}

func (h *Handler) EventReport(c *gin.Context) {
	c.JSON(200, h.svc.EventReport())
}

func (h *Handler) ActivityReport(c *gin.Context) {
	c.JSON(200, h.svc.ActivityReport())
}

/* ========= AUDIT ========= */

func (h *Handler) AuditLogs(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	items, total := h.svc.AuditLogs(AuditFilter{
		UserID: c.Query("user_id"),
		Action: c.Query("action"),
		Page:   page,
		Limit:  limit,
	})

	c.JSON(200, gin.H{
		"items": items,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

/* ========= NOTIFICATION ========= */

func (h *Handler) Notifications(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))

	items, total := h.svc.Notifications(
		c.GetString("user_id"),
		NotificationFilter{Page: page, Limit: limit},
	)

	c.JSON(200, gin.H{
		"items": items,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *Handler) MarkRead(c *gin.Context) {
	_ = h.svc.MarkRead(c.Param("id"))
	c.JSON(200, gin.H{"message": "read"})
}
