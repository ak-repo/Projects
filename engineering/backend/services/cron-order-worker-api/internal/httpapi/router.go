package httpapi

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"cron-order-worker-api/internal/httpapi/handlers"
	"cron-order-worker-api/internal/jobs"
	"cron-order-worker-api/internal/services"
)

type RouterDeps struct {
	Registry     *jobs.Registry
	Runner       *jobs.Runner
	OrderService *services.OrderService
	JobService   *services.JobService
}

func NewRouter(deps RouterDeps) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(corsMiddleware)

	healthHandler := handlers.NewHealthHandler()
	orderHandler := handlers.NewOrderHandler(deps.OrderService)
	jobHandler := handlers.NewJobHandler(deps.Registry, deps.Runner, deps.JobService)

	r.Get("/", serveIndex)
	r.Get("/health", healthHandler.Check)

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/orders", orderHandler.ListOrders)
		r.Post("/orders/seed", orderHandler.SeedOrders)
		r.Get("/jobs", jobHandler.ListJobs)
		r.Post("/jobs/{name}/run", jobHandler.RunJob)
		r.Get("/jobs/history", jobHandler.ListHistory)
	})

	return r
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	content, err := os.ReadFile("web/index.html")
	if err != nil {
		http.Error(w, "index.html not found. Run the app from the project root.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(content)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
