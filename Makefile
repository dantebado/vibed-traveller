.PHONY: run build test clean

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o bin/vibed-traveller main.go

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
