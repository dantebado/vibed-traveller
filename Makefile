.PHONY: run build test clean

# Run the application
run:
	go run cmd/main.go

# Build the application
build:
	go build -o bin/vibed-traveller cmd/main.go

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy

# Run with hot reload (requires air - install with: go install github.com/cosmtrek/air@latest)
dev:
	air

# Stop the server
stop:
	lsof -ti:8080 | xargs kill -9 || true

# Docker commands
docker-build:
	DOCKER_BUILDKIT=1 docker build --secret id=_env,src=frontend/.env -t vibed-traveller-backend .

docker-run:
	docker run -p 8080:8080 -e PORT=8080 --env-file .env vibed-traveller-backend

docker-compose-up:
	docker-compose up -d

docker-compose-down:
	docker-compose down

docker-compose-logs:
	docker-compose logs -f

docker-clean:
	docker system prune -f
	docker image prune -f
