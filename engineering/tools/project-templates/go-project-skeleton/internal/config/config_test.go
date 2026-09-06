package config

import (
	"log/slog"
	"testing"
	"time"
)

func TestLoadDefaults(t *testing.T) {
	t.Setenv("APP_NAME", "")
	t.Setenv("APP_ENV", "")
	t.Setenv("HTTP_ADDR", "")
	t.Setenv("DATABASE_URL", "")
	t.Setenv("LOG_LEVEL", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.AppName != "go-api" {
		t.Fatalf("AppName = %q, want go-api", cfg.AppName)
	}
	if cfg.LogLevel != slog.LevelInfo {
		t.Fatalf("LogLevel = %v, want info", cfg.LogLevel)
	}
}

func TestLoadOverrides(t *testing.T) {
	t.Setenv("APP_NAME", "test-api")
	t.Setenv("APP_ENV", "test")
	t.Setenv("HTTP_ADDR", ":9090")
	t.Setenv("DATABASE_URL", "postgres://user:pass@localhost:5432/test?sslmode=disable")
	t.Setenv("LOG_LEVEL", "debug")
	t.Setenv("READ_TIMEOUT", "2s")
	t.Setenv("BASIC_AUTH_USERNAME", "user")
	t.Setenv("BASIC_AUTH_PASSWORD", "pass")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.AppName != "test-api" || cfg.Environment != "test" || cfg.HTTPAddr != ":9090" {
		t.Fatalf("unexpected app config: %+v", cfg)
	}
	if cfg.LogLevel != slog.LevelDebug {
		t.Fatalf("LogLevel = %v, want debug", cfg.LogLevel)
	}
	if cfg.HTTPServer.ReadTimeout != 2*time.Second {
		t.Fatalf("ReadTimeout = %v, want 2s", cfg.HTTPServer.ReadTimeout)
	}
	if cfg.BasicAuth.Username != "user" || cfg.BasicAuth.Password != "pass" {
		t.Fatalf("unexpected basic auth config: %+v", cfg.BasicAuth)
	}
}

func TestLoadInvalidLogLevel(t *testing.T) {
	t.Setenv("LOG_LEVEL", "trace")

	if _, err := Load(); err == nil {
		t.Fatal("Load() error = nil, want error")
	}
}
