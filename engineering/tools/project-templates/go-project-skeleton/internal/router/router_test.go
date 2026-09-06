package router

import (
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPublicRoutes(t *testing.T) {
	r := New(Config{
		AppName:           "test-api",
		Environment:       "test",
		Logger:            slog.New(slog.NewTextHandler(io.Discard, nil)),
		BasicAuthUsername: "user",
		BasicAuthPassword: "pass",
	})

	tests := []struct {
		path string
		want string
	}{
		{path: "/healthz", want: `"status":"ok"`},
		{path: "/ping", want: `"message":"pong"`},
		{path: "/version", want: `"app":"test-api"`},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			rr := httptest.NewRecorder()

			r.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d", rr.Code, http.StatusOK)
			}
			if body := rr.Body.String(); !strings.Contains(body, tt.want) {
				t.Fatalf("body = %q, want to contain %q", body, tt.want)
			}
		})
	}
}

func TestProtectedRouteRequiresBasicAuth(t *testing.T) {
	r := New(Config{
		Logger:            slog.New(slog.NewTextHandler(io.Discard, nil)),
		BasicAuthUsername: "user",
		BasicAuthPassword: "pass",
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/me", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusUnauthorized)
	}
}

func TestProtectedRouteAllowsBasicAuth(t *testing.T) {
	r := New(Config{
		Logger:            slog.New(slog.NewTextHandler(io.Discard, nil)),
		BasicAuthUsername: "user",
		BasicAuthPassword: "pass",
	})

	req := httptest.NewRequest(http.MethodGet, "/api/v1/me", nil)
	req.SetBasicAuth("user", "pass")
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rr.Code, http.StatusOK)
	}
}
