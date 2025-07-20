package interfaces

import (
	"addis-hiwot/internal/domain/models"
	"context"
	"time"
)

type CommunityRepository interface {
	CreateMessage(message *models.CommunityMessage) (*models.CommunityMessage, error)
	GetAllMessages(ctx context.Context, limit int, before time.Time) ([]*models.CommunityMessage, bool, time.Time, error)
	GetMessageByID(id int) (*models.CommunityMessage, error)
	
}