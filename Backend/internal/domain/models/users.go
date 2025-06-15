package models

import (
	"time"
)

type User struct {
	ID           int       `gorm:"primaryKey;column:id" json:"id"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex" json:"email" validate:"required,email"`
	Username     string    `gorm:"type:varchar(255)" json:"username" validate:"required,min=3,max=32"`
	Name         string    `gorm:"type:varchar(255)" json:"name"`
	ProfileImage string    `gorm:"type:varchar(255)" json:"profile_image"`
	PasswordHash string    `gorm:"type:varchar(255)" json:"password_hash" validate:"required,min=6"`
	IsActive     bool      `json:"is_active"`
	Role         string    `gorm:"type:varchar(255)" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	//relations
	DailyCheckIns     []*DailyCheckIn     `gorm:"constraint:OnDelete:CASCADE"`
	Goals             []*Goal             `gorm:"constraint:OnDelete:CASCADE"`
	GoalRewards       []*GoalReward       `gorm:"constraint:OnDelete:CASCADE"`
	AIChatHistories   []*AIChatHistory    `gorm:"constraint:OnDelete:CASCADE"`
	CommunityMessages []*CommunityMessage `gorm:"constraint:OnDelete:CASCADE"`
	UserResources     []*UserResource     `gorm:"constraint:OnDelete:CASCADE"`
	Habits            []*Habit            `gorm:"constraint:OnDelete:CASCADE"`
	HabitLogs         []*HabitLog         `gorm:"constraint:OnDelete:CASCADE"`
	UserQuotes        []*UserQuote        `gorm:"constraint:OnDelete:CASCADE"`
}

type UserResponse struct {
	ID           int       `json:"id"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	Name         string    `json:"name"`
	ProfileImage string    `json:"profile_image"`
	IsActive     bool      `json:"is_active"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// ToResponse converts User model to UserResponse struct
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:           u.ID,
		Email:        u.Email,
		Username:     u.Username,
		Name:         u.Name,
		ProfileImage: u.ProfileImage,
		IsActive:     u.IsActive,
		Role:         u.Role,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
