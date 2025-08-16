package main

import (
	"log"

	"vibed-traveller/internal/config"
	"vibed-traveller/internal/routes"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Setup routes
	r := routes.SetupRoutes()

	// Start server
	log.Printf("Starting server on :%s", cfg.GetPort())
	log.Fatal(r.Run(":" + cfg.GetPort()))
}
