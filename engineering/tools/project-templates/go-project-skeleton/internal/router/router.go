package router

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"go-project-skeleton/internal/handlers"
	appmiddleware "go-project-skeleton/internal/middleware"
	"go-project-skeleton/internal/response"
)

type Config struct {
	AppName           string
	Environment       string
	DB                *sql.DB
	Logger            *slog.Logger
	BasicAuthUsername string
	BasicAuthPassword string
}

func New(cfg Config) http.Handler {
	r := chi.NewRouter()

	r.Use(appmiddleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(appmiddleware.Recovery(cfg.Logger))
	r.Use(appmiddleware.Logging(cfg.Logger))

	health := handlers.HealthHandler{DB: cfg.DB}
	r.Get("/healthz", health.Healthz)
	r.Get("/readyz", health.Readyz)
	r.Get("/ping", handlers.Ping)
	r.Get("/version", handlers.Version(cfg.AppName, cfg.Environment))

	r.Route("/api/v1", func(r chi.Router) {
		r.Use(appmiddleware.BasicAuth(cfg.BasicAuthUsername, cfg.BasicAuthPassword))
		r.Get("/me", func(w http.ResponseWriter, r *http.Request) {
			response.JSON(w, http.StatusOK, map[string]string{"message": "authenticated"})
		})
	})

	return r
}
