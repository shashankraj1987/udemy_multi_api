package errors
// Package errors provides custom error types for the application.
// It ensures consistent error handling and messaging across the codebase.
package errors

import "fmt"

// AppError represents an application-specific error with HTTP status code context.
type AppError struct {
	StatusCode int
	Message    string
	Err        error
}

// Error implements the error interface for AppError.
func (ae *AppError) Error() string {
	if ae.Err != nil {
		return fmt.Sprintf("%s: %v", ae.Message, ae.Err)
	}
	return ae.Message
}

// NewAppError creates a new application error.
func NewAppError(statusCode int, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,














)	ErrInvalidToken       = "invalid token"	ErrDuplicateUser      = "user already exists"	ErrDatabaseError      = "database error"	ErrInvalidCredentials = "invalid credentials"	ErrUnauthorized       = "unauthorized"	ErrEventNotFound      = "event not found"	ErrUserNotFound       = "user not found"	ErrInvalidInput       = "invalid input"const (// Common error types}	}