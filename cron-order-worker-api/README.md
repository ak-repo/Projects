# Cron Order Worker API

A production-shaped Golang demo project showing how to build cron jobs inside a backend API.

This project demonstrates:

- Go `net/http` backend API
- `chi` router
- MySQL database connection using `database/sql`
- Cron scheduler using `robfig/cron/v3`
- Manual job execution from API
- Job registry pattern
- Job runner with overlap protection
- Job execution history tracking
- Graceful shutdown
- Clean production-style folder structure
- Simple browser UI using `web/index.html`

---

## Project Example

The demo simulates a common backend use case:

> Some orders failed to sync with an external ERP system. A cron job runs periodically, finds retryable orders, and tries to sync them again.

The project includes two background jobs:

| Job Name | Purpose | Demo Schedule |
|---|---|---|
| `retry_failed_orders` | Retries failed or pending order sync records | Every 1 minute |
| `cleanup_job_history` | Deletes old job history records older than 30 days | Every 5 minutes |

You can also run jobs manually through the API.

---

## Folder Structure

```txt
cron-order-worker-api/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── app/
│   ├── config/
│   ├── database/
│   ├── domain/
│   ├── httpapi/
│   ├── jobs/
│   ├── repositories/
│   └── services/
├── migrations/
│   └── 001_init.sql
├── web/
│   └── index.html
├── docker-compose.yml
├── go.mod
├── Makefile
└── README.md
```

### Main folders

| Folder | Responsibility |
|---|---|
| `cmd/api` | Application entry point |
| `internal/app` | Dependency wiring / bootstrap |
| `internal/config` | Environment config loading |
| `internal/database` | MySQL connection setup |
| `internal/domain` | Domain models |
| `internal/httpapi` | Routes, handlers, responses |
| `internal/jobs` | Job interface, registry, runner, scheduler |
| `internal/jobs/tasks` | Actual cron job implementations |
| `internal/repositories` | SQL/database logic |
| `internal/services` | Business logic |
| `migrations` | Database schema |
| `web` | Browser API testing page |

---

## Requirements

- Go 1.22+
- Docker and Docker Compose
- MySQL is provided through Docker Compose

---

## Quick Start

### 1. Extract the zip

```bash
unzip cron-order-worker-api.zip
cd cron-order-worker-api
```

### 2. Start MySQL

```bash
docker compose up -d
```

The database and tables are created automatically from:

```txt
migrations/001_init.sql
```

### 3. Export environment variables

```bash
export APP_PORT=8080
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=cron_user
export DB_PASSWORD=cron_password
export DB_NAME=cron_demo
```

Optional cron schedule variables:

```bash
export JOB_RETRY_FAILED_ORDERS_SCHEDULE="0 */1 * * * *"
export JOB_CLEANUP_HISTORY_SCHEDULE="0 */5 * * * *"
```

### 4. Install dependencies

```bash
go mod tidy
```

### 5. Run the API

```bash
go run ./cmd/api
```

Or using Makefile:

```bash
make run
```

---

## Open Browser UI

After starting the API, open:

```txt
http://localhost:8080
```

This loads `web/index.html`, a simple UI for testing all APIs.

You can test:

- Health check
- List orders
- Seed demo orders
- List cron jobs
- Manually run retry job
- View job history

---

## API Endpoints

### 1. Health Check

```http
GET /health
```

Example:

```bash
curl http://localhost:8080/health
```

Response:

```json
{
  "message": "service is healthy",
  "status": "ok"
}
```

---

### 2. List Orders

```http
GET /api/v1/orders
```

Example:

```bash
curl http://localhost:8080/api/v1/orders
```

Response example:

```json
{
  "status": "ok",
  "data": [
    {
      "id": 1,
      "order_number": "ORD-1001",
      "customer_name": "Customer One",
      "amount": 2500,
      "sync_status": "failed",
      "attempts": 0,
      "last_error": "ERP timeout",
      "created_at": "2026-06-01T10:00:00Z",
      "updated_at": "2026-06-01T10:00:00Z"
    }
  ]
}
```

---

### 3. Seed Demo Orders

```http
POST /api/v1/orders/seed
```

Example:

```bash
curl -X POST http://localhost:8080/api/v1/orders/seed
```

Response:

```json
{
  "message": "demo orders seeded",
  "status": "ok"
}
```

Use this when you want to create more failed/pending orders for testing.

---

### 4. List Registered Jobs

```http
GET /api/v1/jobs
```

Example:

```bash
curl http://localhost:8080/api/v1/jobs
```

Response:

```json
{
  "status": "ok",
  "data": [
    {
      "name": "cleanup_job_history",
      "description": "Deletes old job history records older than 30 days"
    },
    {
      "name": "retry_failed_orders",
      "description": "Retries failed or pending order sync records"
    }
  ]
}
```

---

### 5. Manually Run Retry Job

```http
POST /api/v1/jobs/retry_failed_orders/run
```

Example:

```bash
curl -X POST http://localhost:8080/api/v1/jobs/retry_failed_orders/run
```

