package usecases

import (
	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/models"
	"errors"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo     interfaces.UserRepository
	validate *validator.Validate
}

func NewUserUsecase(repo interfaces.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo:     repo,
		validate: validator.New(),
	}
}

func (uc *UserUsecase) Create(user *models.User) (*models.UserResponse, error) {
	// Validate struct fields using tags in the model
	if err := uc.validate.Struct(user); err != nil {
		return nil, err
	}

	// Hash the password before saving
	if user.PasswordHash == "" {
		return nil, errors.New("password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(hashedPassword)

	// Call repository to save user
	createdUser, err := uc.repo.Create(user)
	return createdUser.ToResponse(), err
}

func (uc *UserUsecase) GetAll() ([]*models.UserResponse, error) {
	return uc.repo.GetAll()
}
