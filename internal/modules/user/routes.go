package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
	}

	users := r.Group("/users")
	{
		users.GET("", h.ListUsers)
		users.POST("", h.CreateUser)
		users.GET("/:id", h.GetUser)
		users.PUT("/:id", h.UpdateUser)
		users.PATCH("/:id/role", h.ChangeRole)
		users.PATCH("/:id/status", h.ChangeStatus)
		users.PUT("/me/password", h.ChangePassword)
	}
}
