package models

import "time"

type Habit struct {
	ID           int `gorm:"primaryKey"`
	UserID       int `gorm:"not null"`
	Name         string
	Frequency    string
	ReminderTime *time.Time
	StreakCount  int
	CreatedAt    time.Time
	UpdatedAt    time.Time

	User      *User       `gorm:"constraint:OnDelete:CASCADE"`
	HabitLogs []*HabitLog `gorm:"constraint:OnDelete:CASCADE"`
}
