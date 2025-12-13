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
	"github.com/PNamGP1120/ougreencampus-go/pkg/jwt"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("❌ Load config failed:", err)
	}
	log.Println("✅ Config loaded")

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}
	log.Println("✅ Database connected")

	if err := database.AutoMigrate(db); err != nil {
		log.Fatal("❌ Auto migrate failed:", err)
	}
	log.Println("✅ Database migrated")

	if cfg.AppEnv != "production" {
		if err := database.Seed(db); err != nil {
			log.Fatal("❌ Seed failed:", err)
		}
		log.Println("🌱 Database seeded")
	}

	jwtSvc := jwt.New(cfg.JWTSecret, cfg.JWTExpire)

	// User
	userRepo := user.NewRepository(db)
	userSvc := user.NewService(userRepo)
	authSvc := user.NewAuthService(userRepo, jwtSvc)
	userHandler := user.NewHandler(userSvc)
	authHandler := user.NewAuthHandler(authSvc)

	// Content
	contentRepo := content.NewRepository(db)
	contentSvc := content.NewService(contentRepo)
	contentHandler := content.NewHandler(contentSvc)

	// Event
	eventRepo := event.NewRepository(db)
	eventSvc := event.NewService(eventRepo)
	eventHandler := event.NewHandler(eventSvc)

	// Activity
	activityRepo := activity.NewRepository(db)
	activitySvc := activity.NewService(activityRepo)
	activityHandler := activity.NewHandler(activitySvc)

	// System
	systemRepo := system.NewRepository(db)
	systemSvc := system.NewService(systemRepo)
	systemHandler := system.NewHandler(systemSvc)

	r := router.SetupRouter(cfg, jwtSvc, router.Handlers{
		User:     userHandler,
		Auth:     authHandler,
		Content:  contentHandler,
		Event:    eventHandler,
		Activity: activityHandler,
		System:   systemHandler,
	})

	log.Println("🚀 Server running on port", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}
