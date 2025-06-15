package models

import "time"

type UserResource struct {
	ID         int `gorm:"primaryKey"`
	UserID     int `gorm:"not null"`
	ResourceID int `gorm:"not null"`
	Action     string
	ActionDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User     *User     `gorm:"constraint:OnDelete:CASCADE"`
	Resource *Resource `gorm:"constraint:OnDelete:CASCADE"`
}
