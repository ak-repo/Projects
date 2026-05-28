# Go Backend Skeleton

Production-ready starter skeleton for Go HTTP APIs using the standard library, chi, slog, database/sql, and PostgreSQL via pgx.

## Features

- Standard `net/http` server with graceful shutdown
- `github.com/go-chi/chi/v5` router
- Structured logging with `log/slog`
- PostgreSQL through `database/sql` and `pgx`
- Environment-based configuration
- Health and readiness endpoints
- Basic Auth middleware
- JSON response helpers
- Docker and Docker Compose for local development
- Makefile developer commands
- Basic tests

## Requirements

- Go 1.23+
- Docker and Docker Compose, optional for local PostgreSQL
- PostgreSQL, if running without Docker Compose

## Setup

```sh
cp .env.example .env
make tidy
make test
```

Start PostgreSQL and the API with Docker Compose:

```sh
make docker-up
```

Or run the API directly against a local database:

```sh
make run
```

The API listens on `:8080` by default.

## Configuration

Configuration is loaded from environment variables.

| Variable | Default | Description |
| --- | --- | --- |
| `APP_NAME` | `go-api` | Application name |
| `APP_ENV` | `development` | Runtime environment |
| `HTTP_ADDR` | `:8080` | HTTP listen address |
| `LOG_LEVEL` | `info` | `debug`, `info`, `warn`, or `error` |
| `DATABASE_URL` | `postgres://postgres:postgres@localhost:5432/app?sslmode=disable` | PostgreSQL connection string |
| `DB_MAX_OPEN_CONNS` | `25` | Max open database connections |
| `DB_MAX_IDLE_CONNS` | `25` | Max idle database connections |
| `DB_CONN_MAX_LIFETIME` | `5m` | Max database connection lifetime |
| `BASIC_AUTH_USERNAME` | `admin` | Basic Auth username |
| `BASIC_AUTH_PASSWORD` | `change-me` | Basic Auth password |
| `READ_TIMEOUT` | `5s` | HTTP read timeout |
| `WRITE_TIMEOUT` | `10s` | HTTP write timeout |
| `IDLE_TIMEOUT` | `60s` | HTTP idle timeout |
| `SHUTDOWN_TIMEOUT` | `10s` | Graceful shutdown timeout |

## Endpoints

| Method | Path | Auth | Description |
| --- | --- | --- | --- |
| `GET` | `/healthz` | No | Liveness check |
| `GET` | `/readyz` | No | Database readiness check |
| `GET` | `/ping` | No | Ping endpoint |
| `GET` | `/version` | No | Build and app metadata |
| `GET` | `/api/v1/me` | Basic Auth | Protected example endpoint |

Example protected request:

```sh
curl -u admin:change-me http://localhost:8080/api/v1/me
```

## Make Commands

```sh
make run          # run the API locally
make test         # run tests
make test-race    # run tests with race detector
make build        # build binary into bin/api
make fmt          # format code
make vet          # run go vet
make tidy         # tidy modules
make docker-build # build Docker image
make docker-up    # start local stack
make docker-down  # stop local stack
```

## Migrations

`migrations/000001_init.sql` is a placeholder for project-specific schema. Docker Compose mounts the directory into PostgreSQL initialization for local development.

For production projects, add a migration tool such as `golang-migrate`, `tern`, or your platform's migration runner.

## Build Metadata

Version metadata is injected with linker flags from the Makefile and Dockerfile into `internal/version`.

## Project Layout

```txt
cmd/api             application entrypoint
internal/app        dependency wiring
internal/config     environment config
internal/httpserver HTTP server wrapper
internal/router     route definitions
internal/handlers   HTTP handlers
internal/middleware HTTP middleware
internal/platform   infrastructure integrations
internal/response   JSON helpers
internal/version    build metadata
migrations          SQL migrations
scripts             developer scripts
```
