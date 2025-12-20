package content

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

	r.GET("/contents", h.ListContents)
	r.GET("/contents/:id", h.GetContent)

	r.GET("/categories", h.ListCategories)
	r.GET("/tags", h.ListTags)

	auth := r.Group("", middleware.AuthMiddleware(jwtSecret))
	{
		auth.GET("/media", h.ListMedia)
	}

	org := r.Group("", middleware.AuthMiddleware(jwtSecret), middleware.RequireRole("organizer"))
	{
		org.POST("/contents", h.CreateContent)
		org.PUT("/contents/:id", h.UpdateContent)
		org.DELETE("/contents/:id", h.DeleteContent)
		org.POST("/media/upload", h.UploadMedia)
		org.POST("/media/:id/attach", h.AttachMedia)
	}

	admin := r.Group("", middleware.AuthMiddleware(jwtSecret), middleware.RequireRole("admin"))
	{
		admin.POST("/categories", h.CreateCategory)
		admin.POST("/tags", h.CreateTag)
		admin.DELETE("/media/:id", h.DeleteMedia)
	}
}
