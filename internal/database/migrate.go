package database

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB, models ...interface{}) {
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Fatal("migration failed:", err)
	}
	log.Println("database migrated")
}
