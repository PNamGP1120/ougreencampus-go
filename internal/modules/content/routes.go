package content

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	// Public
	r.GET("/contents", handler.List)
	r.GET("/contents/:id", handler.Get)
	r.GET("/categories", handler.Categories)
	r.GET("/tags", handler.Tags)

	// Organizer
	org := r.Group("/")
	org.Use(middleware.AuthMiddleware(cfg), middleware.RequireRoles("organizer", "admin"))
	{
		org.POST("/contents", handler.Create)
		org.PUT("/contents/:id", handler.Update)
		org.DELETE("/contents/:id", handler.Delete)
	}

	// Admin
	admin := r.Group("/")
	admin.Use(middleware.AuthMiddleware(cfg), middleware.RequireRoles("admin"))
	{
		admin.POST("/categories", handler.CreateCategory)
		admin.POST("/tags", handler.CreateTag)
	}
}
