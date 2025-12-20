package database

import (
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB, models ...interface{}) {
	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("auto migrate failed: %v", err)
	}
	log.Println("database migrated")
}
