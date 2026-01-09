package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		Server: ServerConfig{
			Name: getEnv("APP_NAME", "golang-todo"),
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8888"),
		},
		Database: DbConfig{
			User:     getEnv("DB_USER", "db"),
			Host:     getEnv("DB_HOST", "db"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", "db"),
			Password: getEnv("DB_PASSWORD", "db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Security: SecurityConfig{
			JWTSecret: mustEnv("JWT_SECRET"),
		},
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func mustEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("Missing required environment variable: %s", key)
	}
	return value
}
