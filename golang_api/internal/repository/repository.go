// Package repository provides data access abstractions for the application.
package repository

import (
	"udemy-multi-api-golang/models"

	"gorm.io/gorm"
)

// UserRepository defines operations for managing users.
type UserRepository interface {
	Create(email, passwordHash string) (int64, error)
	GetByEmail(email string) (int64, string, error)
	Delete(id int64) error
	Exists(email string) (bool, error)
}

// EventRepository defines operations for managing events.
type EventRepository interface {
	Create(event *models.Event) error
	GetAll() ([]models.Event, error)
	GetByID(id int64) (*models.Event, error)
	Update(event *models.Event) error
	Delete(id int64) error
}

// RegistrationRepository defines operations for event registrations.
type RegistrationRepository interface {
	Register(eventID, userID int64) error
	Unregister(eventID, userID int64) error
	IsRegistered(eventID, userID int64) (bool, error)
}

// BaseRepository provides access to the shared *gorm.DB instance.
type BaseRepository struct {
	db *gorm.DB
}

// NewBaseRepository constructs a BaseRepository.
func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

// DB exposes the underlying gorm.DB reference.
func (br *BaseRepository) DB() *gorm.DB {
	return br.db
}
