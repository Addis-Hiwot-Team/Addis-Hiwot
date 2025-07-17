package usecases

import (
	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/schema"
	"context"
	"time"
)

type MessageUsecase struct {
	communityRepo interfaces.CommunityRepository
	userRepo      interfaces.UserRepository
}

func NewMessageUsecase(communityRepo interfaces.CommunityRepository, userRepo interfaces.UserRepository) *MessageUsecase {
	return &MessageUsecase{
		communityRepo: communityRepo,
		userRepo:      userRepo,
	}
}

func(usecase *MessageUsecase) GetAllMessages(ctx context.Context, limit int, before time.Time) ([]*schema.CommunityMessageResponse, bool, time.Time, error) {
	message, hasMore, nextPageTime, err := usecase.communityRepo.GetAllMessages(ctx, limit, before)
	if err != nil {
		return nil, false, time.Time{}, err
	}

	var responses []*schema.CommunityMessageResponse
	for _,msg := range message {
		if msg.IsFlagged {
			continue // Skip flagged messages
		}

		var username string
		if msg.IsAnonymous {
			username = "Anonymous"
		} else {
			if msg.User != nil {
				username = msg.User.Username
			} else {
				user, err := usecase.userRepo.GetUserByID(msg.UserID)
				if err != nil {
					return nil, false, time.Time{}, err
				}
				if user != nil {
					username = user.Username
				} else {
					username = "Unknown User"
				}
			}
		}
		response := schema.ToCommunityMessageResponse(msg, username)
		responses = append(responses, response)
	}
	return responses, hasMore, nextPageTime, nil
}