package activity

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
}

func NewHandler(r *Repository) *Handler {
	return &Handler{r}
}

func RegisterRoutes(r *gin.RouterGroup, h *Handler, jwtSecret string) {

	r.GET("/activities", h.ListActivities)
	r.GET("/activities/:id", h.GetActivity)
	r.GET("/activities/:id/children", h.ListChildren)

	auth := r.Group("", middleware.AuthMiddleware(jwtSecret))
	{
		auth.POST("/activities/:id/join", h.JoinActivity)
		auth.DELETE("/activities/:id/leave", h.LeaveActivity)
		auth.GET("/activities/:id/tasks", h.ListTasks)
	}

	student := r.Group("", middleware.AuthMiddleware(jwtSecret), middleware.RequireRole("student"))
	{
		student.POST("/activities/:id/submissions", h.Submit)
		student.POST("/tasks/:id/progress", h.AddProgress)
	}

	org := r.Group("", middleware.AuthMiddleware(jwtSecret), middleware.RequireRole("organizer"))
	{
		org.POST("/activities", h.CreateActivity)
		org.PUT("/activities/:id", h.UpdateActivity)
		org.DELETE("/activities/:id", h.DeleteActivity)
		org.GET("/activities/:id/participants", h.ListParticipants)
		org.GET("/activities/:id/submissions", h.ListSubmissions)
		org.PATCH("/submissions/:id/review", h.ReviewSubmission)
		org.POST("/activities/:id/tasks", h.CreateTask)
		org.GET("/activities/:id/metrics", func(c *gin.Context) {
			c.JSON(200, gin.H{"metrics": "not implemented yet"})
		})
		org.POST("/activities/:id/children", h.AddChild)
	}
}
