package repository

import (
	"udemy-multi-api-golang/models"

	"gorm.io/gorm"
)

// RegistrationRepositoryImpl implements RegistrationRepository using GORM.
type RegistrationRepositoryImpl struct {
	*BaseRepository
}

// NewRegistrationRepository creates a new RegistrationRepository.
func NewRegistrationRepository(db *gorm.DB) RegistrationRepository {
	return &RegistrationRepositoryImpl{BaseRepository: NewBaseRepository(db)}
}

// Register stores a new registration record.
func (rr *RegistrationRepositoryImpl) Register(eventID, userID int64) error {
	registration := &models.Registration{EventID: eventID, UserID: userID}
	return rr.DB().Create(registration).Error
}

// Unregister removes an existing registration record.
func (rr *RegistrationRepositoryImpl) Unregister(eventID, userID int64) error {
	return rr.DB().Where("event_id = ? AND user_id = ?", eventID, userID).
		Delete(&models.Registration{}).Error
}

// IsRegistered reports whether a registration exists for the given event/user pair.
func (rr *RegistrationRepositoryImpl) IsRegistered(eventID, userID int64) (bool, error) {
	var count int64
	if err := rr.DB().Model(&models.Registration{}).
		Where("event_id = ? AND user_id = ?", eventID, userID).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
