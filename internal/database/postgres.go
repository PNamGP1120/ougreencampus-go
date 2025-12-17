package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(dsn string) *gorm.DB {
	if dsn == "" {
		dsn = os.Getenv("DATABASE_URL")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect database:", err)
	}

	return db
}
