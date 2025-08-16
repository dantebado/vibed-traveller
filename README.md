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

## Frontend

Coming soon...
