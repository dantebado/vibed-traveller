package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
}

func main() {
	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Set Gin to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		response := HealthResponse{
			Status:    "healthy",
			Timestamp: time.Now(),
			Service:   "vibed-traveller-backend",
		}
		c.JSON(http.StatusOK, response)
	})

	// Root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Vibed Traveller Backend",
			"version": "1.0.0",
		})
	})

	log.Printf("Starting server on :%s", port)
	log.Fatal(r.Run(":" + port))
}
