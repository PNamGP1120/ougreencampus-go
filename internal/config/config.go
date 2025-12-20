package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort string

	DBUrl string

	JWTSecret string
	JWTTTL    int // minutes

	SMTPHost string
	SMTPPort string
	SMTPUser string
	SMTPPass string
	SMTPFrom string

	CloudinaryURL string
}

func Load() *Config {
	return &Config{
		AppPort: getEnv("APP_PORT", "8080"),

		DBUrl: getEnv("DATABASE_URL", ""),

		JWTSecret: getEnv("JWT_SECRET", "secret"),
		JWTTTL:    getEnvInt("JWT_TTL", 60),

		SMTPHost: getEnv("SMTP_HOST", ""),
		SMTPPort: getEnv("SMTP_PORT", ""),
		SMTPUser: getEnv("SMTP_USER", ""),
		SMTPPass: getEnv("SMTP_PASS", ""),
		SMTPFrom: getEnv("SMTP_FROM", ""),

		CloudinaryURL: getEnv("CLOUDINARY_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		var i int
		_, err := fmt.Sscan(v, &i)
		if err == nil {
			return i
		}
	}
	return fallback
}
