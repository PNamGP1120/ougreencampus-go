package user

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup, h *Handler) {

	auth := r.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.POST("/forgot-password", h.ForgotPassword)
		auth.POST("/reset-password", h.ResetPassword)
	}

	authSec := r.Group("/auth")
	authSec.Use(middleware.AuthMiddleware())
	{
		authSec.POST("/logout", h.Logout)
		authSec.GET("/me", h.Me)
		authSec.POST("/refresh", h.Refresh)
	}

	admin := r.Group("/users")
	admin.Use(
		middleware.AuthMiddleware(),
		middleware.RequireRoles("admin"),
	)
	{
		admin.GET("", h.ListUsers)
		admin.POST("", h.CreateUser)
		admin.GET("/:id", h.GetUser)
		admin.PUT("/:id", h.UpdateUser)
		admin.PATCH("/:id/role", h.ChangeRole)
		admin.PATCH("/:id/status", h.ChangeStatus)
	}

	me := r.Group("/users/me")
	me.Use(middleware.AuthMiddleware())
	{
		me.GET("", h.GetMe)
		me.PUT("", h.UpdateMe)
		me.PUT("/password", h.ChangePassword)
	}
}
