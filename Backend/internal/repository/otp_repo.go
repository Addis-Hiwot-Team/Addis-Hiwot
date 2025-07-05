package repository

import (
	"addis-hiwot/internal/domain/models"

	"gorm.io/gorm"
)

type OtpRepo interface {
	Create(otp *models.Otp) (*models.Otp, error)
	Get(code, otpType string) (*models.Otp, error)
	Delete(id uint) error
}

type otpRepo struct {
	db *gorm.DB
}

// Create implements OtpRepo.
func (o *otpRepo) Create(otp *models.Otp) (*models.Otp, error) {
	err := o.db.Create(otp).Error
	return otp, err
}

// Delete implements OtpRepo.
func (o *otpRepo) Delete(id uint) error {
	return o.db.Delete(&models.Otp{}, id).Error
}

// Get implements OtpRepo.
func (o *otpRepo) Get(code string, otpType string) (*models.Otp, error) {
	var otp models.Otp
	err := o.db.Where("code = ? AND type = ?", code, otpType).First(&otp).Error

	return &otp, err

}

func NewOtpRepo(db *gorm.DB) OtpRepo {
	return &otpRepo{db: db}
}
