package database

import (
	"log"
	"time"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connect(cfg config.DBConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = gorm.Open(postgres.Open(cfg.Dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Warn),
		})

		if err == nil {
			sqlDB, err2 := db.DB()
			if err2 == nil {
				sqlDB.SetMaxIdleConns(10)
				sqlDB.SetMaxOpenConns(50)
				sqlDB.SetConnMaxLifetime(time.Hour)
				log.Println("PostgreSQL connected")
				return db, nil
			}
		}

		log.Printf("Waiting for database... (%d/10)\n", i)
		time.Sleep(2 * time.Second)
	}

	return nil, err
}
