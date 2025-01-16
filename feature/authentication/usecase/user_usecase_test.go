package usecase

import (
	"testing"
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/dto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthenticationStore is a mock implementation of the AuthenticationStore interface
type MockAuthenticationStore struct {
	mock.Mock
}

func (m *MockAuthenticationStore) GetUserByEmail(email, password string) (*domain.User, string) {
	args := m.Called(email, password)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.String(1)
	}
	return nil, args.String(1)
}

func (m *MockAuthenticationStore) RegisterUser(req *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error) {
	args := m.Called(req)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.RegisterUserResponse), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestUserUsecase_Login(t *testing.T) {
	mockRepo := new(MockAuthenticationStore)
	secretKey := "secret"
	usecase := NewUserUsecase(mockRepo, secretKey)

	req := dto.LoginUserRequest{
		Email:    "test@example.com",
		Password: "password",
	}
	roleId := 456
	mockUser := domain.User{
		ID:     123,
		RoleID: &roleId,
	}
	// Use a pointer to mockUser
	mockRepo.On("GetUserByEmail", req.Email, req.Password).Return(&mockUser, "")

	resp, msg := usecase.Login(req)

	assert.Equal(t, "", msg)
	assert.NotNil(t, resp)
	assert.NotEmpty(t, resp.Token)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(resp.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	assert.NoError(t, err)

	assert.Equal(t, float64(mockUser.ID), claims["userId"])
	if mockUser.RoleID != nil {
		assert.Equal(t, float64(*mockUser.RoleID), claims["roleId"])
	} else {
		t.Error("RoleID should not be nil")
	}
	// Validate the 'exp' claim
	expiration, ok := claims["exp"].(float64)
	assert.True(t, ok, "Token should have a valid 'exp' claim")
	assert.GreaterOrEqual(t, int64(expiration), time.Now().Unix(), "Expiration time should be in the future")

}

func TestUserUsecase_RegisterUser(t *testing.T) {
	mockRepo := new(MockAuthenticationStore)
	secretKey := "secret"
	usecase := NewUserUsecase(mockRepo, secretKey)

	req := dto.RegisterUserRequest{
		Email:    "test@example.com",
		Password: "password",
	}
	mockRepo.On("GetUserByEmail", req.Email, "").Return(nil, "")
	mockResponse := &dto.RegisterUserResponse{
		Message: req.Email,
	}
	mockRepo.On("RegisterUser", &req).Return(mockResponse, nil)

	resp, msg := usecase.RegisterUser(req)

	assert.Equal(t, "", msg)
	assert.NotNil(t, resp)
	assert.Equal(t, mockResponse, resp)
}
