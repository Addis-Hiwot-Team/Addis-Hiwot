package interfaces

import "addis-hiwot/internal/domain/models"

type JWTService interface {
	GenerateToken(user *models.User) (string, error)
	ValidateToken(token string) (*models.UserClaims, error)
}
