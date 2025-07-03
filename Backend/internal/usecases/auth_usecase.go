package usecases

import (
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
	Login(identifier, password string) (*schema.AuthTokenPair, error)
	Logout(accessToken, refreshToken string) error
	Register(req schema.CreateUser) (*models.UserResponse, error)
	Refresh(refreshToken string) (*schema.AuthTokenPair, error)
	OAuthLoginWithCode(provider, code, redirectURI string) (*schema.AuthTokenPair, error)
}
type authUsecase struct {
	ar repository.AuthRepository
	sr repository.SessionRepository
	oa *service.OAuthService // Optional, if you want to add OAuth support later
}

func NewAuthUsecase(
	ar repository.AuthRepository,
	sr repository.SessionRepository) *authUsecase {
	return &authUsecase{ar: ar, sr: sr, oa: service.NewOAuthService()}
}

func (au *authUsecase) Register(req schema.CreateUser) (*models.UserResponse, error) {

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
	_, err = au.generateTokens(user, false)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil

	// user := &models.User{}
}
func (au *authUsecase) Login(identifier, password string) (*schema.AuthTokenPair, error) {
	user, err := au.ar.GetUserByIdentifier(identifier)
	if err != nil {
		return nil, err
	}
	log.Println("secret:", os.Getenv("JWT_AUTH_SECRET"))
	if err := utils.CheckPassword(user.PasswordHash, password); err != nil {
		return nil, err
	}

	// Generate tokens for the user
	return au.generateTokens(user, false)
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
	var tokens *schema.AuthTokenPair
	// Access token (short-lived)
	accessExp := time.Now().Add(15 * time.Minute)
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

	tokens = &schema.AuthTokenPair{
		AccessToken: accessToken,
	}
	return tokens, nil
}
