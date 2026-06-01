package jobs

import (
	"context"
	"log/slog"
	"time"

	"github.com/robfig/cron/v3"
)

type Scheduler struct {
	cron     *cron.Cron
	registry *Registry
	runner   *Runner
	logger   *slog.Logger
}

func NewScheduler(registry *Registry, runner *Runner, logger *slog.Logger) *Scheduler {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(cron.DefaultLogger),
			cron.SkipIfStillRunning(cron.DefaultLogger),
		),
	)

	return &Scheduler{cron: c, registry: registry, runner: runner, logger: logger}
}

func (s *Scheduler) AddJob(ctx context.Context, jobName string, schedule string) error {
	job, err := s.registry.Get(jobName)
	if err != nil {
		return err
	}

	_, err = s.cron.AddFunc(schedule, func() {
		jobCtx, cancel := context.WithTimeout(ctx, 2*time.Minute)
		defer cancel()

		if err := s.runner.Run(jobCtx, job, "scheduler"); err != nil {
			s.logger.Error("scheduled job execution failed", "job", job.Name(), "error", err)
		}
	})
	return err
}

func (s *Scheduler) Start() {
	s.logger.Info("cron scheduler started")
	s.cron.Start()
}

func (s *Scheduler) Stop(ctx context.Context) {
	s.logger.Info("cron scheduler stopping")
	stopCtx := s.cron.Stop()

	select {
	case <-stopCtx.Done():
		s.logger.Info("cron scheduler stopped")
	case <-ctx.Done():
		s.logger.Warn("cron scheduler stop timeout")
	}
}
