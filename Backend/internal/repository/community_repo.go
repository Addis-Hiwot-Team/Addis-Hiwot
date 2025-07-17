package repository

import (
	"addis-hiwot/internal/domain/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type communityRepository struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) *communityRepository {
	return &communityRepository{db: db}
}
// GetAllMessages retrieves up to 'limit' community messages posted before the specified time.
// It returns the messages, a boolean indicating if there are more messages, the timestamp for the next page, and an error if any.
// Pagination is handled by fetching 'limit + 1' messages and trimming the result if more exist.
func (r *communityRepository) GetAllMessages(ctx context.Context, limit int, before time.Time) ([]*models.CommunityMessage, bool, time.Time, error) {
	var messages []*models.CommunityMessage

	query := limit + 1

	err := r.db.WithContext(ctx).Where("PostedAt < ?", before).Order("PostedAt DESC").Limit(query).Preload("User").Find(&messages).Error
	if err != nil {
		return nil, false, time.Time{}, err
	}
	if len(messages) == 0 {
		return nil, false, time.Time{}, nil
	}
	hasMore := len(messages) > limit
	if hasMore {
		messages = messages[:limit]
	}
	var nextPageTime time.Time
	if len(messages) > 0 {
		nextPageTime = messages[len(messages)-1].PostedAt
	}
	return messages, hasMore, nextPageTime, nil
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
