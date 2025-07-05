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

type ForgotPasswordReq struct {
	Email string `json:"email" binding:"required,email"`
}

// @Summary      Forgot password
// @Description  Request a password reset code
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        forgotPasswordRequest body ForgotPasswordReq true "Forgot Password Request"
// @Success      200 {object} schema.APIMessage
// @Failure      400 {object} schema.APIError
// @Failure      500 {object} schema.APIError
// @Router       /users/forgot_password [post]
func (h *UserHandler) ForgotPassword(ctx *gin.Context) {
	var req ForgotPasswordReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.uc.ForgotPassword(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "email has been sent to the provided address!"})
}

type ResetPasswordReq struct {
	Code        string `json:"code" binding:"required,len=6"`
	NewPassword string `json:"new_password" binding:"required,min=3,max=32"`
}

// @Summary      Reset user password
// @Description  Reset user password using OTP code
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        resetPasswordRequest body ResetPasswordReq true "Reset Password Request"
// @Success      200 {object} schema.APIMessage
// @Failure      400 {object} schema.APIError
// @Failure      500 {object} schema.APIError
// @Router       /users/reset_password [post]
func (h *UserHandler) ResetPassword(ctx *gin.Context) {
	var req ResetPasswordReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.uc.ResetPassword(req.Code, req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "password reset success!"})
}

type ChangePasswordReq struct {
	Password    string `json:"password" binding:"required,min=3,max=32"`
	NewPassword string `json:"new_password" binding:"required,min=3,max=32"`
}

// @Summary      Change user password
// @Description  Change user password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        changePasswordRequest body ChangePasswordReq true "Change Password Request"
// @Success      200 {object} schema.APIMessage
// @Failure      400 {object} schema.APIError
// @Failure      500 {object} schema.APIError
// @Security     BearerAuth
// @Router       /users/change_password [post]
func (h *UserHandler) ChangePassword(ctx *gin.Context) {
	var req ChangePasswordReq
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claim := ctx.MustGet("claim").(*schema.AuthClaim)

	err := h.uc.ChangePassword(int(claim.UserID), req.Password, req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, schema.APIMessage{Message: "password changed successfully"})
}
