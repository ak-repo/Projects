package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	App AppConfig
	DB  DBConfig
	Job JobConfig
}

type AppConfig struct {
	Env  string
	Port string
}

type DBConfig struct {
	Host                   string
	Port                   string
	User                   string
	Password               string
	Name                   string
	MaxOpenConns           int
	MaxIdleConns           int
	ConnMaxLifetimeMinutes int
}

type JobConfig struct {
	RetryFailedOrdersSchedule string
	CleanupHistorySchedule    string
}

func Load() Config {
	return Config{
		App: AppConfig{
			Env:  getEnv("APP_ENV", "development"),
			Port: getEnv("APP_PORT", "8080"),
		},
		DB: DBConfig{
			Host:                   getEnv("DB_HOST", "localhost"),
			Port:                   getEnv("DB_PORT", "3306"),
			User:                   getEnv("DB_USER", "cron_user"),
			Password:               getEnv("DB_PASSWORD", "cron_password"),
			Name:                   getEnv("DB_NAME", "cron_demo"),
			MaxOpenConns:           getEnvAsInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns:           getEnvAsInt("DB_MAX_IDLE_CONNS", 5),
			ConnMaxLifetimeMinutes: getEnvAsInt("DB_CONN_MAX_LIFETIME_MINUTES", 30),
		},
		Job: JobConfig{
			RetryFailedOrdersSchedule: getEnv("JOB_RETRY_FAILED_ORDERS_SCHEDULE", "0 */1 * * * *"),
			CleanupHistorySchedule:    getEnv("JOB_CLEANUP_HISTORY_SCHEDULE", "0 */5 * * * *"),
		},
	}
}

func (c DBConfig) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
