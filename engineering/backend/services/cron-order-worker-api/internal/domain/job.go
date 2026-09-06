package domain

import "time"

type JobHistory struct {
	ID           int64      `json:"id"`
	JobName      string     `json:"job_name"`
	Status       string     `json:"status"`
	StartedAt    time.Time  `json:"started_at"`
	FinishedAt   *time.Time `json:"finished_at"`
	DurationMS   *int64     `json:"duration_ms"`
	ErrorMessage *string    `json:"error_message"`
	TriggeredBy  string     `json:"triggered_by"`
	CreatedAt    time.Time  `json:"created_at"`
}
