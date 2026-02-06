package repository

import (
	"udemy-multi-api-golang/models"

	"gorm.io/gorm"
)

// EventRepositoryImpl implements EventRepository using GORM.
type EventRepositoryImpl struct {
	*BaseRepository
}

// NewEventRepository creates a new EventRepository instance.
func NewEventRepository(db *gorm.DB) EventRepository {
	return &EventRepositoryImpl{BaseRepository: NewBaseRepository(db)}
}

// Create persists a new event.
func (er *EventRepositoryImpl) Create(event *models.Event) error {
	return er.DB().Create(event).Error
}

// GetAll returns all events ordered by date.
func (er *EventRepositoryImpl) GetAll() ([]models.Event, error) {
	var events []models.Event
	err := er.DB().Order("dateTime asc").Find(&events).Error
	return events, err
}

// GetByID fetches an event by its primary key.
func (er *EventRepositoryImpl) GetByID(id int64) (*models.Event, error) {
	var event models.Event
	if err := er.DB().First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

// Update saves the provided event.
func (er *EventRepositoryImpl) Update(event *models.Event) error {
	return er.DB().Save(event).Error
}

// Delete removes an event by ID.
func (er *EventRepositoryImpl) Delete(id int64) error {
	return er.DB().Delete(&models.Event{}, id).Error
}
