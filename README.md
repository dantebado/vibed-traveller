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

### Authentication Endpoints

The application now includes Auth0-based authentication with automatic redirects:

- `GET /auth/login-page` - Simple HTML login page
- `GET /auth/login` - Redirects to Auth0 login
- `GET /auth/callback` - Handles Auth0 callback
- `GET /auth/logout` - Redirects to Auth0 logout

### Protected API Endpoints

These endpoints require authentication and will redirect to login if unauthorized:

- `GET /api/profile` - Get user profile (requires authentication)
- `GET /api/me` - Get current user info (requires authentication)

### Authentication Flow

1. **Unauthorized Access**: When a user tries to access a protected endpoint without authentication, they are automatically redirected to the Auth0 login page
2. **Login**: After successful login, users are redirected back to the original page they were trying to access
3. **Session Management**: JWT tokens are validated on each request to protected endpoints
4. **Logout**: Users can logout and are redirected to Auth0 logout page

### Testing Authentication

Use the provided test script to verify the authentication flow:

```bash
./test_auth.sh
```

Or test manually:

```bash
# Try to access protected endpoint (should redirect to login)
curl -L $BASE_URL/api/me

# View login page
curl $BASE_URL/auth/login-page

# Test login redirect
curl -L $BASE_URL/auth/login
```

**Note**: Replace `$BASE_URL` with your actual base URL (e.g., `http://localhost:8081`)

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
- `LOG_LEVEL` - Logging level (defaults to "info")
- `BASE_URL` - Base URL for the application (defaults to "http://localhost:8080")

#### Auth0 Configuration

For authentication to work, you must configure the following environment variables:

```bash
# Copy the example file
cp env.example .env

# Edit with your Auth0 credentials
BASE_URL=http://localhost:8080
AUTH0_DOMAIN=your-tenant.auth0.com
AUTH0_AUDIENCE=https://your-api-identifier
AUTH0_ISSUER_URL=https://your-tenant.auth0.com
AUTH0_CLIENT_ID=your-client-id
AUTH0_CLIENT_SECRET=your-client-secret
```

See `AUTH0_SETUP.md` for detailed setup instructions.

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
- **Global Middleware**: Automatic logging of all HTTP requests with comprehensive details
- **Request Details**: Method, path, status, latency, IP, user agent, content length
- **Error Logging**: Special error logging for 4xx and 5xx responses
- **Request Tracing**: Unique request ID links all logs for a single request
- **Context-Aware**: Uses `slog.InfoContext` and `slog.ErrorContext`