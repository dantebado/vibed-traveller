package main

import (
	"log/slog"
	"os"

	"vibed-traveller/internal/config"
	"vibed-traveller/internal/routes"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Setup structured logging with configured level
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: cfg.GetSlogLevel(),
	}))
	slog.SetDefault(logger)

	// Setup routes
	r := routes.SetupRoutes()

	// Start server
	slog.Info("Starting server", "port", cfg.GetPort())
	if err := r.Run(":" + cfg.GetPort()); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
