package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	AppName         string
	Environment     string
	HTTPAddr        string
	LogLevel        slog.Level
	ShutdownTimeout time.Duration
	HTTPServer      HTTPServerConfig
	Database        DatabaseConfig
	BasicAuth       BasicAuthConfig
}

type HTTPServerConfig struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

type DatabaseConfig struct {
	URL             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type BasicAuthConfig struct {
	Username string
	Password string
}

func Load() (Config, error) {
	cfg := Config{
		AppName:         getEnv("APP_NAME", "go-api"),
		Environment:     getEnv("APP_ENV", "development"),
		HTTPAddr:        getEnv("HTTP_ADDR", ":8080"),
		ShutdownTimeout: getDuration("SHUTDOWN_TIMEOUT", 10*time.Second),
		HTTPServer: HTTPServerConfig{
			ReadTimeout:  getDuration("READ_TIMEOUT", 5*time.Second),
			WriteTimeout: getDuration("WRITE_TIMEOUT", 10*time.Second),
			IdleTimeout:  getDuration("IDLE_TIMEOUT", 60*time.Second),
		},
		Database: DatabaseConfig{
			URL:             getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/app?sslmode=disable"),
			MaxOpenConns:    getInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:    getInt("DB_MAX_IDLE_CONNS", 25),
			ConnMaxLifetime: getDuration("DB_CONN_MAX_LIFETIME", 5*time.Minute),
		},
		BasicAuth: BasicAuthConfig{
			Username: getEnv("BASIC_AUTH_USERNAME", "admin"),
			Password: getEnv("BASIC_AUTH_PASSWORD", "change-me"),
		},
	}

	level, err := parseLogLevel(getEnv("LOG_LEVEL", "info"))
	if err != nil {
		return Config{}, err
	}
	cfg.LogLevel = level

	if cfg.Database.URL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL is required")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok || strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}

func getInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok || strings.TrimSpace(value) == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func getDuration(key string, fallback time.Duration) time.Duration {
	value, ok := os.LookupEnv(key)
	if !ok || strings.TrimSpace(value) == "" {
		return fallback
	}

	parsed, err := time.ParseDuration(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func parseLogLevel(value string) (slog.Level, error) {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "debug":
		return slog.LevelDebug, nil
	case "info", "":
		return slog.LevelInfo, nil
	case "warn", "warning":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelInfo, fmt.Errorf("invalid LOG_LEVEL %q", value)
	}
}
