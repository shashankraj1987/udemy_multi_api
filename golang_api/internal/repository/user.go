package repository

import (
	"udemy-multi-api-golang/models"

	"gorm.io/gorm"
)

// UserRepositoryImpl implements UserRepository using GORM.
type UserRepositoryImpl struct {
	*BaseRepository
}

// NewUserRepository creates a new UserRepository instance.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		BaseRepository: NewBaseRepository(db),
	}
}

// Create inserts a new user record.
func (ur *UserRepositoryImpl) Create(email, passwordHash string) (int64, error) {
	user := &models.User{Email: email, Password: passwordHash}
	if err := ur.DB().Create(user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// GetByEmail returns the user's ID and hashed password for a given email.
func (ur *UserRepositoryImpl) GetByEmail(email string) (int64, string, error) {
	var user models.User
	if err := ur.DB().Where("email = ?", email).First(&user).Error; err != nil {
		return 0, "", err
	}
	return user.ID, user.Password, nil
}

// Delete removes a user by ID.
func (ur *UserRepositoryImpl) Delete(id int64) error {
	return ur.DB().Delete(&models.User{}, id).Error
}

// Exists checks whether a user with the given email already exists.
func (ur *UserRepositoryImpl) Exists(email string) (bool, error) {
	var count int64
	if err := ur.DB().Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
