package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds the port configuration for the application
type Config struct {
	Port string
}

// Load loads the port configuration from environment variables and .env file
func Load() *Config {
	// Load .env file if it exists (ignores error if file doesn't exist)
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port: port,
	}
}

// GetPort returns the configured port
func (c *Config) GetPort() string {
	return c.Port
}
