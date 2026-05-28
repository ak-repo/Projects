package handlers

import (
	"net/http"

	"go-project-skeleton/internal/response"
	"go-project-skeleton/internal/version"
)

func Version(appName, environment string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, map[string]string{
			"app":         appName,
			"environment": environment,
			"version":     version.Version,
			"commit":      version.Commit,
			"date":        version.Date,
		})
	}
}
