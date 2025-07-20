package usecases

import (
	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/models"
	"addis-hiwot/internal/domain/schema"
	"addis-hiwot/internal/repository"
	"addis-hiwot/internal/service"
	"addis-hiwot/utils"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase interface {
	GetMe(userID int) (*models.UserResponse, error)
	Login(identifier, password string) (*schema.AuthResponse, error)
	Logout(accessToken, refreshToken string) error
	Register(req schema.CreateUser) (*schema.AuthResponse, error)
	Refresh(refreshToken string) (*schema.AuthTokenPair, error)
	ActivateAccount(token string) error
	OAuthLoginWithCode(provider, code, redirectURI string) (*schema.AuthTokenPair, error)
}
type authUsecase struct {
	ar        repository.AuthRepository
	userRepo  interfaces.UserRepository
	sr        repository.SessionRepository
	oa        *service.OAuthService // Optional, if you want to add OAuth support later
	otpRepo   repository.OtpRepo
	emailSrvs service.EmailService
}

func NewAuthUsecase(
	ar repository.AuthRepository,
	sr repository.SessionRepository,
	or repository.OtpRepo,
	ur interfaces.UserRepository,
	es service.EmailService,
) *authUsecase {
	return &authUsecase{ar: ar, sr: sr, oa: service.NewOAuthService(), otpRepo: or, emailSrvs: es, userRepo: ur}
}

// @Summary      Get current user info
// @Description  Get the currently authenticated user's information
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UserResponse
// @Failure      401  {object}  schema.APIError
// @Failure      500  {object}  schema.APIError
// @Security     BearerAuth
// @Router       /auth/me [get]
func (au *authUsecase) GetMe(userID int) (*models.UserResponse, error) {
	user, err := au.userRepo.Get(userID)
	if err != nil {
		return nil, err
	}
	return user.ToResponse(), nil

}

func (au *authUsecase) Register(req schema.CreateUser) (*schema.AuthResponse, error) {

	hashedPwd, _ := utils.HashPassword(req.Password)
	user := &models.User{
		Name:         req.Name,
		Username:     req.Username,
		Email:        req.Email,
		ProfileImage: req.ProfileImage,
		PasswordHash: hashedPwd,
		Role:         "user", // default
	}
	err := au.ar.Create(user)
	if err != nil {
		return nil, err
	}

	// Generate tokens for the new user
	tokens, err := au.generateTokens(user, false)
	if err != nil {
		return nil, err
	}

	resp := &schema.AuthResponse{
		User:   user.ToResponse(),
		Tokens: tokens,
	}
	//send verification email
	code := utils.GenerateOTP()
	otp := &models.Otp{
		UserID: uint(user.ID),
		Code:   code,
		Type:   "email_verification",
		Exp:    time.Now().Add(24 * time.Hour), // 24 hours
	}
	if _, err := au.otpRepo.Create(otp); err != nil {
		log.Println("failed to create otp:", err)
		return nil, fmt.Errorf("failed to create otp")
	}
	err = au.emailSrvs.SendEmail(user.Email, "Activate Account", "activate_account.html", map[string]any{
		"ActivationLink": os.Getenv("FRONTEND_URL") + "/activate?token=" + code,
		"CurrentYear":    time.Now().Year(),
		"WebsiteLink":    "https://www.addishiwt.com",
		"SupportLink":    "mailto:support@addishiwt.com",
	})

	if err != nil {
		log.Println("failed to send verification email:", err)
	}
	return resp, nil

}
func (au *authUsecase) Login(identifier, password string) (*schema.AuthResponse, error) {
	user, err := au.ar.GetUserByIdentifier(identifier)
	if err != nil {
		return nil, err
	}
	log.Println("secret:", os.Getenv("JWT_AUTH_SECRET"))
	if err := utils.CheckPassword(user.PasswordHash, password); err != nil {
		return nil, err
	}

	// Generate tokens for the user
	tokens, err := au.generateTokens(user, false)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}
	resp := &schema.AuthResponse{
		User:   user.ToResponse(),
		Tokens: tokens,
	}
	return resp, nil
}
func (au *authUsecase) Logout(accessToken, refreshToken string) error {
	// Blacklist access token
	var claims schema.AuthClaim
	err := utils.ParseJwt(&claims, accessToken, os.Getenv("JWT_AUTH_SECRET"))
	if err == nil {
		_ = au.sr.Blacklist(accessToken, claims.ExpiresAt.Time)
	}

	// Delete refresh session
	return au.sr.Delete(refreshToken)
}
func (au *authUsecase) Refresh(refreshToken string) (*schema.AuthTokenPair, error) {
	// Validate refresh token signature & expiry
	var claims schema.AuthClaim
	err := utils.ParseJwt(&claims, refreshToken, os.Getenv("JWT_REFRESH_SECRET"))
	if err != nil {
		return nil, err
	}

	// Check session
	session, err := au.sr.Get(refreshToken)
	if err != nil || session == nil {
		return nil, fmt.Errorf("invalid refresh token")
	}

	// Issue new access token
	accessExp := time.Now().Add(time.Minute * 15)
	accessClaims := &schema.AuthClaim{
		UserID: claims.UserID,
		Role:   claims.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExp),
		},
	}
	accessToken, err := utils.GenerateJwt(accessClaims, os.Getenv("JWT_AUTH_SECRET"))
	if err != nil {
		return nil, err
	}

	return &schema.AuthTokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken, // Reuse existing
	}, nil
}
func (au *authUsecase) OAuthLoginWithCode(provider, code, redirectURI string) (*schema.AuthTokenPair, error) {
	var idToken string
	var userInfo *service.GoogleUser
	var err error

	switch provider {
	case "google":
		// 1. Exchange code for id_token
		idToken, err = au.oa.ExchangeGoogleCodeForIDToken(code, redirectURI)
		if err != nil {
			return nil, fmt.Errorf("exchange failed: %w", err)
		}

		// 2. Verify id_token
		userInfo, err = au.oa.VerifyGoogleIDToken(idToken)
		if err != nil {
			return nil, fmt.Errorf("verify failed: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider)
	}

	// 3. Find or create user in DB
	user, err := au.ar.GetUserByIdentifier(userInfo.Email)
	if err != nil || user == nil {
		user = &models.User{
			Email:        userInfo.Email,
			Name:         userInfo.Name,
			ProfileImage: userInfo.Picture,
			Role:         "user",
			IsActive:     true,
		}
		if err := au.ar.Create(user); err != nil {
			return nil, fmt.Errorf("failed to create user: %w", err)
		}
	}

	// Generate tokens for the user
	return au.generateTokens(user, false)
}

// Extracted helper function for token generation
// if refresh is true it will assume you want to reuse  refresh token will be created
func (au *authUsecase) generateTokens(user *models.User, refresh bool) (*schema.AuthTokenPair, error) {
	tokens := &schema.AuthTokenPair{}
	// Access token (short-lived)
	accessExp := time.Now().Add(15 * time.Hour)
	accessClaims := &schema.AuthClaim{
		UserID: uint(user.ID),
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExp),
		},
	}
	accessToken, err := utils.GenerateJwt(accessClaims, os.Getenv("JWT_AUTH_SECRET"))
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// Refresh token (long-lived)
	refreshExp := time.Now().Add(7 * 24 * time.Hour)
	refreshClaims := &schema.AuthClaim{
		UserID: uint(user.ID),
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExp),
		},
	}
	if !refresh {

		refreshToken, err := utils.GenerateJwt(refreshClaims, os.Getenv("JWT_REFRESH_SECRET"))
		if err != nil {
			return nil, fmt.Errorf("failed to generate refresh token: %w", err)
		}

		// Store refresh session
		if err := au.sr.Create(refreshToken, uint(user.ID), refreshExp); err != nil {
			return nil, fmt.Errorf("failed to store refresh session: %w", err)
		}
		tokens.RefreshToken = refreshToken
	}

	tokens.AccessToken = accessToken
	return tokens, nil
}

func (au *authUsecase) ActivateAccount(token string) error {
	// Validate OTP
	otp, err := au.otpRepo.Get(token, "email_verification")
	if err != nil {
		return fmt.Errorf("invalid activation token: %w", err)
	}
	if otp.Exp.Before(time.Now()) {
		return fmt.Errorf("activation token expired")
	}

	// Activate user account
	err = au.userRepo.Activate(int(otp.UserID))
	if err != nil {
		return fmt.Errorf("faild to activate user: %w", err)
	}
	// Delete OTP after successful activation
	if err := au.otpRepo.Delete(otp.ID); err != nil {
		log.Printf("failed to delete OTP after activation: %v", err)
	}

	return nil
}
