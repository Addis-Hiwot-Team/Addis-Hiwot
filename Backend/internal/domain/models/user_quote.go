package models

import "time"

type UserQuote struct {
	ID         int `gorm:"primaryKey"`
	UserID     int `gorm:"not null"`
	QuoteID    int `gorm:"not null"`
	Action     string
	ActionDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User  *User              `gorm:"constraint:OnDelete:CASCADE"`
	Quote *MotivationalQuote `gorm:"constraint:OnDelete:CASCADE"`
}
