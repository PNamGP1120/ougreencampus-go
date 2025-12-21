package system

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	// Admin only
	admin := r.Group("/")
	admin.Use(middleware.AuthMiddleware(cfg), middleware.RequireRoles("admin"))
	{
		admin.GET("/system/config", h.GetConfig)
		admin.PUT("/system/config", h.UpdateConfig)
		admin.GET("/reports/overview", h.Overview)
		admin.GET("/reports/events", h.EventReport)
		admin.GET("/reports/activities", h.ActivityReport)
		admin.GET("/audit/logs", h.Audit)
	}

	// Authenticated users
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(cfg))
	{
		auth.GET("/notifications", h.Notifications)
		auth.PATCH("/notifications/:id/read", h.ReadNotification)
	}
}
