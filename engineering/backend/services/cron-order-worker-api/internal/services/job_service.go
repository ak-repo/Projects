package services

import (
	"context"
	"time"

	"cron-order-worker-api/internal/domain"
	"cron-order-worker-api/internal/repositories"
)

type JobService struct {
	jobHistoryRepo *repositories.JobHistoryRepository
}

func NewJobService(jobHistoryRepo *repositories.JobHistoryRepository) *JobService {
	return &JobService{jobHistoryRepo: jobHistoryRepo}
}

func (s *JobService) ListHistory(ctx context.Context, limit int) ([]domain.JobHistory, error) {
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	return s.jobHistoryRepo.List(ctx, limit)
}

func (s *JobService) CleanupOldHistory(ctx context.Context, olderThanDays int) (int64, error) {
	if olderThanDays <= 0 {
		olderThanDays = 30
	}
	olderThan := time.Now().AddDate(0, 0, -olderThanDays)
	return s.jobHistoryRepo.DeleteOlderThan(ctx, olderThan)
}
