package app

import (
	"context"
	"database/sql"
	"log/slog"

	"go-project-skeleton/internal/config"
	"go-project-skeleton/internal/httpserver"
	"go-project-skeleton/internal/platform/db"
	"go-project-skeleton/internal/router"
)

type App struct {
	DB     *sql.DB
	Server *httpserver.Server
}

func New(ctx context.Context, cfg config.Config, logger *slog.Logger) (*App, error) {
	database, err := db.Open(ctx, cfg.Database)
	if err != nil {
		return nil, err
	}

	r := router.New(router.Config{
		AppName:           cfg.AppName,
		Environment:       cfg.Environment,
		DB:                database,
		Logger:            logger,
		BasicAuthUsername: cfg.BasicAuth.Username,
		BasicAuthPassword: cfg.BasicAuth.Password,
	})

	server := httpserver.New(cfg.HTTPAddr, r, cfg.HTTPServer)

	return &App{DB: database, Server: server}, nil
}

func (a *App) Close() {
	if a.DB != nil {
		a.DB.Close()
	}
}
