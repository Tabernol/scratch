package database

import (
	"database/sql"
	"fmt"
	"scratch/config"
)

// NewDB establishes a connection to the database using the provided configuration
func NewDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	// Test the database connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging the database: %w", err)
	}
	return db, nil
}