Response:

```json
{
  "job": "retry_failed_orders",
  "message": "job executed successfully",
  "status": "ok"
}
```

If the same job is already running, you will get:

```json
{
  "message": "job is already running",
  "status": "error"
}
```

with HTTP status:

```txt
409 Conflict
```

---

### 6. Manually Run Cleanup Job

```http
POST /api/v1/jobs/cleanup_job_history/run
```

Example:

```bash
curl -X POST http://localhost:8080/api/v1/jobs/cleanup_job_history/run
```

---

### 7. View Job History

```http
GET /api/v1/jobs/history
```

Example:

```bash
curl http://localhost:8080/api/v1/jobs/history
```

With limit:

```bash
curl "http://localhost:8080/api/v1/jobs/history?limit=20"
```

Response example:

```json
{
  "status": "ok",
  "data": [
    {
      "id": 1,
      "job_name": "retry_failed_orders",
      "status": "success",
      "started_at": "2026-06-01T10:00:00Z",
      "finished_at": "2026-06-01T10:00:02Z",
      "duration_ms": 2040,
      "error_message": null,
      "triggered_by": "manual",
      "created_at": "2026-06-01T10:00:00Z"
    }
  ]
}
```

---

## Demo Testing Flow

Use this sequence to clearly demonstrate the cron job system:

### Step 1: Check API health

```bash
curl http://localhost:8080/health
```

### Step 2: Check existing orders

```bash
curl http://localhost:8080/api/v1/orders
```

You should see orders with statuses like:

```txt
failed
pending
synced
```

### Step 3: Seed more failed orders

```bash
curl -X POST http://localhost:8080/api/v1/orders/seed
```

### Step 4: Run the retry job manually

```bash
curl -X POST http://localhost:8080/api/v1/jobs/retry_failed_orders/run
```

### Step 5: Check orders again

```bash
curl http://localhost:8080/api/v1/orders
```

Some orders will move from:

```txt
failed/pending -> synced
```

or:

```txt
pending/failed -> failed with attempts increased
```

### Step 6: Check job history

```bash
curl http://localhost:8080/api/v1/jobs/history
```

You should see the manual job run stored in `job_history`.

### Step 7: Wait for cron

The `retry_failed_orders` cron runs every 1 minute by default.

After one minute, check history again:

```bash
curl http://localhost:8080/api/v1/jobs/history
```

You should see records with:

```txt
triggered_by = scheduler
```

---

## Cron Schedule Format

This project uses:

```go
cron.WithSeconds()
```

So cron expressions use 6 fields:

```txt
second minute hour day month weekday
```

Examples:

| Schedule | Meaning |
|---|---|
| `*/10 * * * * *` | Every 10 seconds |
| `0 */1 * * * *` | Every 1 minute |
| `0 */10 * * * *` | Every 10 minutes |
| `0 0 2 * * *` | Every day at 2 AM |
| `0 0 1 1 * *` | First day of every month at 1 AM |

---

## Production Concepts Used

### 1. Job Registry

Jobs are registered once during app startup:

```go
registry.Register(tasks.NewRetryFailedOrdersJob(orderService, logger))
registry.Register(tasks.NewCleanupJobHistoryJob(jobService, logger))
```

The API and scheduler both use the same registry.

---

### 2. Job Runner

The job runner handles common execution behavior:

- Prevents duplicate execution of the same job
- Creates `job_history` record
- Tracks success/failure/skipped status
- Logs duration
- Stores errors

This keeps each job clean and small.

---

### 3. Overlap Protection

There are two protections:

1. `cron.SkipIfStillRunning`
2. Custom per-job mutex in `Runner`

This protects against this case:

```txt
Cron starts retry_failed_orders
User manually clicks run at the same time
Manual request returns 409 Conflict
```

---

### 4. Clean Layering

The flow is:

```txt
HTTP / Cron
    ↓
Job Runner
    ↓
Job Task
    ↓
Service
    ↓
Repository
    ↓
Database
```

Jobs do not write SQL directly. They call services.

---

## Useful Commands

```bash
make docker-up
make run
make tidy
make fmt
make test
make docker-down
```

---

## Troubleshooting

### MySQL connection refused

Make sure Docker is running:

```bash
docker compose ps
```

Check logs:

```bash
docker compose logs -f mysql
```

### Tables not created

If MySQL volume already existed before the migration file was added, recreate the volume:

```bash
docker compose down -v
docker compose up -d
```

### Browser UI cannot call API

Open the UI through the Go server:

```txt
http://localhost:8080
```

Do not open `web/index.html` directly unless you understand browser CORS behavior.

---

## Next Production Improvements

For a real production system, add:

- Authentication for manual job execution
- Role-based access for job APIs
- Migration tool such as `golang-migrate`
- Pagination metadata for job history
- Unit tests for services
- Integration tests for repositories
- External ERP client interface
- Retry backoff rules
- Metrics with Prometheus
- Distributed locking if running multiple app instances

