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

func (m *MockUserIdentityRepository) CreateUserIdentity(identity *domain.UserIdentity) error {
	args := m.Called(identity)
	return args.Error(0)
}

func (m *MockUserIdentityRepository) UpdateUserIdentity(identity *domain.UserIdentity) error {
	args := m.Called(identity)
	return args.Error(0)
}

func (m *MockUserIdentityRepository) FindUserIdentityByID(id int) (*domain.UserIdentity, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.UserIdentity), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestCreateUserIdentity(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)
	usecase := NewUserIdentityUsecase(mockRepo)

	req := dto.CreateUserIdentityRequest{
		UserID:      1,
		Number:      "0987123456",
		Type:        " Citizen ID",
		Status:      01,
		ExpiryDate:  "2025-01-01",
		PlaceIssued: "City X",
	}

	expiryDate, _ := time.Parse("2006-01-02", req.ExpiryDate)
	identity := &domain.UserIdentity{
		UserID:      req.UserID,
		Number:      req.Number,
		Type:        req.Type,
		Status:      req.Status,
		ExpiryDate:  expiryDate,
		PlaceIssued: req.PlaceIssued,
	}

	mockRepo.On("CreateUserIdentity", identity).Return(nil)

	err := usecase.CreateUserIdentity(req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserIdentity(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)
	usecase := NewUserIdentityUsecase(mockRepo)

	req := dto.UpdateUserIdentityRequest{
		UserID:      1,
		Number:      "654321",
		Type:        "ID",
		Status:      01,
		ExpiryDate:  "2026-01-01",
		PlaceIssued: "City Y",
	}

	id := 1
	expiryDate, _ := time.Parse("2006-01-02", req.ExpiryDate)
	identity := &domain.UserIdentity{
		ID:          id,
		UserID:      req.UserID,
		Number:      req.Number,
		Type:        req.Type,
		Status:      req.Status,
		ExpiryDate:  expiryDate,
		PlaceIssued: req.PlaceIssued,
	}

	mockRepo.On("UpdateUserIdentity", identity).Return(nil)

	err := usecase.UpdateUserIdentity(id, req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindUserIdentityByID(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)
	usecase := NewUserIdentityUsecase(mockRepo)

	id := 1
	identity := &domain.UserIdentity{
		ID:          id,
		UserID:      2,
		Number:      "789123",
		Type:        "Passport",
		Status:      01,
		ExpiryDate:  time.Date(2027, 1, 1, 0, 0, 0, 0, time.UTC),
		PlaceIssued: "City Z",
	}

	response := &dto.UserIdentityResponse{
		ID:          identity.ID,
		UserID:      identity.UserID,
		Number:      identity.Number,
		Type:        identity.Type,
		Status:      identity.Status,
		ExpiryDate:  "2027-01-01",
		PlaceIssued: identity.PlaceIssued,
	}

	mockRepo.On("FindUserIdentityByID", id).Return(identity, nil)

	result, err := usecase.FindUserIdentityByID(id)

	assert.NoError(t, err)
	assert.Equal(t, response, result)
	mockRepo.AssertExpectations(t)
}

func TestFindUserIdentityByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)
	usecase := NewUserIdentityUsecase(mockRepo)

	id := 999
	mockRepo.On("FindUserIdentityByID", id).Return(nil, errors.New("record not found"))

	result, err := usecase.FindUserIdentityByID(id)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
