package db

import (
	"context"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"

	"go-project-skeleton/internal/config"
)

func Open(ctx context.Context, cfg config.DatabaseConfig) (*sql.DB, error) {
	database, err := sql.Open("pgx", cfg.URL)
	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(cfg.MaxOpenConns)
	database.SetMaxIdleConns(cfg.MaxIdleConns)
	database.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err := database.PingContext(ctx); err != nil {
		database.Close()
		return nil, err
	}

	return database, nil
}
