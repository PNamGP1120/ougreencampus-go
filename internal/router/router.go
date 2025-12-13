package router

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(
	db *gorm.DB,
	jwtSvc *jwt.JWTService,
) *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())

	Health(r)

	// ===== INIT REPOS =====
	userRepo := user.NewRepository(db)
	contentRepo := content.NewRepository(db)
	eventRepo := event.NewRepository(db)
	activityRepo := activity.NewRepository(db)

	// ===== INIT SERVICES =====
	userSvc := user.NewService(userRepo)
	authSvc := user.NewAuthService(userRepo, jwtSvc)

	contentSvc := content.NewService(contentRepo)
	eventSvc := event.NewService(eventRepo)
	activitySvc := activity.NewService(activityRepo)

	// ===== INIT HANDLERS =====
	userHandler := user.NewHandler(userSvc)
	authHandler := user.NewAuthHandler(authSvc)

	contentHandler := content.NewHandler(contentSvc)
	eventHandler := event.NewHandler(eventSvc)
	activityHandler := activity.NewHandler(activitySvc)

	// ===== PUBLIC =====
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	// ===== AUTH =====
	api := r.Group("/api")
	api.Use(middleware.Auth(jwtSvc))

	// ----- USER (ADMIN) -----
	api.POST("/users", middleware.RequireRole(config.RoleAdmin), userHandler.CreateUser)
	api.GET("/users", middleware.RequireRole(config.RoleAdmin), userHandler.GetAllUsers)

	// ----- CONTENT -----
	api.POST("/contents",
		middleware.RequireRole(config.RoleOrganizer, config.RoleAdmin),
		contentHandler.Create,
	)
	api.GET("/contents", contentHandler.GetAll)

	// ----- EVENT -----
	api.POST("/events",
		middleware.RequireRole(config.RoleOrganizer),
		eventHandler.Create,
	)
	api.GET("/events", eventHandler.GetAll)
	api.POST("/events/:id/register",
		middleware.RequireRole(config.RoleStudent),
		eventHandler.Register,
	)

	// ----- ACTIVITY -----
	api.POST("/activities",
		middleware.RequireRole(config.RoleOrganizer),
		activityHandler.Create,
	)
	api.GET("/activities", activityHandler.GetAll)
	api.POST("/activities/:id/submit",
		middleware.RequireRole(config.RoleStudent),
		activityHandler.Submit,
	)

	// ----- SYSTEM (ADMIN) -----
	_ = system.NewReportService(db)

	return r
}
