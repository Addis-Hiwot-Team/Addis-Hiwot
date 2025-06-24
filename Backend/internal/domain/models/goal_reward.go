package models

import "time"

type GoalReward struct {
	ID                int `gorm:"primaryKey"`
	GoalID            int `gorm:"not null"`
	UserID            int `gorm:"not null"`
	RewardType        string
	RewardDescription string `gorm:"type:text"`
	AwardedAt         time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time

	Goal *Goal `gorm:"constraint:OnDelete:CASCADE"`
	User *User `gorm:"constraint:OnDelete:CASCADE"`
}
