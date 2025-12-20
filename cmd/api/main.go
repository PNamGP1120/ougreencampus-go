package main

import (
	"log"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/database"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/media"
	"github.com/PNamGP1120/ougreencampus-go/internal/modules/user"
	"github.com/PNamGP1120/ougreencampus-go/internal/router"
)

func main() {
	// Load config
	cfg := config.Load()

	// Connect database
	db := database.Connect(cfg)

	// Auto migrate (PHASE 1)
	database.Migrate(
		db,
		&user.User{},
		&media.Media{},
	)

	// Seed initial data
	database.Seed(db)

	// Setup router
	r := router.Setup(db, cfg)

	log.Println("server started at port", cfg.AppPort)
	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
