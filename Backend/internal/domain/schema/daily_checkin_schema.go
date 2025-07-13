package schema

import "time"

// DailyCheckinCreate represents the data for creating a new daily check-in.
type DailyCheckinCreate struct {
	Mood  string `json:"mood" binding:"required"`
	Notes string `json:"notes"`
}

// DailyCheckinResponse represents a daily check-in record.
type DailyCheckinResponse struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Mood        string    `json:"mood"`
	Notes       string    `json:"notes"`
	CheckInDate time.Time `json:"checkin_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// DailyCheckinUpdate represents the data for updating a daily check-in.
type DailyCheckinUpdate struct {
	Mood  string `json:"mood"`
	Notes string `json:"notes"`
}
