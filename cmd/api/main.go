package main

import (
	"log"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/database"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/PNamGP1120/ougreencampus-go/internal/router"
)

func main() {
	cfg := config.Load()

	db := database.ConnectPostgres(cfg.DatabaseURL)

	// USER
	userRepo := user.NewRepository(db)
	userSvc := user.NewService(userRepo, cfg.JWTSecret)
	userHandler := user.NewHandler(userSvc)

	// CONTENT
	contentRepo := content.NewRepository(db)
	contentSvc := content.NewService(contentRepo)
	contentHandler := content.NewHandler(contentSvc)

	// ACTIVITY
	activityRepo := activity.NewRepository(db)
	activitySvc := activity.NewService(activityRepo)
	activityHandler := activity.NewHandler(activitySvc)

	// EVENT
	eventRepo := event.NewRepository(db)
	eventSvc := event.NewService(eventRepo)
	eventHandler := event.NewHandler(eventSvc)

	// SYSTEM
	systemRepo := system.NewRepository(db)
	systemSvc := system.NewService(systemRepo)
	systemHandler := system.NewHandler(systemSvc)

	r := router.SetupRouter(&router.Handlers{
		User:     userHandler,
		Content:  contentHandler,
		Activity: activityHandler,
		Event:    eventHandler,
		System:   systemHandler,
	}, cfg.JWTSecret)

	log.Println("🚀 API running on :" + cfg.Port)
	r.Run(":" + cfg.Port)
}
