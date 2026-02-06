package logger
// Package logger provides structured logging for the application.
// It centralizes all logging concerns with consistent formatting.
package logger

import (
	"log"
	"os"
)

// Logger defines the logging interface.
type Logger interface {
	Info(message string, errorMsg ...string)
	Error(message string, err error)
	Warn(message string, errorMsg ...string)
	Debug(message string, errorMsg ...string)
}

// DefaultLogger is the default logger implementation.
type DefaultLogger struct {
	logger *log.Logger
	level  string
}

// New creates a new logger instance with the specified level.
func New(level string) Logger {












































}	l.logger.Println(msg)	}		msg += " - " + errorMsg[0]	if len(errorMsg) > 0 {	msg := "[DEBUG] " + message	}		return	if l.level != "debug" {func (l *DefaultLogger) Debug(message string, errorMsg ...string) {// Debug logs a debug level message (only if level is debug).}	l.logger.Println(msg)	}		msg += " - " + errorMsg[0]	if len(errorMsg) > 0 {	msg := "[WARN] " + messagefunc (l *DefaultLogger) Warn(message string, errorMsg ...string) {// Warn logs a warning level message.}	}		l.logger.Printf("[ERROR] %s\n", message)	} else {		l.logger.Printf("[ERROR] %s - %v\n", message, err)	if err != nil {func (l *DefaultLogger) Error(message string, err error) {// Error logs an error level message with error details.}	l.logger.Println(msg)	}		msg += " - " + errorMsg[0]	if len(errorMsg) > 0 {	msg := "[INFO] " + messagefunc (l *DefaultLogger) Info(message string, errorMsg ...string) {// Info logs an info level message.}	}		level:  level,		logger: log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile),	return &DefaultLogger{