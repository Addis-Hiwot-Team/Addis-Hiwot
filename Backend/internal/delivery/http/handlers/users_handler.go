package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"addis-hiwot/internal/domain/models"
	"addis-hiwot/internal/usecases"
)

type UserHandler struct {
	uc *usecases.UserUsecase
}

func NewUserHandler(uc *usecases.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
