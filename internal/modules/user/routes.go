package user

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, cfg *config.Config) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service, cfg)

	// AUTH
	r.POST("/auth/register", handler.Register)
	r.POST("/auth/login", handler.Login)

	// USER
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(cfg))
	{
		auth.GET("/users/me", handler.Me)
		auth.PUT("/users/me", handler.UpdateMe)
		auth.PUT("/users/me/password", handler.ChangePassword)
	}

	admin := r.Group("/users")
	admin.Use(middleware.AuthMiddleware(cfg), middleware.RequireRoles("admin"))
	{
		admin.GET("", handler.List)
		admin.GET("/:id", handler.GetByID)
		admin.PATCH("/:id/role", handler.UpdateRole)
		admin.PATCH("/:id/status", handler.UpdateStatus)
	}
}
