package system

import (
	"strconv"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{svc: s}
}

// 59 GET /system/config
func (h *Handler) GetConfig(c *gin.Context) {
	items, _ := h.svc.GetConfig()
	common.SuccessResponse(c, items)
}

// 60 PUT /system/config
func (h *Handler) UpdateConfig(c *gin.Context) {
	var req UpdateConfigRequest
	c.ShouldBindJSON(&req)
	h.svc.UpdateConfig(req.Key, req.Value)
	common.SuccessResponse(c, gin.H{"message": "updated"})
}

// 61 GET /reports/overview
func (h *Handler) Overview(c *gin.Context) {
	data, _ := h.svc.Overview()
	common.SuccessResponse(c, data)
}

// 62 GET /reports/events
func (h *Handler) EventReport(c *gin.Context) {
	data, _ := h.svc.EventReport()
	common.SuccessResponse(c, data)
}

// 63 GET /reports/activities
func (h *Handler) ActivityReport(c *gin.Context) {
	data, _ := h.svc.ActivityReport()
	common.SuccessResponse(c, data)
}

// 64 GET /audit/logs
func (h *Handler) Audit(c *gin.Context) {
	page, limit := common.GetPagination(c)

	var uid *uint
	if v := c.Query("user_id"); v != "" {
		id, _ := strconv.Atoi(v)
		u := uint(id)
		uid = &u
	}

	items, total, _ := h.svc.AuditLogs(uid, c.Query("action"), page, limit)
	common.SuccessResponse(c, gin.H{
		"items": items,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// 65 GET /notifications
func (h *Handler) Notifications(c *gin.Context) {
	page, limit := common.GetPagination(c)
	items, total, _ := h.svc.Notifications(c.GetUint("user_id"), page, limit)

	common.SuccessResponse(c, gin.H{
		"items": items,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
		},
	})
}

// 66 PATCH /notifications/:id/read
func (h *Handler) ReadNotification(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.svc.ReadNotification(uint(id), c.GetUint("user_id"))
	common.SuccessResponse(c, gin.H{"message": "read"})
}
