// Package models defines data structures for the application.
package models

import "time"

// User represents a user in the system.
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// Event represents an event that users can register for.
type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required,min=1"`
	Description string    `json:"description" binding:"required,min=1"`
	Location    string    `json:"location" binding:"required,min=1"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int64     `json:"userId"`
}

// LoginRequest represents a login request.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// SignupRequest represents a signup request.
type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// CreateEventRequest represents a request to create an event.
type CreateEventRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
}

// UpdateEventRequest represents a request to update an event.
type UpdateEventRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
}
