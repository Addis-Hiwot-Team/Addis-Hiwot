package schema

import (
	"addis-hiwot/internal/domain/models"
	"time"
)

type CreateCommunityMessage struct {
	Content     string `json:"content" binding:"required,min=1,max=5000"`
	IsAnonymous bool   `json:"is_anonymous"`
}

func (cm *CreateCommunityMessage) ToModel(userID int) *models.CommunityMessage {
	return &models.CommunityMessage{
		UserID:      userID,
		Content:     cm.Content,
		IsAnonymous: cm.IsAnonymous,
		PostedAt:    time.Now(),
		CreatedAt:   time.Now(),
	}
}

type CommunityMessageResponse struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Username    string    `json:"username,omitempty"` // if not anonymous
	Content     string    `json:"content"`
	IsAnonymous bool      `json:"is_anonymous"`
	IsEdited    bool      `json:"is_edited"`
	IsFlagged   bool      `json:"is_flagged"`
	PostedAt    time.Time `json:"posted_at"`
}

func ToCommunityMessageResponse(msg *models.CommunityMessage, username string) *CommunityMessageResponse {
	return &CommunityMessageResponse{
		ID:          msg.ID,
		UserID:      msg.UserID,
		Username:    username,
		Content:     msg.Content,
		IsAnonymous: msg.IsAnonymous,
		IsEdited:    msg.IsEdited,
		IsFlagged:   msg.IsFlagged,
		PostedAt:    msg.PostedAt,
	}
}

type FlagMessageRequest struct {
	MessageID int    `json:"message_id" binding:"required"`
	Reason    string `json:"reason" binding:"required,min=5,max=1000"`
}
