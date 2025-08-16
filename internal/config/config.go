package config

import (
	"os"
)

// Config holds the port configuration for the application
type Config struct {
	Port string
}

// Load loads the port configuration from environment variables
func Load() *Config {
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
