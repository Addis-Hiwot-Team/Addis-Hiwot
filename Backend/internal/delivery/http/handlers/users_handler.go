package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"addis-hiwot/internal/domain/schema"
	"addis-hiwot/internal/usecases"
)

type UserHandler struct {
	uc *usecases.UserUsecase
}

func NewUserHandler(uc *usecases.UserUsecase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userCreateSchema schema.CreateUser
	if err := c.ShouldBindJSON(&userCreateSchema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.uc.Register(&userCreateSchema)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "User registered successfully",
		"access_token": token,
	})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var loginSchema schema.LoginUser
	if err := c.ShouldBindJSON(&loginSchema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.uc.Login(&loginSchema)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user logged in successfully", "access_token": token})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
