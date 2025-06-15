package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID       int            `gorm:"primaryKey;column:user_id" json:"user_id"`
	Email        string         `gorm:"type:varchar(255);uniqueIndex" json:"email" validate:"required,email"`
	Username     string         `gorm:"type:varchar(255)" json:"username" validate:"required,min=3,max=32"`
	PasswordHash string         `gorm:"type:varchar(255)" json:"password_hash" validate:"required,min=6"`
	IsActive     bool           `json:"is_active"`
	Role         string         `gorm:"type:varchar(255)" json:"role" validate:"required"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type UserResponse struct {
	UserID    int       `json:"user_id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	IsActive  bool      `json:"is_active"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToResponse converts User model to UserResponse struct
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		UserID:    u.UserID,
		Email:     u.Email,
		Username:  u.Username,
		IsActive:  u.IsActive,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
