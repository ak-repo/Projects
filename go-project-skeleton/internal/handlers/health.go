package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"go-project-skeleton/internal/response"
)

type HealthHandler struct {
	DB *sql.DB
}

func (h HealthHandler) Healthz(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h HealthHandler) Readyz(w http.ResponseWriter, r *http.Request) {
	if h.DB == nil {
		response.Error(w, http.StatusServiceUnavailable, "database not configured")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	if err := h.DB.PingContext(ctx); err != nil {
		response.Error(w, http.StatusServiceUnavailable, "database unavailable")
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"status": "ready"})
}
