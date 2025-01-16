package storage

import (
	"errors"
	"testing"
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/domain"
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
	return args.Get(0).(*domain.UserIdentity), args.Error(1)
}

func TestCreateUserIdentity(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)

	identity := &domain.UserIdentity{
		ID:          1,
		UserID:      2,
		Number:      "123456789",
		Type:        "Citizen ID",
		Status:      0,
		ExpiryDate:  time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC),
		PlaceIssued: "Some city",
	}

	mockRepo.On("CreateUserIdentity", identity).Return(nil).Once()

	err := mockRepo.CreateUserIdentity(identity)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUserIdentity(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)

	identity := &domain.UserIdentity{
		ID:          1,
		UserID:      2,
		Number:      "010293984",
		Type:        "Citizen ID",
		Status:      0,
		ExpiryDate:  time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC),
		PlaceIssued: "Some city",
	}

	mockRepo.On("UpdateUserIdentity", identity).Return(nil).Once()

	err := mockRepo.UpdateUserIdentity(identity)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindUserIdentityByID(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)

	identity := &domain.UserIdentity{
		ID:          1,
		UserID:      2,
		Number:      "123456789",
		Type:        "Citizen ID",
		Status:      0,
		ExpiryDate:  time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC),
		PlaceIssued: "Some city",
	}

	mockRepo.On("FindUserIdentityByID", 1).Return(identity, nil).Once()

	result, err := mockRepo.FindUserIdentityByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, identity.ID, result.ID)
	mockRepo.AssertExpectations(t)
}

func TestFindUserIdentityByID_NotFound(t *testing.T) {
	mockRepo := new(MockUserIdentityRepository)

	mockRepo.On("FindUserIdentityByID", 1).Return((*domain.UserIdentity)(nil), errors.New("user identity not found")).Once()

	result, err := mockRepo.FindUserIdentityByID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "user identity not found", err.Error())

	mockRepo.AssertExpectations(t)
}
