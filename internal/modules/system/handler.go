package system

import (
	"net/http"
	"strconv"

	"github.com/PNamGP1120/ougreencampus-go/internal/common"
	"github.com/gin-gonic/gin"
)

/* ---------- CONFIG ---------- */

func (h *Handler) GetConfig(c *gin.Context) {
	items, _ := h.repo.GetConfigs()
	c.JSON(http.StatusOK, common.Success(items))
}

func (h *Handler) UpdateConfig(c *gin.Context) {
	var req UpdateConfigRequest
	c.ShouldBindJSON(&req)
	h.repo.UpsertConfig(req.Key, req.Value)
	c.JSON(http.StatusOK, common.Message("config updated"))
}

/* ---------- REPORT ---------- */

func (h *Handler) OverviewReport(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	c.JSON(http.StatusOK, common.Success(h.svc.OverviewReport(from, to)))
}

func (h *Handler) EventReport(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	c.JSON(http.StatusOK, common.Success(h.svc.EventReport(from, to)))
}

func (h *Handler) ActivityReport(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	c.JSON(http.StatusOK, common.Success(h.svc.ActivityReport(from, to)))
}

/* ---------- AUDIT ---------- */

func (h *Handler) AuditLogs(c *gin.Context) {
	items, _ := h.repo.ListAuditLogs()
	c.JSON(http.StatusOK, common.Success(items))
}

/* ---------- NOTIFICATION ---------- */

func (h *Handler) Notifications(c *gin.Context) {
	uid, _ := strconv.Atoi(c.GetString("user_id"))
	items, _ := h.repo.ListNotifications(uint(uid))
	c.JSON(http.StatusOK, common.Success(items))
}

func (h *Handler) MarkRead(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	h.repo.MarkRead(uint(id))
	c.JSON(http.StatusOK, common.Message("marked as read"))
}
