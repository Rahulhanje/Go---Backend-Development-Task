package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	ServerPort  string
}

func Load() *Config {
	// Load .env file if exists
	_ = godotenv.Load()

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "postgres"),
		DBName:      getEnv("DB_NAME", "userdb"),
		ServerPort:  getEnv("SERVER_PORT", "3000"),
	}
}

func (c *Config) GetDBConnectionString() string {
	// Use DATABASE_URL if provided (for cloud databases like Aiven)
	if c.DatabaseURL != "" {
		return c.DatabaseURL
	}
	// Fallback to individual params
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
