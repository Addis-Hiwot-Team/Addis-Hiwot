package repository

import (
	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/models"
	"time"

	"gorm.io/gorm"
)

type DailyCheckInRepository struct {
	db *gorm.DB
}

func NewDailyCheckInRepository(db *gorm.DB) interfaces.DailyCheckInRepository {
	return &DailyCheckInRepository{db: db}
}

func (r *DailyCheckInRepository) CreateCheckIn(userID int, mood, note string, date time.Time) (*models.DailyCheckIn, error) {
	checkIn := &models.DailyCheckIn{
		UserID:      userID,
		Mood:        mood,
		Notes:       note,
		CheckInDate: date,
	}
	err := r.db.Create(checkIn).Error
	return checkIn, err
}

func (r *DailyCheckInRepository) GetCheckInsByUser(userID int) ([]*models.DailyCheckIn, error) {
	var checkIns []*models.DailyCheckIn
	err := r.db.Where("user_id = ?", userID).Order("check_in_date desc").Find(&checkIns).Error
	return checkIns, err
}

func (r *DailyCheckInRepository) GetCheckInByDate(userID int, date time.Time) (*models.DailyCheckIn, error) {
	var checkIn models.DailyCheckIn
	err := r.db.Where("user_id = ? AND check_in_date = ?", userID, date.Format("2006-01-02")).First(&checkIn).Error
	if err != nil {
		return nil, err
	}
	return &checkIn, nil
}
