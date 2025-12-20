package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

/* ===================== STRUCTS ===================== */

type AppConfig struct {
	Env  string
	Port string
}

type DBConfig struct {
	Dsn         string
	AutoMigrate bool
	Seed        bool
}

type JWTConfig struct {
	Secret    string
	AccessTTL time.Duration
}

type SMTPConfig struct {
	Host string
	Port int
	User string
	Pass string
	From string
}

type Config struct {
	App  AppConfig
	DB   DBConfig
	JWT  JWTConfig
	SMTP SMTPConfig
}

/* ===================== GLOBAL CONFIG ===================== */

// Cfg là biến global DUY NHẤT dùng trong toàn app
var Cfg *Config

/* ===================== LOAD ===================== */

// Init phải được gọi 1 lần trong main()
func Init() {
	Cfg = load()
}

// Load giữ lại để test / cli dùng nếu cần
func Load() *Config {
	return load()
}

func load() *Config {
	return &Config{
		App: AppConfig{
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		DB: DBConfig{
			Dsn:         mustEnv("DATABASE_URL"),
			AutoMigrate: getEnvBool("DB_AUTO_MIGRATE", true),
			Seed:        getEnvBool("DB_SEED", false),
		},
		JWT: JWTConfig{
			Secret:    mustEnv("JWT_SECRET"),
			AccessTTL: getEnvDuration("JWT_ACCESS_TTL", time.Hour*24),
		},
		SMTP: SMTPConfig{
			Host: getEnv("SMTP_HOST", ""),
			Port: getEnvInt("SMTP_PORT", 587),
			User: getEnv("SMTP_USER", ""),
			Pass: getEnv("SMTP_PASS", ""),
			From: getEnv("SMTP_FROM", ""),
		},
	}
}

/* ===================== HELPERS ===================== */

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("missing required env: %s", key)
	}
	return v
}

func getEnvBool(key string, def bool) bool {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return def
	}
	return b
}

func getEnvInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return i
}

func getEnvDuration(key string, def time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return def
	}
	return d
}
