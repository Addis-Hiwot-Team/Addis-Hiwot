package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"addis-hiwot/internal/domain/interfaces"
	"addis-hiwot/internal/domain/models"
	"addis-hiwot/internal/domain/schema"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(handler *UserHandler) *gin.Engine {
	r := gin.Default()
	r.POST("/register", handler.CreateUser)
	r.POST("/login", handler.LoginUser)
	r.GET("/users", handler.GetUsers)
	return r
}

func TestCreateUser_Success(t *testing.T) {
	mockUC := new(interfaces.MockUserUsecase)
	h := NewUserHandler(mockUC)
	r := setupRouter(h)

	input := schema.CreateUser{
		Email:    "test@example.com",
		Username: "testuser",
		Name:     "Test User",
		Password: "password123",
	}
	mockUC.On("Register", mock.AnythingOfType("*schema.CreateUser")).Return("token123", nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User registered successfully")
	assert.Contains(t, w.Body.String(), "token123")
	mockUC.AssertExpectations(t)
}

func TestCreateUser_BadRequest(t *testing.T) {
	mockUC := new(interfaces.MockUserUsecase)
	h := NewUserHandler(mockUC)
	r := setupRouter(h)

	// Invalid input (missing required fields)
	body := []byte(`{"email":"not-an-email"}`)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestCreateUser_Error(t *testing.T) {
	mockUC := new(interfaces.MockUserUsecase)
	h := NewUserHandler(mockUC)
	r := setupRouter(h)

	input := schema.CreateUser{
		Email:    "test@example.com",
		Username: "testuser",
		Name:     "Test User",
		Password: "password123",
	}
	mockUC.On("Register", mock.AnythingOfType("*schema.CreateUser")).Return("", errors.New("registration failed"))

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "registration failed")
	mockUC.AssertExpectations(t)
}

func TestLoginUser_Success(t *testing.T) {
	mockUC := new(interfaces.MockUserUsecase)
	h := NewUserHandler(mockUC)
	r := setupRouter(h)

	input := schema.LoginUser{
		Email:    "test@example.com",
		Password: "password123",
	}
	mockUC.On("Login", mock.AnythingOfType("*schema.LoginUser")).Return("token456", nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user logged in successfully")
	assert.Contains(t, w.Body.String(), "token456")
	mockUC.AssertExpectations(t)
}

func TestLoginUser_BadRequest(t *testing.T) {
	mockUC := new(interfaces.MockUserUsecase)
	h := NewUserHandler(mockUC)
	r := setupRouter(h)

	body := []byte(`{"email":"not-an-email"}`)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestLoginUser_Error(t *testing.T) {
	mockUC := new(interfaces.MockUserUsecase)
	h := NewUserHandler(mockUC)
	r := setupRouter(h)

	input := schema.LoginUser{
		Email:    "test@example.com",
		Password: "password123",
	}
	mockUC.On("Login", mock.AnythingOfType("*schema.LoginUser")).Return("", errors.New("login failed"))

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "login failed")
	mockUC.AssertExpectations(t)
}

func TestGetUsers_Success(t *testing.T) {
	mockUC := new(interfaces.MockUserUsecase)
	h := NewUserHandler(mockUC)
	r := setupRouter(h)

	mockUsers := []*models.UserResponse{
		{ID: 1, Email: "test@example.com", Username: "testuser", Name: "Test User"},
	}
	mockUC.On("GetAll").Return(mockUsers, nil)

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "testuser")
	mockUC.AssertExpectations(t)
}

func TestGetUsers_Error(t *testing.T) {
	mockUC := new(interfaces.MockUserUsecase)
	h := NewUserHandler(mockUC)
	r := setupRouter(h)

	mockUC.On("GetAll").Return(nil, errors.New("db error"))

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "db error")
	mockUC.AssertExpectations(t)
}
