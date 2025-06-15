package schema

import "addis-hiwot/internal/domain/models"

type CreateUser struct {
	Email        string `json:"email" binding:"required,email"`
	Username     string `json:"username" binding:"required,min=3,max=32"`
	Name         string `json:"name" binding:"max=255,min=2"`
	Password     string `json:"password" binding:"required,min=6"`
	ProfileImage string `json:"profile_image" binding:"omitempty,max=255"`
}

func (cu *CreateUser) DBUser() *models.User {
	return &models.User{
		Name:         cu.Name,
		Email:        cu.Email,
		Username:     cu.Username,
		ProfileImage: cu.ProfileImage,
		PasswordHash: cu.Password, // hash the password before saving
	}
}
