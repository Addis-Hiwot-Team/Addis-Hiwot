package models

import "time"

type Resource struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Description string `gorm:"type:text"`
	Type        string
	Topic       string
	SourceURL   string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	UserResources []*UserResource `gorm:"constraint:OnDelete:CASCADE"`
}
