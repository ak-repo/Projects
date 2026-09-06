package repositories

import (
	"context"
	"database/sql"
	"time"

	"cron-order-worker-api/internal/domain"
)

type JobHistoryRepository struct {
	db *sql.DB
}

func NewJobHistoryRepository(db *sql.DB) *JobHistoryRepository {
	return &JobHistoryRepository{db: db}
}

func (r *JobHistoryRepository) CreateRunning(ctx context.Context, jobName string, triggeredBy string) (int64, error) {
	result, err := r.db.ExecContext(ctx, `
		INSERT INTO job_history (job_name, status, started_at, triggered_by)
		VALUES (?, 'running', NOW(), ?)
	`, jobName, triggeredBy)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *JobHistoryRepository) FinishSuccess(ctx context.Context, id int64, durationMS int64) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE job_history
		SET status = 'success', finished_at = NOW(), duration_ms = ?
		WHERE id = ?
	`, durationMS, id)
	return err
}

func (r *JobHistoryRepository) FinishFailed(ctx context.Context, id int64, durationMS int64, errorMessage string) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE job_history
		SET status = 'failed', finished_at = NOW(), duration_ms = ?, error_message = ?
		WHERE id = ?
	`, durationMS, errorMessage, id)
	return err
}

func (r *JobHistoryRepository) CreateSkipped(ctx context.Context, jobName string, triggeredBy string, reason string) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO job_history (job_name, status, started_at, finished_at, duration_ms, error_message, triggered_by)
		VALUES (?, 'skipped', NOW(), NOW(), 0, ?, ?)
	`, jobName, reason, triggeredBy)
	return err
}

func (r *JobHistoryRepository) List(ctx context.Context, limit int) ([]domain.JobHistory, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, job_name, status, started_at, finished_at, duration_ms, error_message, triggered_by, created_at
		FROM job_history
		ORDER BY id DESC
		LIMIT ?
	`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	histories := make([]domain.JobHistory, 0)
	for rows.Next() {
		var history domain.JobHistory
		if err := rows.Scan(
			&history.ID,
			&history.JobName,
			&history.Status,
			&history.StartedAt,
			&history.FinishedAt,
			&history.DurationMS,
			&history.ErrorMessage,
			&history.TriggeredBy,
			&history.CreatedAt,
		); err != nil {
			return nil, err
		}
		histories = append(histories, history)
	}
	return histories, rows.Err()
}

func (r *JobHistoryRepository) DeleteOlderThan(ctx context.Context, olderThan time.Time) (int64, error) {
	result, err := r.db.ExecContext(ctx, `
		DELETE FROM job_history
		WHERE created_at < ?
	`, olderThan)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
