package main

import (
	"log"
	"os"

	"github.com/PNamGP1120/ougreencampus-go/internal/database"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/activity"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/content"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/event"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/system"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db := database.ConnectPostgres(dsn)

	// ✅ BẮT BUỘC cho UUID v4
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	log.Println("📦 Running migrations...")

	err := db.AutoMigrate(
		&user.User{},
		&content.Category{},
		&content.Tag{},
		&content.Content{},
		&activity.Activity{},
		&event.Event{},
		&system.SystemConfig{},
		&system.AuditLog{},
		&system.Notification{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Migration completed")
}
