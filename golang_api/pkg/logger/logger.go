// Package logger provides structured logging utilities.
package logger

import (
	"log"
	"os"
)

// Logger defines the logging interface used across the app.
type Logger interface {
	Info(message string, details ...string)
	Error(message string, err error)
	Warn(message string, details ...string)
	Debug(message string, details ...string)
}

// DefaultLogger is a simple log.Logger-backed implementation.
type DefaultLogger struct {
	logger *log.Logger
	level  string
}

// New returns a DefaultLogger with the provided log level.
func New(level string) Logger {
	return &DefaultLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),
		level:  level,
	}
}

// Info logs informational messages.
func (l *DefaultLogger) Info(message string, details ...string) {
	msg := "[INFO] " + message
	if len(details) > 0 {
		msg += " - " + details[0]
	}
	l.logger.Println(msg)
}

// Error logs errors with optional error details.
func (l *DefaultLogger) Error(message string, err error) {
	if err != nil {
		l.logger.Printf("[ERROR] %s - %v\n", message, err)
		return
	}
	l.logger.Printf("[ERROR] %s\n", message)
}

// Warn logs warning messages.
func (l *DefaultLogger) Warn(message string, details ...string) {
	msg := "[WARN] " + message
	if len(details) > 0 {
		msg += " - " + details[0]
	}
	l.logger.Println(msg)
}

// Debug logs debug messages if the logger is configured for it.
func (l *DefaultLogger) Debug(message string, details ...string) {
	if l.level != "debug" {
		return
	}
	msg := "[DEBUG] " + message
	if len(details) > 0 {
		msg += " - " + details[0]
	}
	l.logger.Println(msg)
}
