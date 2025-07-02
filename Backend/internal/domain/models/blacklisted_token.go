package models

import "time"

type BlacklistedToken struct {
	Token     string `gorm:"primaryKey;size:512"`
	Exp       time.Time
	CreatedAt time.Time
}
