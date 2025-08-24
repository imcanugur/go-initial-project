package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	App struct {
		Port string
		Env  string
	}
	DB struct {
		Host     string
		Port     string
		User     string
		Pass     string
		Name     string
		SSLMode  string
		TimeZone string
	}
	JWT struct {
		Secret string
	}
}

var AppConfig *EnvConfig

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env file not found, using system env vars")
	}

	AppConfig = &EnvConfig{}

	AppConfig.App.Port = getEnv("APP_PORT", "8080")
	AppConfig.App.Env = getEnv("APP_ENV", "development")

	AppConfig.DB.Host = getEnv("DB_HOST", "localhost")
	AppConfig.DB.Port = getEnv("DB_PORT", "5432")
	AppConfig.DB.User = getEnv("DB_USER", "postgres")
	AppConfig.DB.Pass = getEnv("DB_PASS", "secret")
	AppConfig.DB.Name = getEnv("DB_NAME", "initial")
	AppConfig.DB.SSLMode = getEnv("DB_SSLMODE", "disable")
	AppConfig.DB.TimeZone = getEnv("DB_TIMEZONE", "Europe/Istanbul")

	AppConfig.JWT.Secret = getEnv("JWT_SECRET", "super-secret-key")
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
