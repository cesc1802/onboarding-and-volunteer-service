package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserIdentityRepository struct {
	mock.Mock
}

func (m *MockUserIdentityRepository) CreateUserIdentity(UserIdentity *domain.UserIdentity) error {
	args := m.Called(UserIdentity)
	return args.Error(0)
}

func (m *MockUserIdentityRepository) UpdateUserIdentity(UserIdentity *domain.UserIdentity) error {
	args := m.Called(UserIdentity)
	return args.Error(0)
}

func (m *MockUserIdentityRepository) FindUserIdentityByID(id int) (*domain.UserIdentity, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.UserIdentity), args.Error(1)
}

func TestCreateUserIdentity(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)
	usecase := NewUserIdentityUsecase(mockRepo)

	input := dto.CreateUserIdentityRequest{
		UserID:      2,
		Number:      "123456789",
		Type:        "Citizen ID",
		Status:      0,
		ExpiryDate:  "12-12-2025",
		PlaceIssued: "Some city",
	}

	mockRepo.On("CreateUserIdentity", mock.Anything).Return(nil)

	err := usecase.CreateUserIdentity(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserIdentity(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)
	usecase := NewUserIdentityUsecase(mockRepo)

	input := dto.UpdateUserIdentityRequest{
		UserID:      2,
		Number:      "123555789",
		Type:        "Passport",
		Status:      0,
		ExpiryDate:  "12-12-2028",
		PlaceIssued: "Some city",
	}
	userIdentity := &domain.UserIdentity{
		ID:          1,
		UserID:      2,
		Number:      "123456987",
		Type:        "Citizen ID",
		Status:      0,
		ExpiryDate:  time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC),
		PlaceIssued: "Some city",
	}

	mockRepo.On("FindUserIdentityByID", 1).Return(userIdentity, nil)
	mockRepo.On("UpdateUserIdentity", userIdentity).Return(nil)

	err := usecase.UpdateUserIdentity(1, input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindUserIdentityByID(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)
	usecase := NewUserIdentityUsecase(mockRepo)

	userIdentity := &domain.UserIdentity{
		ID:          1,
		UserID:      2,
		Number:      "123456987",
		Type:        "Citizen ID",
		Status:      0,
		ExpiryDate:  time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC),
		PlaceIssued: "Some city",
	}

	mockRepo.On("FindUserIdentityByID", 1).Return(userIdentity, nil)

	result, err := usecase.FindUserIdentityByID(1)

	assert.NoError(t, err)
	assert.Equal(t, &dto.UserIdentityResponse{}, result)
	mockRepo.AssertExpectations(t)
}

func TestFindUserIdentityByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)
	usecase := NewUserIdentityUsecase(mockRepo)

	mockRepo.On("FindUserIdentityByID", 1).Return(nil, errors.New("record not found"))

	result, err := usecase.FindUserIdentityByID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
