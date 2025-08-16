# Vibed Traveller

A fullstack travel application with a Go backend and modern frontend.

## Backend

The backend is built with Go using the Gin framework and provides a simple HTTP API. The project follows Go best practices with a clean, modular structure and simple configuration management.

### Prerequisites

- Go 1.21 or later

### Running the Backend

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the server:
   ```bash
   go run cmd/main.go
   ```

3. The server will start on port 8080

### API Endpoints

- `GET /` - Welcome message
- `GET /health` - Health check endpoint

### Testing the Health Endpoint

```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T12:00:00Z",
  "service": "vibed-traveller-backend"
}
```

## Docker

The application is containerized for easy deployment and development.

### Building and Running with Docker

1. **Build the Docker image:**
   ```bash
   make docker-build
   # or
   docker build -t vibed-traveller-backend .
   ```

2. **Run with Docker:**
   ```bash
   make docker-run
   # or
   docker run -p 8080:8080 -e PORT=8080 vibed-traveller-backend
   ```

3. **Using Docker Compose (recommended):**
   ```bash
   # Start the service
   make docker-compose-up
   # or
   docker-compose up -d

   # View logs
   make docker-compose-logs
   # or
   docker-compose logs -f

   # Stop the service
   make docker-compose-down
   # or
   docker-compose down
   ```

### Configuration

The application uses annotation-based configuration with struct tags to automatically map environment variables to configuration fields.

#### Configuration Structure

```go
type Config struct {
    Port     string `env:"PORT" default:"8080"`
    LogLevel string `env:"LOG_LEVEL" default:"info"`
}
```

#### Environment Variables

- `PORT` - Server port (defaults to 8080)
- `LOG_LEVEL` - Log level (defaults to info)

#### Usage

```bash
# Use default port (8080)
go run cmd/main.go

# Use custom port via environment variable
PORT=3000 go run cmd/main.go

# Use custom log level
LOG_LEVEL=debug go run cmd/main.go

# Use .env file
cp .env.example .env
# Edit .env file to set PORT=3000 and LOG_LEVEL=debug
go run cmd/main.go
```

#### Priority Order

1. **Environment variables** (highest priority)
2. **.env file** (if exists)
3. **Default values** (lowest priority)

#### Benefits of Annotation-Based Configuration

- **Self-Documenting**: Struct tags show exactly which env vars map to which fields
- **Type Safe**: Automatic type conversion for strings, ints, and bools
- **Extensible**: Easy to add new configuration fields with appropriate tags
- **Maintainable**: Clear mapping between code and environment variables

#### Logging

The application uses Go's built-in `slog` package for structured logging with configurable levels:

- **Available Levels**: `debug`, `info`, `warn`, `error`
- **Structured Output**: JSON-like key-value pairs for easy parsing
- **Configurable**: Set log level via `LOG_LEVEL` environment variable
- **Request Logging**: Automatic logging of HTTP requests with IP and user agent

### Docker Commands

- `make docker-build` - Build Docker image
- `make docker-run` - Run container directly
- `make docker-compose-up` - Start with Docker Compose
- `make docker-compose-down` - Stop Docker Compose services
- `make docker-compose-logs` - View logs
- `make docker-clean` - Clean up Docker resources

## Project Structure

```
vibed-traveller/
├── cmd/
│   └── main.go          # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go    # Configuration management
│   └── routes/
│       └── routes.go    # HTTP route definitions
├── bin/                  # Build artifacts
├── Dockerfile           # Container configuration
├── docker-compose.yml   # Docker Compose setup
├── Makefile             # Build and development commands
└── README.md            # This file
```

### Package Organization

- **`cmd/main.go`**: Application entry point that orchestrates the startup
- **`internal/config`**: Configuration management and environment variable handling
- **`internal/routes`**: HTTP route definitions and handlers using Gin framework
- **`bin/`**: Compiled binary output directory

## Frontend

Coming soon...
