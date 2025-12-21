package event

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

	// Public
	r.GET("/events", h.List)
	r.GET("/events/:id", h.Get)

	// Student
	st := r.Group("/")
	st.Use(middleware.AuthMiddleware(cfg), middleware.RequireRoles("student"))
	{
		st.POST("/events/:id/register", h.Register)
	}

	// Organizer
	org := r.Group("/")
	org.Use(middleware.AuthMiddleware(cfg), middleware.RequireRoles("organizer", "admin"))
	{
		org.POST("/events", h.Create)
		org.PUT("/events/:id", h.Update)
		org.DELETE("/events/:id", h.Delete)
		org.GET("/events/:id/registrations", h.Registrations)
		org.POST("/events/:id/send-confirmation", h.SendConfirmation)
		org.POST("/events/:id/checkin", h.CheckinQR)
		org.POST("/events/:id/checkin/manual", h.CheckinManual)
		org.GET("/events/:id/stats", h.Stats)
		org.GET("/events/:id/export", h.Export)
	}
}
