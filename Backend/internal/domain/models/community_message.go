package models

import "time"

type CommunityMessage struct {
	ID          int    `gorm:"primaryKey"`
	UserID      int    `gorm:"not null"`
	Content     string `gorm:"type:text"`
	IsAnonymous bool   `gorm:"default:false"`
	PostedAt    time.Time
	IsEdited    bool `gorm:"default:false"`
	IsFlagged   bool `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE"`
}
