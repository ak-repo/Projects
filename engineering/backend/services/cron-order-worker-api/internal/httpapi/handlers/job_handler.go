package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"cron-order-worker-api/internal/httpapi/response"
	"cron-order-worker-api/internal/jobs"
	"cron-order-worker-api/internal/services"
)

type JobHandler struct {
	registry   *jobs.Registry
	runner     *jobs.Runner
	jobService *services.JobService
}

func NewJobHandler(registry *jobs.Registry, runner *jobs.Runner, jobService *services.JobService) *JobHandler {
	return &JobHandler{registry: registry, runner: runner, jobService: jobService}
}

func (h *JobHandler) ListJobs(w http.ResponseWriter, r *http.Request) {
	allJobs := h.registry.List()
	items := make([]map[string]string, 0, len(allJobs))

	for _, job := range allJobs {
		items = append(items, map[string]string{
			"name":        job.Name(),
			"description": job.Description(),
		})
	}

	response.WriteJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
		"data":   items,
	})
}

func (h *JobHandler) RunJob(w http.ResponseWriter, r *http.Request) {
	jobName := chi.URLParam(r, "name")
	job, err := h.registry.Get(jobName)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, "job not found")
		return
	}

	jobCtx, cancel := context.WithTimeout(r.Context(), 2*time.Minute)
	defer cancel()

	err = h.runner.Run(jobCtx, job, "manual")
	if err != nil {
		if errors.Is(err, jobs.ErrJobAlreadyRunning) {
			response.WriteError(w, http.StatusConflict, "job is already running")
			return
		}
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]any{
		"status":  "ok",
		"message": "job executed successfully",
		"job":     jobName,
	})
}

func (h *JobHandler) ListHistory(w http.ResponseWriter, r *http.Request) {
	limit := 50
	if value := r.URL.Query().Get("limit"); value != "" {
		parsed, err := strconv.Atoi(value)
		if err == nil {
			limit = parsed
		}
	}

	history, err := h.jobService.ListHistory(r.Context(), limit)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
		"data":   history,
	})
}
