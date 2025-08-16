package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"vibed-traveller/internal/middleware"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
}

// SetupRoutes configures all the routes for the application
func SetupRoutes() *gin.Engine {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router (without default middleware)
	r := gin.New()

	// Add custom logging middleware
	r.Use(middleware.RequestLoggingMiddleware())

	// Add recovery middleware
	r.Use(gin.Recovery())

	// Health check endpoint
	r.GET("/health", healthHandler)

	// Root endpoint
	r.GET("/", rootHandler)

	return r
}

// healthHandler handles the health check endpoint
func healthHandler(c *gin.Context) {
	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "vibed-traveller-backend",
	}

	c.JSON(http.StatusOK, response)
}

// rootHandler handles the root endpoint
func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Vibed Traveller Backend",
		"version": "1.0.0",
	})
}
