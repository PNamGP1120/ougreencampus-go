package system

import "github.com/gin-gonic/gin"

func RegisterRoutes(
	r *gin.RouterGroup,
	h *Handler,
	authMW gin.HandlerFunc,
	adminMW gin.HandlerFunc,
) {
	system := r.Group("/system", authMW, adminMW)
	{
		system.GET("/config", h.GetConfig)
		system.PUT("/config", h.UpdateConfig)
	}

	reports := r.Group("/reports", authMW, adminMW)
	{
		reports.GET("/overview", h.OverviewReport)
		reports.GET("/events", h.EventReport)
		reports.GET("/activities", h.ActivityReport)
	}

	r.GET("/audit/logs", authMW, adminMW, h.AuditLogs)

	r.GET("/notifications", authMW, h.Notifications)
	r.PATCH("/notifications/:id/read", authMW, h.MarkRead)
}
