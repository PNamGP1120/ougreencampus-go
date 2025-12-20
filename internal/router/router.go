package router

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/middleware"

	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	api := r.Group("/api")

	/* ================= USER ================= */
	userRepo := user.NewRepository(db)
	userSvc := user.NewService(userRepo, cfg.JWT.Secret)
	userHandler := user.NewHandler(userSvc, userRepo)
	user.RegisterRoutes(api, userHandler, cfg.JWT.Secret)

	/* ================= CONTENT ================= */
	contentRepo := content.NewRepository(db)
	contentHandler := content.NewHandler(contentRepo)
	content.RegisterRoutes(api, contentHandler, cfg.JWT.Secret)

	/* ================= ACTIVITY ================= */
	activityRepo := activity.NewRepository(db)
	activityHandler := activity.NewHandler(activityRepo)
	activity.RegisterRoutes(api, activityHandler, cfg.JWT.Secret)

	/* ================= EVENT ================= */
	eventRepo := event.NewRepository(db)
	eventHandler := event.NewHandler(eventRepo)
	event.RegisterRoutes(api, eventHandler, cfg.JWT.Secret)

	/* ================= SYSTEM ================= */
	systemRepo := system.NewRepository(db)
	systemSvc := system.NewService(systemRepo)
	systemHandler := system.NewHandler(systemRepo, systemSvc)
	system.RegisterRoutes(api, systemHandler, cfg.JWT.Secret)

	return r
}
