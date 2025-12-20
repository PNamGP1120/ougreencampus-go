package main

import (
	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"github.com/PNamGP1120/ougreencampus-go/internal/database"
	"github.com/PNamGP1120/ougreencampus-go/internal/router"
)

func main() {
	// 🔥 BẮT BUỘC
	config.Init()

	db := database.Connect(config.Cfg.DB)

	if config.Cfg.DB.AutoMigrate {
		database.AutoMigrate(db)
	}
	if config.Cfg.DB.Seed {
		database.Seed(db)
	}

	r := router.New(db)
	r.Run(":" + config.Cfg.App.Port)
}
