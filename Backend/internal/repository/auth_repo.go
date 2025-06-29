package repository

import (
	"addis-hiwot/internal/domain/models"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}
type AuthRepository interface {
	GetUserByIdentifier(identifier string) (*models.User, error)
	Create(user *models.User) error
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepo{db}
}
func (ar *authRepo) GetUserByIdentifier(identifier string) (*models.User, error) {
	var user models.User
	err := ar.db.Where("email = ? OR username = ?", identifier, identifier).First(&user).Error
	return &user, err
}

func (ar *authRepo) Create(user *models.User) error {
	return ar.db.Create(&user).Error
}
