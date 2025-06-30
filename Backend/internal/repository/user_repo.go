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

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	return user, err
}

// GetAll returns users as []UserResponse (only selected fields)
func (r *userRepository) GetAll() ([]*models.UserResponse, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	// Convert to []UserResponse
	userResponses := make([]*models.UserResponse, len(users))
	for i, u := range users {
		userResponses[i] = u.ToResponse()
	}

	return userResponses, nil
}

// GetUserByID returns a user by ID as UserResponse
func (r *userRepository) GetUserByID(id int) (*models.UserResponse, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

func (r *userRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
