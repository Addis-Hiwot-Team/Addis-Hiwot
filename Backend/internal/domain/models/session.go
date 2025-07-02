package models

import "time"

type Session struct {
	RefreshToken string    `gorm:"primaryKey;size:512"`
	UserID       uint      `gorm:"index"`
	Exp          time.Time `gorm:"index"`
	CreatedAt    time.Time
}
