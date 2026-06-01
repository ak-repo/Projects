package app

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"time"

	"cron-order-worker-api/internal/config"
	"cron-order-worker-api/internal/database"
	"cron-order-worker-api/internal/httpapi"
	"cron-order-worker-api/internal/jobs"
	"cron-order-worker-api/internal/jobs/tasks"
	"cron-order-worker-api/internal/repositories"
	"cron-order-worker-api/internal/services"
)

type App struct {
	cfg       config.Config
	db        *sql.DB
	server    *http.Server
	scheduler *jobs.Scheduler
	logger    *slog.Logger
}

func New(ctx context.Context, cfg config.Config, logger *slog.Logger) (*App, error) {
	db, err := database.NewMySQLDB(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}

	orderRepo := repositories.NewOrderRepository(db)
	jobHistoryRepo := repositories.NewJobHistoryRepository(db)

	orderService := services.NewOrderService(orderRepo, logger)
	jobService := services.NewJobService(jobHistoryRepo)

	registry := jobs.NewRegistry()
	runner := jobs.NewRunner(jobHistoryRepo, logger)

	registry.Register(tasks.NewRetryFailedOrdersJob(orderService, logger))
	registry.Register(tasks.NewCleanupJobHistoryJob(jobService, logger))

	scheduler := jobs.NewScheduler(registry, runner, logger)
	if err := scheduler.AddJob(ctx, "retry_failed_orders", cfg.Job.RetryFailedOrdersSchedule); err != nil {
		_ = db.Close()
		return nil, err
	}
	if err := scheduler.AddJob(ctx, "cleanup_job_history", cfg.Job.CleanupHistorySchedule); err != nil {
		_ = db.Close()
		return nil, err
	}

	router := httpapi.NewRouter(httpapi.RouterDeps{
		Registry:     registry,
		Runner:       runner,
		OrderService: orderService,
		JobService:   jobService,
	})

	server := &http.Server{
		Addr:         ":" + cfg.App.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &App{cfg: cfg, db: db, server: server, scheduler: scheduler, logger: logger}, nil
}

func (a *App) Start() error {
	a.scheduler.Start()
	a.logger.Info("http server started", "addr", a.server.Addr)
	return a.server.ListenAndServe()
}

func (a *App) Shutdown(ctx context.Context) error {
	a.logger.Info("application shutdown started")
	a.scheduler.Stop(ctx)

	if err := a.server.Shutdown(ctx); err != nil {
		return err
	}
	if err := a.db.Close(); err != nil {
		return err
	}

	a.logger.Info("application shutdown completed")
	return nil
}
