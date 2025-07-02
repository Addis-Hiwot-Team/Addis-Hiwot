package usecases

import (
	"addis-hiwot/internal/domain/models"
	"addis-hiwot/internal/domain/schema"
	"addis-hiwot/internal/repository"
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
}
type authUsecase struct {
	ar repository.AuthRepository
	sr repository.SessionRepository
}

func NewAuthUsecase(
	ar repository.AuthRepository,
	sr repository.SessionRepository) *authUsecase {
	return &authUsecase{ar, sr}
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

	return user.ToResponse(), err

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
	accessExp := time.Now().Add(time.Minute * 15)
	accessClaims := &schema.AuthClaim{
		UserID: uint(user.ID),
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExp),
		},
	}
	accessToken, err := utils.GenerateJwt(accessClaims, os.Getenv("JWT_AUTH_SECRET"))
	if err != nil {
		return nil, err
	}
	// Refresh token (long-lived)
	refreshExp := time.Now().Add(time.Hour * 24 * 7)
	refreshClaims := &schema.AuthClaim{
		UserID: uint(user.ID),
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExp),
		},
	}
	refreshToken, err := utils.GenerateJwt(refreshClaims, os.Getenv("JWT_REFRESH_SECRET"))
	if err != nil {
		return nil, err
	}
	// Store refresh session
	if err := au.sr.Create(refreshToken, uint(user.ID), refreshExp); err != nil {
		return nil, err
	}
	return &schema.AuthTokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

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
