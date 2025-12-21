package activity

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
	r.GET("/activities", h.List)
	r.GET("/activities/:id", h.Get)
	r.GET("/activities/:id/children", h.Children)

	// Student
	st := r.Group("/")
	st.Use(middleware.AuthMiddleware(cfg), middleware.RequireRoles("student"))
	{
		st.POST("/activities/:id/join", h.Join)
		st.DELETE("/activities/:id/leave", h.Leave)
		st.POST("/activities/:id/submissions", h.Submit)
		st.POST("/tasks/:id/progress", h.Progress)
	}

	// Organizer
	org := r.Group("/")
	org.Use(middleware.AuthMiddleware(cfg), middleware.RequireRoles("organizer", "admin"))
	{
		org.POST("/activities", h.Create)
		org.PUT("/activities/:id", h.Update)
		org.DELETE("/activities/:id", h.Delete)
		org.GET("/activities/:id/participants", h.Participants)
		org.GET("/activities/:id/submissions", h.Submissions)
		org.PATCH("/submissions/:id/review", h.Review)
		org.POST("/activities/:id/tasks", h.CreateTask)
		org.GET("/activities/:id/tasks", h.Tasks)
		org.GET("/activities/:id/metrics", h.Metrics)
		org.POST("/activities/:id/children", h.AddChild)
	}
}
