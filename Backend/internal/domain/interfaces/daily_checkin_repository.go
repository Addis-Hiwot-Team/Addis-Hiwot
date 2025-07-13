package interfaces

import (
	"addis-hiwot/internal/domain/models"
	"time"
)

type DailyCheckInRepository interface {
	CreateCheckIn(userID int, mood, note string, date time.Time) (*models.DailyCheckIn, error)
	GetCheckInsByUser(userID int) ([]*models.DailyCheckIn, error)
	GetCheckInByDate(userID int, date time.Time) (*models.DailyCheckIn, error)
}
