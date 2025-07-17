package usecases

import (
	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/schema"
	"context"
	"time"
)

type CommunityMessageUsecase struct {
	communityRepo interfaces.CommunityRepository
	userRepo      interfaces.UserRepository
}

func NewCommunityMessageUsecase(communityRepo interfaces.CommunityRepository, userRepo interfaces.UserRepository) *CommunityMessageUsecase {
	return &CommunityMessageUsecase{
		communityRepo: communityRepo,
		userRepo:      userRepo,
	}
}

func(usecase *CommunityMessageUsecase) GetAllMessages(ctx context.Context, limit int, before time.Time) ([]*schema.CommunityMessageResponse, bool, time.Time, error) {
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