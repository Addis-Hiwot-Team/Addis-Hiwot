package handlers

import (
	"addis-hiwot/internal/domain/schema"
	"addis-hiwot/internal/usecases"
	"addis-hiwot/utils"
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

func (ah *AuthHandler) GetMe(ctx *gin.Context) {
	authClaim := ctx.MustGet("claim").(*schema.AuthClaim)

	user, err := ah.auc.GetMe(int(authClaim.UserID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
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
	data, err := ah.auc.Register(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, schema.APIResponse{
		Message: "User registered successfully",
		Data:    data,
	})
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=3"`
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
		ctx.JSON(400, gin.H{"error": utils.ValidationErrorToText(err, req)})
		return
	}
	tokens, err := ah.auc.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, schema.APIResponse{
		Message: "Login successful",
		Data:    tokens})

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

// @Summary 	Activate account request handler
// @Description 	Activates user account using activation code
// @Tags 		auth
// @Param 		code path string true "Activation code"
// @Success 	200 {object} schema.APIMessage
// @Failure 	400 {object} ErrorResponse
// @Failure 	500 {object} ErrorResponse
// @Router 	/auth/activate/{code} [get]
func (h *AuthHandler) ActivateAccount(c *gin.Context) {
	activationCode := c.Param("code")
	if activationCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "activation code is required"})
		return
	}

	err := h.auc.ActivateAccount(activationCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "account activated successfully"})
}
