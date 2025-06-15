package models

import "time"

type CommunityMessage struct {
	ID          int    `gorm:"primaryKey"`
	UserID      int    `gorm:"not null"`
	Content     string `gorm:"type:text"`
	IsAnonymous bool
	PostedAt    time.Time
	IsEdited    bool
	IsFlagged   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE"`
}
