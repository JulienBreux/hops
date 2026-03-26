.PHONY: build test lint docker-build docker-run

# Default target
all: build

# Build Go backend and Svelte frontend
build:
	@echo "Building Svelte frontend..."
	cd web && npm install && npm run build
	@echo "Building Go backend..."
	go build -o hops main.go

# Run tests
test:
	@echo "Running Go tests..."
	go test ./...
	@echo "Running Svelte checks..."
	cd web && npm run check

# Linting
lint:
	@echo "Linting Go code..."
	go vet ./...
	@echo "Checking Svelte code..."
	cd web && npm run check

# Docker build
docker-build:
	docker build -t hops .

# Docker run
docker-run:
	docker run -p 8080:8080 hops
