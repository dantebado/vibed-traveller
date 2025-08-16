package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware creates a middleware that logs all HTTP requests
func LoggingMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// Log structured information about the request
		slog.Info("HTTP Request",
			"method", param.Method,
			"path", param.Path,
			"status", param.StatusCode,
			"latency", param.Latency,
			"client_ip", param.ClientIP,
			"user_agent", param.Request.UserAgent(),
			"timestamp", param.TimeStamp,
		)

		// Return empty string since we're handling logging ourselves
		return ""
	})
}

// RequestLoggingMiddleware creates a custom middleware for detailed request logging
func RequestLoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(start)

		// Get response status
		status := c.Writer.Status()
		statusText := getStatusText(status)

		// Log request details
		slog.Info("HTTP Request Completed",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"query", c.Request.URL.RawQuery,
			"status", status,
			"status_text", statusText,
			"latency", latency,
			"latency_ms", latency.Milliseconds(),
			"client_ip", c.ClientIP(),
			"user_agent", c.Request.UserAgent(),
			"content_length", c.Writer.Size(),
		)

		// Log errors with additional context
		if status >= 400 {
			slog.Error("HTTP Request Error",
				"method", c.Request.Method,
				"path", c.Request.URL.Path,
				"status", status,
				"status_text", statusText,
				"latency", latency,
				"client_ip", c.ClientIP(),
				"user_agent", c.Request.UserAgent(),
				"error", c.Errors.String(),
			)
		}
	}
}

// getStatusText returns the HTTP status text for a given status code
func getStatusText(status int) string {
	switch status {
	case 200:
		return "OK"
	case 201:
		return "Created"
	case 204:
		return "No Content"
	case 400:
		return "Bad Request"
	case 401:
		return "Unauthorized"
	case 403:
		return "Forbidden"
	case 404:
		return "Not Found"
	case 500:
		return "Internal Server Error"
	case 502:
		return "Bad Gateway"
	case 503:
		return "Service Unavailable"
	default:
		return "Unknown"
	}
}
