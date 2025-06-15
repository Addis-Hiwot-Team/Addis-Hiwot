package interfaces

import "addis-hiwot/internal/domain/models"

type UserRepository interface {
	Create(user *models.User) error
	GetAll() ([]models.UserResponse, error)
}
