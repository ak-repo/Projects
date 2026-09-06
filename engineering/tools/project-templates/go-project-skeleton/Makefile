APP_NAME := api
PKG := ./...
VERSION ?= dev
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || printf none)
DATE ?= $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
LDFLAGS := -s -w -X go-project-skeleton/internal/version.Version=$(VERSION) -X go-project-skeleton/internal/version.Commit=$(COMMIT) -X go-project-skeleton/internal/version.Date=$(DATE)

.PHONY: run test test-race build fmt vet tidy docker-build docker-up docker-down clean

run:
	go run ./cmd/api

test:
	go test $(PKG)

test-race:
	go test -race $(PKG)

build:
	mkdir -p bin
	go build -ldflags "$(LDFLAGS)" -o bin/$(APP_NAME) ./cmd/api

fmt:
	go fmt $(PKG)

vet:
	go vet $(PKG)

tidy:
	go mod tidy

docker-build:
	docker build --build-arg VERSION=$(VERSION) --build-arg COMMIT=$(COMMIT) --build-arg DATE=$(DATE) -t go-api:local .

docker-up:
	docker compose up --build

docker-down:
	docker compose down

clean:
	rm -rf bin coverage.out
