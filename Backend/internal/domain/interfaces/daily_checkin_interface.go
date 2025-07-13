package interfaces

import "addis-hiwot/internal/domain/models"

type DailyCheckInUsecase interface {
	AddCheckIn(userID int, mood, note string) (*models.DailyCheckIn, error)
	GetCheckInsByUser(userID int) ([]*models.DailyCheckIn, error)
}
