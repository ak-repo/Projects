package jobs

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"

	"cron-order-worker-api/internal/repositories"
)

var ErrJobAlreadyRunning = errors.New("job is already running")

type Runner struct {
	historyRepo *repositories.JobHistoryRepository
	logger      *slog.Logger
	mu          sync.Mutex
	locks       map[string]*sync.Mutex
}

func NewRunner(historyRepo *repositories.JobHistoryRepository, logger *slog.Logger) *Runner {
	return &Runner{historyRepo: historyRepo, logger: logger, locks: make(map[string]*sync.Mutex)}
}

func (r *Runner) Run(ctx context.Context, job Job, triggeredBy string) error {
	jobLock := r.getLock(job.Name())
	if !jobLock.TryLock() {
		_ = r.historyRepo.CreateSkipped(ctx, job.Name(), triggeredBy, ErrJobAlreadyRunning.Error())
		return ErrJobAlreadyRunning
	}
	defer jobLock.Unlock()

	start := time.Now()
	historyID, err := r.historyRepo.CreateRunning(ctx, job.Name(), triggeredBy)
	if err != nil {
		return err
	}

	r.logger.Info("job started", "job", job.Name(), "triggered_by", triggeredBy)
	err = job.Run(ctx)
	durationMS := time.Since(start).Milliseconds()

	if err != nil {
		_ = r.historyRepo.FinishFailed(ctx, historyID, durationMS, err.Error())
		r.logger.Error("job failed", "job", job.Name(), "triggered_by", triggeredBy, "duration_ms", durationMS, "error", err)
		return err
	}

	if err := r.historyRepo.FinishSuccess(ctx, historyID, durationMS); err != nil {
		return err
	}

	r.logger.Info("job finished", "job", job.Name(), "triggered_by", triggeredBy, "duration_ms", durationMS)
	return nil
}

func (r *Runner) getLock(jobName string) *sync.Mutex {
	r.mu.Lock()
	defer r.mu.Unlock()

	lock, exists := r.locks[jobName]
	if exists {
		return lock
	}

	lock = &sync.Mutex{}
	r.locks[jobName] = lock
	return lock
}
