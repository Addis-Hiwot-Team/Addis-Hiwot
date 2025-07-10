package repository

import (
	"addis-hiwot/internal/domain/models"

	"gorm.io/gorm"
)

type communityRepository struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) *communityRepository {
	return &communityRepository{db: db}
}
func (r *communityRepository) GetAllMessages() ([]*models.CommunityMessage, error) {
	var messages []*models.CommunityMessage
	err := r.db.Preload("User").Order("posted_at DESC").Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *communityRepository) CreateMessage(message *models.CommunityMessage) (*models.CommunityMessage, error) {
	err := r.db.Create(message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *communityRepository) GetMessageByID(id int) (*models.CommunityMessage, error) {
	var message models.CommunityMessage
	err := r.db.Preload("User").First(&message, id).Error
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (r *communityRepository) UpdateMessage(message *models.CommunityMessage) (*models.CommunityMessage, error) {
	err := r.db.Save(message).Error
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *communityRepository) DeleteMessage(id int) error {
	return r.db.Delete(&models.CommunityMessage{}, id).Error
}
