package repository

import (
	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// GetAll returns users as []UserResponse (only selected fields)
func (r *userRepository) GetAll() ([]models.UserResponse, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	// Convert to []UserResponse
	userResponses := make([]models.UserResponse, len(users))
	for i, u := range users {
		userResponses[i] = u.ToResponse()
	}

	return userResponses, nil
}
