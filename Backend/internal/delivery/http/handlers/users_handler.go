package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/schema"
	// "addis-hiwot/internal/usecases"
)

type UserHandler struct {
	uc interfaces.UserUsecaseInterface
}

  

func NewUserHandler(uc interfaces.UserUsecaseInterface) *UserHandler {
	return &UserHandler{uc: uc}
}

// @Summary      Get user by ID
// @Description  Get user details by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true "User ID"
// @Success      200  {object}  models.UserResponse
// @Failure      400  {object}  schema.APIError
// @Failure      404  {object}  schema.APIError
// @Failure      500  {object}  schema.APIError
// @Security     BearerAuth
// @Router       /users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	type GetUserByIDReq struct {
		ID int `uri:"id" binding:"required"`
	}
	var req GetUserByIDReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID is not integer"})
		return
	}
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user ID is required"})
		return
	}

	user, err := h.uc.GetByID(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
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
