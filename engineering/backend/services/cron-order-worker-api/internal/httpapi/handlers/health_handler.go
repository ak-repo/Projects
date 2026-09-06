package handlers

import (
	"net/http"

	"cron-order-worker-api/internal/httpapi/response"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, http.StatusOK, map[string]any{
		"status":  "ok",
		"message": "service is healthy",
	})
}
