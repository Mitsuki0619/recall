.PHONY: help dev build clean test docker-up docker-down docker-restart docker-logs

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

# Development (Docker-based)
dev: docker-up ## Start development environment with hot reload in Docker
	@echo "Development environment started! Access at http://localhost:8080"
	@echo "Use 'make logs' to see application logs"

dev-local: ## Start development server locally (without Docker)
	go run ./cmd/dev

build: ## Build the application in Docker
	docker-compose exec app go build -o bin/server ./cmd/dev

build-lambda: ## Build for AWS Lambda in Docker
	docker-compose exec app sh -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags \"-static\"' -o bin/bootstrap ./cmd/lambda"

# Testing (Docker-based)
test: ## Run tests in Docker
	docker-compose exec app go test -v ./...

test-coverage: ## Run tests with coverage in Docker
	docker-compose exec app go test -v -coverprofile=coverage.out ./...
	docker-compose exec app go tool cover -html=coverage.out -o coverage.html

test-local: ## Run tests locally
	go test -v ./...

# Dependencies
deps: ## Install/update dependencies in Docker
	docker-compose exec app go mod download
	docker-compose exec app go mod tidy

deps-local: ## Install/update dependencies locally
	go mod download
	go mod tidy

# GraphQL
generate: ## Generate GraphQL code in Docker
	docker-compose exec app go run github.com/99designs/gqlgen generate

generate-local: ## Generate GraphQL code locally
	go run github.com/99designs/gqlgen generate

# Shell access
shell: ## Access app container shell
	docker-compose exec app sh

# Code quality
lint: ## Run linter in Docker (using go vet as fallback)
	docker-compose exec app go vet ./...

format: ## Format code in Docker
	docker-compose exec app go fmt ./...

# Docker operations
docker-up: ## Start all containers
	docker-compose up -d

docker-down: ## Stop all containers
	docker-compose down

docker-restart: ## Restart all containers
	docker-compose restart

docker-logs: ## Show container logs
	docker-compose logs -f

docker-logs-app: ## Show app container logs
	docker-compose logs -f app

docker-build: ## Build Docker images
	docker-compose build

docker-clean: ## Clean Docker containers and volumes
	docker-compose down -v
	docker system prune -f

# Database operations
db-migrate: ## Run database migrations
	docker exec server-postgres-1 psql -U recall -d recall_dev -c "SELECT version();"

db-reset: ## Reset database (remove all data)
	docker-compose down postgres postgres-test
	docker volume rm server_postgres_data server_postgres_test_data || true
	docker-compose up -d postgres postgres-test

# Utility
clean: ## Clean build artifacts
	rm -rf bin/
	rm -rf tmp/
	rm -f coverage.out coverage.html

# Development workflow shortcuts
start: dev ## Start development (alias for dev)
stop: docker-down ## Stop development environment  
restart: docker-restart ## Restart development environment
status: ## Show container status
	docker-compose ps

# Quick shortcuts
up: docker-up ## Alias for docker-up
down: docker-down ## Alias for docker-down
logs: docker-logs-app ## Alias for docker-logs-app