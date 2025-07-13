package handlers

import (
	"net/http"

	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/schema"

	"github.com/gin-gonic/gin"
)

type DailyCheckInHandler struct {
	uc interfaces.DailyCheckInUsecase
}

func NewDailyCheckInHandler(uc interfaces.DailyCheckInUsecase) *DailyCheckInHandler {
	return &DailyCheckInHandler{uc: uc}
}

// POST /api/v1/checkin
func (h *DailyCheckInHandler) AddCheckIn(c *gin.Context) {
	claim, exists := c.Get("claim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	authClaim, ok := claim.(*schema.AuthClaim)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claim type"})
		return
	}

	var req struct {
		Mood  string `json:"mood" binding:"required"`
		Notes string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkIn, err := h.uc.AddCheckIn(int(authClaim.UserID), req.Mood, req.Notes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, checkIn)
}

// GET /api/v1/checkin
func (h *DailyCheckInHandler) GetCheckIns(c *gin.Context) {
	claim, exists := c.Get("claim")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	authClaim, ok := claim.(*schema.AuthClaim)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claim type"})
		return
	}

	checkIns, err := h.uc.GetCheckInsByUser(int(authClaim.UserID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, checkIns)
}
