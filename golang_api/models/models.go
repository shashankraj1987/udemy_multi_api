// Package models defines data structures for the application.
package models

import "time"

// User represents a user in the system.
type User struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string    `json:"email" binding:"required,email" gorm:"uniqueIndex;size:255;not null"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Event represents an event that users can register for.
type Event struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" binding:"required,min=1" gorm:"size:255;not null"`
	Description string    `json:"description" binding:"required,min=1" gorm:"type:text;not null"`
	Location    string    `json:"location" binding:"required,min=1" gorm:"size:255;not null"`
	DateTime    time.Time `json:"dateTime" binding:"required" gorm:"column:dateTime;not null"`
	UserID      int64     `json:"userId" gorm:"column:user_id;not null"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// Registration represents a user's registration to an event.
type Registration struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	EventID   int64     `json:"eventId" gorm:"column:event_id;not null;uniqueIndex:idx_event_user"`
	UserID    int64     `json:"userId" gorm:"column:user_id;not null;uniqueIndex:idx_event_user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
