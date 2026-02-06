// Package errors defines reusable error helpers for the API.
package errors

import "fmt"

// AppError represents an application-specific error with HTTP context.
type AppError struct {
	StatusCode int
	Message    string
	Err        error
}

// Error satisfies the error interface.
func (ae *AppError) Error() string {
	if ae.Err != nil {
		return fmt.Sprintf("%s: %v", ae.Message, ae.Err)
	}
	return ae.Message
}

// NewAppError constructs a new AppError instance.
func NewAppError(statusCode int, message string, err error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

// Common error identifiers.
const (
	ErrInvalidInput       = "invalid input"
	ErrUserNotFound       = "user not found"
	ErrEventNotFound      = "event not found"
	ErrUnauthorized       = "unauthorized"
	ErrInvalidCredentials = "invalid credentials"
	ErrDatabaseError      = "database error"
	ErrDuplicateUser      = "user already exists"
	ErrInvalidToken       = "invalid token"
)
