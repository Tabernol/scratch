package database

import (
	"database/sql"
	"scratch/config"
)

// NewDB establishes a connection to the database using the provided configuration
func NewDB(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
