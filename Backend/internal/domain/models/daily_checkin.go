package models

import "time"

type DailyCheckIn struct {
	ID          int `gorm:"primaryKey"`
	UserID      int `gorm:"not null"`
	Mood        string
	Notes       string    `gorm:"type:text"`
	CheckInDate time.Time `gorm:"type:date"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User *User `gorm:"constraint:OnDelete:CASCADE"`
}
