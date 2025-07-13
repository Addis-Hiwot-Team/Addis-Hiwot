package usecases

import (
	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/models"
	"errors"
	"time"
)

type DailyCheckInUsecase struct {
	repo interfaces.DailyCheckInRepository
}

func NewDailyCheckInUsecase(repo interfaces.DailyCheckInRepository) interfaces.DailyCheckInUsecase {
	return &DailyCheckInUsecase{repo: repo}
}

func (uc *DailyCheckInUsecase) AddCheckIn(userID int, mood, note string) (*models.DailyCheckIn, error) {
	today := time.Now().Truncate(24 * time.Hour)
	existing, err := uc.repo.GetCheckInByDate(userID, today)
	if err == nil && existing != nil {
		return nil, errors.New("already checked in today")
	}
	return uc.repo.CreateCheckIn(userID, mood, note, today)
}

func (uc *DailyCheckInUsecase) GetCheckInsByUser(userID int) ([]*models.DailyCheckIn, error) {
	return uc.repo.GetCheckInsByUser(userID)
}
