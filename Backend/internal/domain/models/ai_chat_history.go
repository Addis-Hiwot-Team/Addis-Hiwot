package models

import "time"

type AIChatHistory struct {
	ID            int    `gorm:"primaryKey"`
	UserID        int    `gorm:"not null"`
	Message       string `gorm:"type:text"`
	IsUserMessage bool
	SentAt        time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE"`
}
