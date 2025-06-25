package models

import "time"

type MotivationalQuote struct {
	ID        int    `gorm:"primaryKey"`
	QuoteText string `gorm:"type:text"`
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time

	UserQuotes []*UserQuote `gorm:"foreignKey:QuoteID;constraint:OnDelete:CASCADE"`
}
