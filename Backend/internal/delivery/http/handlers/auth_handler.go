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

type LoginReq struct {
	Identifier string `json:"identifier" binding:"required,min=3"`
	Password   string `json:"password" binding:"required,min=3"`
}

// @Summary      login request hanlder
// @Description  logs in user using their crediential
// @Tags         auth
// @Param        loginRequest   body      LoginReq  true "Login request body"
// @Success      200  {object}  schema.AuthTokenPair
// @Failure      401  {object}  ErrorResponse
// @Router       /auth/login [post]
func (ah *AuthHandler) Login(ctx *gin.Context) {
	var req LoginReq

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}
	tokens, err := ah.auc.Login(req.Identifier, req.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, tokens)

}

type RefreshReq struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// @Summary      logout request hanlder
// @Description  logs out user
// @Tags         auth
// @Success      204
// @Param		 logoutRequest  body 	RefreshReq  true "Logout request body"
// @Failure      401  {object}  ErrorResponse
// @Security	 BearerAuth
// @Router       /auth/logout [post]
func (ah *AuthHandler) Logout(ctx *gin.Context) {
	var req RefreshReq
	if ctx.BindJSON(&req) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if req.RefreshToken == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is required"})
		return
	}

	accessToken := strings.Split(ctx.GetHeader("Authorization"), " ")[1]

	err := ah.auc.Logout(accessToken, req.RefreshToken)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusNoContent)
}

// @Summary      refresh request handler
// @Description  refreshes the access token using the refresh token
// @Tags         auth
// @Param        refreshRequest   body      RefreshReq  true "Refresh request body"
// @Success      200  {object}  schema.AuthTokenPair
// @Failure      401  {object}  ErrorResponse
// @Router       /auth/refresh [post]
func (ah *AuthHandler) Refresh(ctx *gin.Context) {

	var req RefreshReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := ah.auc.Refresh(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tokens)
}

type OAuthCodereq struct {
	Provider    string `json:"provider"`
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

// @Summary      OAuth login request handler
// @Description  logs in user using OAuth provider
// @Tags         auth
// @Param 	  oauthCodeRequest   body      OAuthCodereq  true "OAuth Code request body"
// @Success      200  {object}  schema.AuthTokenPair
// @Failure      401  {object}  ErrorResponse
// @Router  /auth/oauth  [post]
func (h *AuthHandler) OAuthCodeLoginHandler(c *gin.Context) {

	var req OAuthCodereq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	tokens, err := h.auc.OAuthLoginWithCode(req.Provider, req.Code, req.RedirectURI)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}
