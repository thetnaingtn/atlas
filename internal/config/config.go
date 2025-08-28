package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	// Server configuration
	Addr string
	Port string

	// Database configuration
	DatabaseURL string

	// Environment
	Environment string
}

// NewConfig creates a new configuration instance with default values
func NewConfig() *Config {
	return &Config{
		Addr:        getEnv("ADDR", "localhost"),
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		Environment: getEnv("ENVIRONMENT", "development"),
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
