package database

import (
	"fmt"
	"os"
	"strings"

	"github.com/PNamGP1120/ougreencampus-go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")

	if dsn != "" {
		// log an toàn, không lộ password
		masked := dsn
		if i := strings.Index(masked, "@"); i > 0 {
			masked = "postgresql://***:***" + masked[i:]
		}
		fmt.Println("✅ Using DATABASE_URL:", masked)
	} else {
		fmt.Println("⚠️ DATABASE_URL is empty; falling back to DB_HOST/DB_PORT")
		dsn = fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBName,
		)
	}

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
