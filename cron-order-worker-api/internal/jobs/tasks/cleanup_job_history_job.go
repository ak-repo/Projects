package tasks

import (
	"context"
	"log/slog"

	"cron-order-worker-api/internal/services"
)

type CleanupJobHistoryJob struct {
	jobService *services.JobService
	logger     *slog.Logger
}

func NewCleanupJobHistoryJob(jobService *services.JobService, logger *slog.Logger) *CleanupJobHistoryJob {
	return &CleanupJobHistoryJob{jobService: jobService, logger: logger}
}

func (j *CleanupJobHistoryJob) Name() string {
	return "cleanup_job_history"
}

func (j *CleanupJobHistoryJob) Description() string {
	return "Deletes old job history records older than 30 days"
}

func (j *CleanupJobHistoryJob) Run(ctx context.Context) error {
	deleted, err := j.jobService.CleanupOldHistory(ctx, 30)
	if err != nil {
		return err
	}
	j.logger.Info("cleanup job history completed", "deleted", deleted)
	return nil
}
