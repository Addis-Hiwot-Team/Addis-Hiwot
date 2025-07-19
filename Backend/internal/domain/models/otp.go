package models

import "time"

type Otp struct {
	ID        uint
	UserID    uint
	Code      string `gorm:"index"`
	Type      string
	Exp       time.Time
	CreatedAt time.Time
}
