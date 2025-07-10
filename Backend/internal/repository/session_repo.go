package repository

import (
	"addis-hiwot/internal/domain/models"
	"time"

	"gorm.io/gorm"
)

type SessionRepository interface {
	Create(token string, userID uint, exp time.Time) error
	Get(token string) (*models.Session, error)
	Delete(token string) error
	Blacklist(token string, exp time.Time) error
	IsBlacklisted(token string) (bool, error)
}

type sessionRepo struct{ db *gorm.DB }

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepo{db}
}

func (sr *sessionRepo) Create(token string, userID uint, exp time.Time) error {
	return sr.db.Create(&models.Session{
		RefreshToken: token,
		UserID:       userID,
		Exp:          exp,
	}).Error
}
func (sr *sessionRepo) Get(token string) (*models.Session, error) {
	var session models.Session
	return &session, sr.db.Where("refresh_token = ?", token).First(&session).Error
}
func (sr *sessionRepo) Delete(token string) error {
	return sr.db.Delete(&models.Session{}, "refresh_token = ?", token).Error
}

func (r *sessionRepo) Blacklist(token string, exp time.Time) error {
	return r.db.Create(&models.BlacklistedToken{
		Token:     token,
		Exp:       exp,
		CreatedAt: time.Now(),
	}).Error
}

func (r *sessionRepo) IsBlacklisted(token string) (bool, error) {
	var b models.BlacklistedToken
	err := r.db.First(&b, "token = ?", token).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	return err == nil, err
}
