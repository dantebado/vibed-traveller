package main

import (
	"log/slog"
	"os"
	"vibed-traveller/internal/middleware"

	"vibed-traveller/internal/config"
	"vibed-traveller/internal/routes"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Logger
	level := cfg.GetSlogLevel()
	base := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level, AddSource: false})
	logger := slog.New(middleware.New(base))
	slog.SetDefault(logger)

	// Setup routes with configuration
	r := routes.SetupRoutes(cfg)

	// Start server
	slog.Info("Starting server", "port", cfg.GetPort())
	if err := r.Run(":" + cfg.GetPort()); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
