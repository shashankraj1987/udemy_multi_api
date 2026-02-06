// Package config provides configuration management for the event registration API.
package config

import (
	"log"
	"os"
	"strconv"
)

// Config aggregates all configuration sections.
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
	Logging  LoggingConfig
}

// DatabaseConfig holds database-specific configuration options.
type DatabaseConfig struct {
	Path            string
	MaxOpenConns    int
	MaxIdleConns    int
	MaxConnLifetime int // seconds
}

// ServerConfig tracks HTTP server settings.
type ServerConfig struct {
	Port string
	Host string
}

// JWTConfig holds token-related configuration.
type JWTConfig struct {
	SecretKey      string
	TokenExpiryHrs int
}

// LoggingConfig controls logging behavior.
type LoggingConfig struct {
	Level string
}

// Load reads configuration from environment variables (with defaults).
func Load() *Config {
	return &Config{
		Database: DatabaseConfig{
			Path:            getEnv("DB_PATH", "./api.db"),
			MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 10),
			MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 5),
			MaxConnLifetime: getEnvInt("DB_MAX_CONN_LIFETIME", 3600),
		},
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8081"),
			Host: getEnv("SERVER_HOST", "127.0.0.1"),
		},
		JWT: JWTConfig{
			SecretKey:      getEnv("JWT_SECRET_KEY", "superSecretKey"),
			TokenExpiryHrs: getEnvInt("JWT_TOKEN_EXPIRY_HRS", 2),
		},
		Logging: LoggingConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid integer value for %s: %v, using default: %d\n", key, err, defaultValue)
		return defaultValue
	}
	return intVal
}
