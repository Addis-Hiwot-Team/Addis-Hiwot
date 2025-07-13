package interfaces

import (
	"addis-hiwot/internal/domain/models"
	"addis-hiwot/internal/domain/schema"

	"github.com/stretchr/testify/mock"
)

// UserUsecaseInterface defines the contract for user usecase methods
type UserUsecaseInterface interface {
	Register(input *schema.CreateUser) (string, error)
	Login(input *schema.LoginUser) (string, error)
	GetAll() ([]*models.UserResponse, error)
	ForgotPassword(email string) error
	ResetPassword(code, newPassword string) error
	ChangePassword(userID int, oldPassword, newPassword string) error
}

// MockUserUsecase is a mock implementation of UserUsecaseInterface
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(input *schema.CreateUser) (string, error) {
	args := m.Called(input)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) Login(input *schema.LoginUser) (string, error) {
	args := m.Called(input)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) GetAll() ([]*models.UserResponse, error) {
	args := m.Called()
	// Check if the first return argument is nil before casting.
	// This prevents a panic when the test expects an error.
	var r0 []*models.UserResponse
	if args.Get(0) != nil {
		r0 = args.Get(0).([]*models.UserResponse)
	}
	return r0, args.Error(1)
}

func (m *MockUserUsecase) ForgotPassword(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func (m *MockUserUsecase) ResetPassword(code, newPassword string) error {
	args := m.Called(code, newPassword)
	return args.Error(0)
}

func (m *MockUserUsecase) ChangePassword(userID int, oldPassword, newPassword string) error {
	args := m.Called(userID, oldPassword, newPassword)
	return args.Error(0)
}
