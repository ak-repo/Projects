package handlers

import (
	"net/http"

	"go-project-skeleton/internal/response"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{"message": "pong"})
}
