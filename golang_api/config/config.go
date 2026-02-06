package config
// Package config provides configuration management for the event registration API.
// It centralizes all configuration concerns and provides accessor functions.
package config

import (
	"log"
	"os"
	"strconv"
)

// Config holds all configuration values for the application.
type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	JWT      JWTConfig
	Logging  LoggingConfig
}






































































}	return intVal	}		return defaultValue		log.Printf("Invalid integer value for %s: %v, using default: %d\n", key, err, defaultValue)	if err != nil {	intVal, err := strconv.Atoi(value)	}		return defaultValue	if value == "" {	value := getEnv(key, "")func getEnvInt(key string, defaultValue int) int {// getEnvInt retrieves an environment variable as integer or returns a default value.}	return defaultValue	}		return value	if value, exists := os.LookupEnv(key); exists {func getEnv(key, defaultValue string) string {// getEnv retrieves an environment variable or returns a default value.}	}		},			Level: getEnv("LOG_LEVEL", "info"),		Logging: LoggingConfig{		},			TokenExpiryHrs: getEnvInt("JWT_TOKEN_EXPIRY_HRS", 2),			SecretKey:      getEnv("JWT_SECRET_KEY", "superSecretKey"),		JWT: JWTConfig{		},			Host: getEnv("SERVER_HOST", "127.0.0.1"),			Port: getEnv("SERVER_PORT", "8081"),		Server: ServerConfig{		},			MaxConnLifetime: getEnvInt("DB_MAX_CONN_LIFETIME", 3600),			MaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 5),			MaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 10),			Path:            getEnv("DB_PATH", "./api.db"),		Database: DatabaseConfig{	return &Config{func Load() *Config {// Load loads configuration from environment variables with sensible defaults.}	Level stringtype LoggingConfig struct {// LoggingConfig holds logging-specific configuration.}	TokenExpiryHrs int	SecretKey      stringtype JWTConfig struct {// JWTConfig holds JWT-specific configuration.}	Host string	Port stringtype ServerConfig struct {// ServerConfig holds server-specific configuration.}	MaxConnLifetime int // seconds	MaxIdleConns    int	MaxOpenConns    int	Path            stringtype DatabaseConfig struct {// DatabaseConfig holds database-specific configuration.