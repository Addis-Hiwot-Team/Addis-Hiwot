package usecases

import (
	"addis-hiwot/internal/domain/models"
	"addis-hiwot/internal/domain/schema"
	"addis-hiwot/internal/repository"
	"addis-hiwot/utils"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthUsecase interface {
	Login(identifier, password string) (string, error)
	Logout(token string) error
	Register(req schema.CreateUser) (*models.UserResponse, error)
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
	}
	err := au.ar.Create(user)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), err

	// user := &models.User{}
}
func (au *authUsecase) Login(identifier, password string) (string, error) {
	user, err := au.ar.GetUserByIdentifier(identifier)
	if err != nil {
		return "", err
	}
	log.Println("secret:", os.Getenv("JWT_AUTH_SECRET"))
	if err := utils.CheckPassword(user.PasswordHash, password); err != nil {
		return "", err
	}
	var authTokenExp = time.Now().Add(time.Hour * 24 * 7)
	claims := &schema.AuthClaim{
		UserID: uint(user.ID),
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(authTokenExp),
		},
	}
	token, err := utils.GenerateJwt(claims, os.Getenv("JWT_AUTH_SECRET"))

	au.sr.Create(token, uint(user.ID), authTokenExp)

	return token, err

}
func (au *authUsecase) Logout(token string) error {
	return au.sr.Delete(token)
}
