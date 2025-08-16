# Build stage for frontend
FROM node:18-alpine AS frontend-builder

# Set working directory for frontend
WORKDIR /frontend

# Copy frontend package files
COPY frontend/package*.json ./

# Install frontend dependencies
RUN npm ci --only=production

# Copy frontend source code
COPY frontend/ ./

# Build the frontend
RUN npm run build

# Build stage for Go backend
FROM golang:1.21-alpine AS backend-builder

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy the binary from backend builder stage
COPY --from=backend-builder /app/main .

# Create dist directory and copy frontend build files
RUN mkdir -p dist
COPY --from=frontend-builder /frontend/build/ ./dist/

# Expose port (will be overridden by PORT env var)
EXPOSE 8080

# Run the application
CMD ["./main"]
