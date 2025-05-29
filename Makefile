.PHONY: build run test clean dev deps sqlc-gen docker-build docker-run fmt lint help

BINARY_NAME=rss-aggregator
PORT=8080
DB_URL=

build:
	@echo "Building $(BINARY_NAME)..."
	go build -o bin/$(BINARY_NAME) cmd/server/main.go

run:
	@echo "Running $(BINARY_NAME)..."
	./bin/$(BINARY_NAME)

dev:
	@echo "Starting development server..."
	air -c .air.toml

clean:
	@echo "Cleaning up..."
	rm -rf bin/
	rm -f coverage.out coverage.html

deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

sqlc-gen:
	@echo "Generating SQL code..."
	sqlc generate

migrate-up:
	@echo "Running database migrations..."
	goose up

migrate-down:
	@echo "Rolling back database migrations..."
	goose down

install-tools:
	@echo "Installing development tools..."
	go install github.com/air-verse/air@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

setup: install-tools deps
	@echo "Project setup complete!"

help:
	@echo "Available commands:"
	@echo "  build          - Build the application"
	@echo "  run            - Build and run the application"
	@echo "  dev            - Run in development mode with hot reload"
	@echo "  clean          - Clean build artifacts"
	@echo "  deps           - Install dependencies"
	@echo "  sqlc-gen       - Generate SQL code using sqlc"
	@echo "  migrate-up     - Run database migrations"
	@echo "  migrate-down   - Roll back database migrations"
	@echo "  install-tools  - Install development tools"
	@echo "  setup          - Setup project (install tools and deps)"
	@echo "  help           - Show this help message"