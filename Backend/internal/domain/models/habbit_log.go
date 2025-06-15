package models

import "time"

type HabitLog struct {
	ID          int       `gorm:"primaryKey"`
	HabitID     int       `gorm:"not null"`
	UserID      int       `gorm:"not null"`
	LogDate     time.Time `gorm:"type:date"`
	IsCompleted bool
	Notes       string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Habit *Habit `gorm:"constraint:OnDelete:CASCADE"`
	User  *User  `gorm:"constraint:OnDelete:CASCADE"`
}
