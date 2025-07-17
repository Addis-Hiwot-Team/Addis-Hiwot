package handlers

import (
	"addis-hiwot/internal/usecases"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CommunityMessageHandler struct {
	communityUsecase usecases.CommunityMessageUsecase
}

func NewCommunityMessageHandler(communityUsecase usecases.CommunityMessageUsecase) *CommunityMessageHandler {
	return &CommunityMessageHandler{
		communityUsecase: communityUsecase,
	}
}

func (h *CommunityMessageHandler) GetAllMessages(ctx *gin.Context){
	limitstr := ctx.Query("limit")
	limit, err := strconv.Atoi(limitstr)
	if err != nil || limit <= 0 {
		limit = 20 // Default limit
	}
	beforestr := ctx.Query("before")
	before, err := time.Parse(time.RFC3339, beforestr)
	if err != nil {
		before = time.Now()
	}

	messages, hasMore, nextPageTime, err := h.communityUsecase.GetAllMessages(ctx, limit, before)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Something went wrong","error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Community messages retrieved successfully",
		"data": gin.H{
			"messages":   messages,
			"has_more":   hasMore,
			"next_page":  nextPageTime,
		},
	})
}