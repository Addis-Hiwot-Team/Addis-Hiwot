package service

import (
	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTService(secret string, duration time.Duration) interfaces.JWTService {
	return &jwtService{
		secretKey:     secret,
		tokenDuration: duration,
	}
}

func (s *jwtService) GenerateToken(user *models.User) (string, error) {
	claims := models.UserClaims{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *jwtService) ValidateToken(tokenString string) (*models.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*models.UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
