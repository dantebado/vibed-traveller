package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

const (
	// RequestIDHeader is the header name for the request ID
	RequestIDHeader = "X-Request-ID"
	// RequestIDKey is the key used to store request ID in gin context
	RequestIDKey = "request_id"
)

// RequestIDMiddleware adds a unique request ID to each request
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if request ID is already provided in headers
		requestID := c.GetHeader(RequestIDHeader)

		// If no request ID provided, generate one
		if requestID == "" {
			requestID = generateRequestID()
		}

		// Add request ID to gin context
		c.Set(RequestIDKey, requestID)

		// Add request ID to response headers
		c.Header(RequestIDHeader, requestID)

		// Continue to next middleware/handler
		c.Next()
	}
}

// GetRequestID retrieves the request ID from gin context
func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get(RequestIDKey); exists {
		if id, ok := requestID.(string); ok {
			return id
		}
	}
	return "unknown"
}

// generateRequestID generates a random 16-byte hex string
func generateRequestID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
