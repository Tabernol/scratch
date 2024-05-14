package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"scratch/internal/database"
)

// APIConfig holds API-related configuration.
type APIConfig struct {
	DB *database.Queries
}

// Config represents the application configuration.
type Config struct {
	Port        string
	DatabaseURL string
	// Add more configuration fields as needed
}

// Load loads the configuration from environment variables.
func Load() (*Config, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT environment variable is not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL environment variable is not set")
	}

	return &Config{
		Port:        port,
		DatabaseURL: dbURL,
	}, nil
}
