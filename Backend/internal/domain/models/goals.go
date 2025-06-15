package models

import "time"

type Goal struct {
	ID             int `gorm:"primaryKey"`
	UserID         int `gorm:"not null"`
	Title          string
	Description    string    `gorm:"type:text"`
	Deadline       time.Time `gorm:"type:date"`
	Frequency      string
	IsCompleted    bool
	CompletionDate *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time

	User        *User         `gorm:"constraint:OnDelete:CASCADE"`
	GoalRewards []*GoalReward `gorm:"constraint:OnDelete:CASCADE"`
}
