package config

import (
	"os"
)

type Config struct {
	Port        string
	DatabaseURL string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", ":8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""), // If no set it to empty string to trigger no-database in main
	}
}

// Helper function: If variables dont exist in env assign the local variables for testing
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		if key == "PORT" && value[0] != ':' {
			return ":" + value
		}
		return value
	}
	return fallback
}