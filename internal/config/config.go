package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string

	// ưu tiên cho Railway
	DatabaseURL string

	// fallback cho local/docker
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret string
	JWTExpire int
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	expire, _ := strconv.Atoi(getEnv("JWT_EXPIRE_MINUTES", "60"))

	cfg := &Config{
		AppEnv:  getEnv("APP_ENV", "development"),
		AppPort: getEnv("APP_PORT", "8080"),

		DatabaseURL: os.Getenv("DATABASE_URL"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "ougc"),
		DBPassword: getEnv("DB_PASSWORD", "ougc"),
		DBName:     getEnv("DB_NAME", "ougreencampus"),

		JWTSecret: getEnv("JWT_SECRET", "secret"),
		JWTExpire: expire,
	}

	// Fail fast: production phải có DATABASE_URL (Railway chuẩn)
	if cfg.AppEnv == "production" && cfg.DatabaseURL == "" {
		return nil, errors.New("DATABASE_URL is required in production (Railway)")
	}

	return cfg, nil
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
