set dotenv-load := true

dev-up:
    docker compose up -d

dev-down:
    docker compose down

dev-logs:
    docker compose logs -f

migrate:
    goose  up

migrate-down:
    goose  down

migrate-status:
    goose  status

run-api:
    go run ./cmd/api

run-worker:
    go run ./cmd/worker

run-mcp:
    go run ./cmd/mcp

test:
    go test ./...

test-race:
    go test -race ./...

build:
    go build ./cmd/api
    go build ./cmd/worker
    go build ./cmd/mcp

fmt:
    gofmt -w .