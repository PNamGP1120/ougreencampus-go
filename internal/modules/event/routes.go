package event

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

	r.GET("/events", h.ListEvents)
	r.GET("/events/:id", h.GetEvent)

	auth := r.Group("", middleware.AuthMiddleware(jwtSecret))
	{
		auth.POST("/events/:id/register", h.Register)
	}

	org := r.Group("", middleware.AuthMiddleware(jwtSecret), middleware.RequireRole("organizer"))
	{
		org.POST("/events", h.CreateEvent)
		org.PUT("/events/:id", h.UpdateEvent)
		org.DELETE("/events/:id", h.DeleteEvent)
		org.GET("/events/:id/registrations", h.ListRegistrations)
		org.POST("/events/:id/send-confirmation", h.SendConfirmation)
		org.POST("/events/:id/checkin", h.CheckinQR)
		org.POST("/events/:id/checkin/manual", h.CheckinManual)
		org.GET("/events/:id/stats", h.Stats)
		org.GET("/events/:id/export", h.Export)
	}
}
