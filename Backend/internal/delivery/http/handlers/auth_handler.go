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

// @Summary      register request hanlder
// @Description  registers a new user
// @Param        registerRequest   body      schema.CreateUser  true "Register request body"
// @Tags         auth
// @Success      201  {object}  models.UserResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /auth/register [post]
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

type ErrorResponse struct {
	Error string `json:"error"`
}
type LoginResponse struct {
	AuthToken string `json:"auth_token"`
}

type LoginReq struct {
	Identifier string `json:"identifier" binding:"required,min=3"`
	Password   string `json:"password" binding:"required,min=3"`
}

// @Summary      login request hanlder
// @Description  logs in user using their crediential
// @Tags         auth
// @Param        loginRequest   body      LoginReq  true "Login request body"
// @Success      200  {object}  LoginResponse
// @Failure      401  {object}  ErrorResponse
// @Router       /auth/login [post]
func (ah *AuthHandler) Login(ctx *gin.Context) {
	var req LoginReq

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

// @Summary      logout request hanlder
// @Description  logs out user
// @Tags         auth
// @Success      204
// @Failure      401  {object}  ErrorResponse
// @Router       /auth/logout [post]
func (ah *AuthHandler) Logout(ctx *gin.Context) {
	token := strings.Split(ctx.GetHeader("Authorization"), " ")[1]

	err := ah.auc.Logout(token)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}
