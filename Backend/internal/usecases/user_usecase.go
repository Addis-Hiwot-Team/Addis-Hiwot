package usecases

import (
	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/models"
	"addis-hiwot/internal/domain/schema"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo       interfaces.UserRepository
	jwtService interfaces.JWTService
}

func NewUserUsecase(r interfaces.UserRepository, j interfaces.JWTService) *UserUsecase {
	return &UserUsecase{
		repo:       r,
		jwtService: j,
	}
}

func (uc *UserUsecase) Register(input *schema.CreateUser) (string, error) {
	_, err := uc.repo.GetByEmail(input.Email)
	if err == nil {
		return "", errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := input.DBUser()
	user.PasswordHash = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	newUser, err := uc.repo.Create(user)

	if err != nil {
		return "", err
	}

	return uc.jwtService.GenerateToken(newUser)
}

func (uc *UserUsecase) Login(input *schema.LoginUser) (string, error) {
	user, err := uc.repo.GetByEmail(input.Email)
	if err != nil || user == nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	return uc.jwtService.GenerateToken(user)
}

func (uc *UserUsecase) GetAll() ([]*models.UserResponse, error) {
	return uc.repo.GetAll()
}
