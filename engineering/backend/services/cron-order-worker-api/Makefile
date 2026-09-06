APP_NAME=cron-order-worker-api

run:
	go run ./cmd/api

build:
	go build -o bin/$(APP_NAME) ./cmd/api

tidy:
	go mod tidy

fmt:
	go fmt ./...

test:
	go test ./...

docker-up:
	docker compose up -d

docker-down:
	docker compose down

docker-logs:
	docker compose logs -f mysql
