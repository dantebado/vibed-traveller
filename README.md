# Vibed Traveller

A fullstack travel application with a Go backend and modern frontend.

## Backend

The backend is built with Go and provides a simple HTTP API.

### Prerequisites

- Go 1.21 or later

### Running the Backend

1. Install dependencies:
   ```bash
   go mod tidy
   ```

2. Run the server:
   ```bash
   go run main.go
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

### Environment Variables

- `PORT` - Server port (defaults to 8080)

### Docker Commands

- `make docker-build` - Build Docker image
- `make docker-run` - Run container directly
- `make docker-compose-up` - Start with Docker Compose
- `make docker-compose-down` - Stop Docker Compose services
- `make docker-compose-logs` - View logs
- `make docker-clean` - Clean up Docker resources

## Frontend

Coming soon...
