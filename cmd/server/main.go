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
	// ================= LOAD CONFIG =================
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("❌ Load config failed:", err)
	}

	log.Println("✅ Config loaded")

	// ================= CONNECT DATABASE =================
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("❌ Database connection failed:", err)
	}

	log.Println("✅ Database connected")

	// ================= AUTO MIGRATE =================
	if err := database.AutoMigrate(
		db,
		&user.User{},
		&content.Content{},
		&event.Event{},
		&event.EventRegistration{},
		&activity.Activity{},
		&activity.Submission{},
		&system.SystemConfig{},
		&system.AuditLog{},
	); err != nil {
		log.Fatal("❌ Auto migrate failed:", err)
	}

	log.Println("✅ Database migrated")

	// ================= SEED DATA =================
	if err := database.Seed(db); err != nil {
		log.Fatal("❌ Seed data failed:", err)
	}

	log.Println("✅ Database seeded")

	// ================= JWT SERVICE =================
	jwtSvc := jwt.NewJWTService(cfg.JWTSecret, cfg.JWTExpire)

	// ================= ROUTER =================
	r := router.SetupRouter(db, jwtSvc)

	log.Println("🚀 Server running on port:", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("❌ Server failed:", err)
	}
}
