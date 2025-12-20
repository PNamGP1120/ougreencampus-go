package system

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
	svc  *Service
}

func NewHandler(r *Repository, s *Service) *Handler {
	return &Handler{r, s}
}

func RegisterRoutes(r *gin.RouterGroup, h *Handler, jwtSecret string) {

	admin := r.Group("", middleware.AuthMiddleware(jwtSecret), middleware.RequireRole("admin"))
	{
		admin.GET("/system/config", h.GetConfig)
		admin.PUT("/system/config", h.UpdateConfig)

		admin.GET("/reports/overview", h.OverviewReport)
		admin.GET("/reports/events", h.EventReport)
		admin.GET("/reports/activities", h.ActivityReport)

		admin.GET("/audit/logs", h.AuditLogs)
	}

	auth := r.Group("", middleware.AuthMiddleware(jwtSecret))
	{
		auth.GET("/notifications", h.Notifications)
		auth.PATCH("/notifications/:id/read", h.MarkRead)
	}
}
