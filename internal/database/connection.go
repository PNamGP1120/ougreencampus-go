package database

import (
	"fmt"
	"strings"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.DatabaseURL
	if dsn != "" {
		// đảm bảo sslmode=require nếu Railway không kèm
		if !strings.Contains(dsn, "sslmode=") {
			if strings.Contains(dsn, "?") {
				dsn += "&sslmode=require"
			} else {
				dsn += "?sslmode=require"
			}
		}
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	// fallback local/docker
	dsn = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
