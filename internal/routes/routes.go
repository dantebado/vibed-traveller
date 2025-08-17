package routes

import (
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"vibed-traveller/internal/config"
	"vibed-traveller/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
}

// SetupRoutes configures all the routes for the application
func SetupRoutes(cfg *config.Config) *gin.Engine {
	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router (without default middleware)
	r := gin.New()

	// Setup CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.BaseURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Add request ID middleware first
	r.Use(middleware.RequestIDMiddleware())

	// Add custom logging middleware
	r.Use(middleware.RequestLoggingMiddleware())

	// Add recovery middleware
	r.Use(gin.Recovery())

	// Health check endpoint
	r.GET("/health", healthHandler)

	// Setup authenticated routes if Auth0 is configured
	SetupAuthRoutes(r, cfg)

	// Serve static files from dist directory
	r.Static("/static", "./dist/static")

	// Serve the React app for all other routes (SPA routing)
	r.NoRoute(func(c *gin.Context) {
		// Try to serve the requested file first
		filePath := filepath.Join("./dist", c.Request.URL.Path)
		if _, err := os.Stat(filePath); err == nil {
			c.File(filePath)
			return
		}

		// If file doesn't exist, serve index.html for SPA routing
		c.File("./dist/index.html")
	})

	// Root endpoint - serve the React app
	r.GET("/", func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	return r
}

// healthHandler handles the health check endpoint
func healthHandler(c *gin.Context) {
	slog.InfoContext(c.Request.Context(), "Health check endpoint hit")

	response := HealthResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "vibed-traveller-backend",
	}

	c.JSON(http.StatusOK, response)
}
