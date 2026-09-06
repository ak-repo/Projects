package middleware

import (
	"crypto/subtle"
	"net/http"

	"go-project-skeleton/internal/response"
)

func BasicAuth(username, password string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			providedUsername, providedPassword, ok := r.BasicAuth()
			if !ok || !compare(username, providedUsername) || !compare(password, providedPassword) {
				w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
				response.Error(w, http.StatusUnauthorized, "unauthorized")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func compare(expected, actual string) bool {
	if expected == "" || actual == "" {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(expected), []byte(actual)) == 1
}
