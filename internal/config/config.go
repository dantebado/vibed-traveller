package config

import (
	"log/slog"
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds the configuration for the application
type Config struct {
	Port     string `env:"PORT" default:"8080"`
	LogLevel string `env:"LOG_LEVEL" default:"info"`
	BaseURL  string `env:"BASE_URL" default:"http://localhost:3000"`
	APIURL   string `env:"API_URL" default:"http://localhost:8080"`

	// Auth0 Configuration
	Auth0Domain       string `env:"AUTH0_DOMAIN" default:""`
	Auth0Audience     string `env:"AUTH0_AUDIENCE" default:""`
	Auth0IssuerURL    string `env:"AUTH0_ISSUER_URL" default:""`
	Auth0ClientID     string `env:"AUTH0_CLIENT_ID" default:""`
	Auth0ClientSecret string `env:"AUTH0_CLIENT_SECRET" default:""`
}

// Load loads configuration from environment variables and .env file
func Load() *Config {
	// Load .env file if it exists (ignores error if file doesn't exist)
	err := godotenv.Load()
	if err != nil {
		slog.Warn("failed to load .env file")
	}

	config := &Config{}

	// Use reflection to automatically populate config from environment
	config.loadFromEnv()

	return config
}

// loadFromEnv uses reflection to automatically populate config fields from environment variables
func (c *Config) loadFromEnv() {
	val := reflect.ValueOf(c).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		// Get the env tag
		envKey := fieldType.Tag.Get("env")
		if envKey == "" {
			continue
		}

		// Get the default value
		defaultValue := fieldType.Tag.Get("default")

		// Get value from environment
		envValue := os.Getenv(envKey)

		// Use default if environment variable is not set
		if envValue == "" {
			envValue = defaultValue
		}

		// Set the field value based on its type
		c.setFieldValue(field, envValue)
	}
}

// setFieldValue sets the field value based on its type
func (c *Config) setFieldValue(field reflect.Value, value string) {
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Int:
		if value == "" {
			field.SetInt(0)
		} else {
			if intVal, err := strconv.Atoi(value); err == nil {
				field.SetInt(int64(intVal))
			}
		}
	case reflect.Bool:
		if value == "" {
			field.SetBool(false)
		} else {
			if boolVal, err := strconv.ParseBool(value); err == nil {
				field.SetBool(boolVal)
			}
		}
	}
}

// GetPort returns the configured port
func (c *Config) GetPort() string {
	return c.Port
}

// GetLogLevel returns the configured log level
func (c *Config) GetLogLevel() string {
	return c.LogLevel
}

// GetBaseURL returns the configured base URL
func (c *Config) GetBaseURL() string {
	return c.BaseURL
}

// GetSlogLevel returns the slog.Level for the configured log level
func (c *Config) GetSlogLevel() slog.Level {
	switch c.LogLevel {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// GetAuth0Domain returns the Auth0 domain
func (c *Config) GetAuth0Domain() string {
	return c.Auth0Domain
}

// GetAuth0Audience returns the Auth0 audience
func (c *Config) GetAuth0Audience() string {
	return c.Auth0Audience
}

// GetAuth0IssuerURL returns the Auth0 issuer URL
func (c *Config) GetAuth0IssuerURL() string {
	return c.Auth0IssuerURL
}

// GetAuth0ClientID returns the Auth0 client ID
func (c *Config) GetAuth0ClientID() string {
	return c.Auth0ClientID
}

// GetAuth0ClientSecret returns the Auth0 client secret
func (c *Config) GetAuth0ClientSecret() string {
	return c.Auth0ClientSecret
}

// IsAuth0Configured checks if Auth0 is properly configured
func (c *Config) IsAuth0Configured() bool {
	// Check if all required Auth0 fields are set
	if c.Auth0Domain == "" {
		slog.Error("AUTH0_DOMAIN is not set")
		return false
	}
	if c.Auth0Audience == "" {
		slog.Error("AUTH0_AUDIENCE is not set")
		return false
	}
	if c.Auth0IssuerURL == "" {
		slog.Error("AUTH0_ISSUER_URL is not set")
		return false
	}
	if c.Auth0ClientID == "" {
		slog.Error("AUTH0_CLIENT_ID is not set")
		return false
	}
	if c.Auth0ClientSecret == "" {
		slog.Error("AUTH0_CLIENT_SECRET is not set")
		return false
	}
	return true
}

// Debug prints the current configuration values for debugging
func (c *Config) Debug() {
	slog.Info("Current configuration",
		"port", c.Port,
		"log_level", c.LogLevel,
		"base_url", c.BaseURL,
		"auth0_domain", c.Auth0Domain,
		"auth0_audience", c.Auth0Audience,
		"auth0_issuer_url", c.Auth0IssuerURL,
		"auth0_client_id", c.Auth0ClientID,
		"auth0_client_secret_set", c.Auth0ClientSecret != "",
	)
}
