package handlers

import (
	"addis-hiwot/internal/domain/schema"
	"addis-hiwot/internal/usecases"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auc usecases.AuthUsecase
}

func NewAuthHander(auc usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{auc}
}

func (ah *AuthHandler) Register(ctx *gin.Context) {
	var req schema.CreateUser
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := ah.auc.Register(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var req struct {
		Identifier string `json:"identifier" binding:"required,min=3"`
		Password   string `json:"password" binding:"required,min=3"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	token, err := ah.auc.Login(req.Identifier, req.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"auth_token": token})

}

func (ah *AuthHandler) Logout(ctx *gin.Context) {
	token := strings.Split(ctx.GetHeader("Authorization"), " ")[1]

	err := ah.auc.Logout(token)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
