package config

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.GetDBConnectionString())
	if err != nil {
		return nil, err
	}

	// Test connection
	if err := db.PingContext(context.Background()); err != nil {
		return nil, err
	}

	return db, nil
}
