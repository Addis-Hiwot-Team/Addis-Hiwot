package interfaces

import "addis-hiwot/internal/domain/models"

type UserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetAll() ([]*models.UserResponse, error)
	GetUserByID(id int) (*models.UserResponse, error)
	Get(id int) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	ChangePassword(id int, newpassowrd string) error
	Activate(id int) error
	IsActive(id int) bool
}
