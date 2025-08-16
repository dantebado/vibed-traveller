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
}

// Load loads configuration from environment variables and .env file
func Load() *Config {
	// Load .env file if it exists (ignores error if file doesn't exist)
	godotenv.Load()

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
