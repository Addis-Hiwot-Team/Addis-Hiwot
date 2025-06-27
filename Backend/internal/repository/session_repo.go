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
}

type sessionRepo struct{ db *gorm.DB }

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepo{db}
}

func (sr *sessionRepo) Create(token string, userID uint, exp time.Time) error {
	return sr.db.Create(&models.Session{
		Token:  token,
		UserID: userID,
		Exp:    exp,
	}).Error
}
func (sr *sessionRepo) Get(token string) (*models.Session, error) {
	var session models.Session
	return &session, sr.db.Where("token = ?", token).First(&session).Error
}
func (sr *sessionRepo) Delete(token string) error {
	return sr.db.Delete(&models.Session{}, "token = ?", token).Error
}
