package router

import (
	"net/http"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User     *user.Handler
	Auth     *user.AuthHandler
	Content  *content.Handler
	Event    *event.Handler
	Activity *activity.Handler
	System   *system.Handler
}

func SetupRouter(cfg *config.Config, jwtSvc *jwt.JWTService, h Handlers) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	// Health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Auth
	r.POST("/auth/login", h.Auth.Login)

	// Protected API
	api := r.Group("/api")
	api.Use(middleware.Auth(jwtSvc))

	// ===== USER (ADMIN) =====
	api.GET("/users",
		middleware.RequireRole(config.RoleAdmin),
		h.User.GetAllUsers,
	)
	api.POST("/users",
		middleware.RequireRole(config.RoleAdmin),
		h.User.CreateUser,
	)

	// ===== CONTENT =====
	api.GET("/contents", h.Content.GetAll)
	api.POST("/contents",
		middleware.RequireRole(config.RoleAdmin, config.RoleOrganizer),
		h.Content.Create,
	)

	// ===== EVENT =====
	api.GET("/events", h.Event.GetAll)
	api.POST("/events",
		middleware.RequireRole(config.RoleOrganizer, config.RoleAdmin),
		h.Event.Create,
	)
	api.POST("/events/:id/register",
		middleware.RequireRole(config.RoleStudent),
		h.Event.Register,
	)

	// ===== ACTIVITY (SIMPLIFIED) =====
	api.GET("/activities", h.Activity.List)
	api.POST("/activities",
		middleware.RequireRole(config.RoleOrganizer, config.RoleAdmin),
		h.Activity.Create,
	)
	api.POST("/activities/:id/contest/submissions",
		middleware.RequireRole(config.RoleStudent),
		h.Activity.SubmitContest,
	)
	api.GET("/activities/:id/contest/submissions",
		middleware.RequireRole(config.RoleOrganizer, config.RoleAdmin),
		h.Activity.Submissions,
	)

	// ===== SYSTEM (ADMIN) =====
	admin := api.Group("/admin")
	admin.Use(middleware.RequireRole(config.RoleAdmin))
	{
		admin.GET("/config", h.System.ListConfigs)
		admin.PUT("/config", h.System.UpsertConfig)
		admin.GET("/reports/overview", h.System.Overview)
		admin.GET("/audit", h.System.ListAudit)
	}

	return r
}
