package main

import (
	"log"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/database"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/media"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/PNamGP1120/ougreencampus-go/internal/router"
)

func main() {
	// Load config (env from Docker / Railway)
	cfg := config.Load()

	// Connect DB
	db := database.Connect(cfg)

	// Auto migrate
	database.Migrate(
		db,

		// User
		&user.User{},

		// Media
		&media.Media{},

		// Content
		&content.Content{},
		&content.Category{},
		&content.Tag{},

		// Activity
		&activity.Activity{},
		&activity.ActivityParticipant{},
		&activity.ContestSubmission{},
		&activity.CampaignTask{},
		&activity.CampaignProgress{},
		&activity.ProgramRelation{},

		// Event
		&event.Event{},
		&event.EventRegistration{},

		// System
		&system.SystemConfig{},
		&system.AuditLog{},
		&system.Notification{},
	)

	// Seed data
	database.Seed(db)

	// Router
	r := router.Setup(db, cfg)

	log.Println("🚀 OU Green Campus API running on port", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
