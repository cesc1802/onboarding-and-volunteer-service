package storage

import (
	"testing"
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
	gorm.DB
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func TestCreateUserIdentity(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewUserIdentityRepository(&mockDB.DB)

	userIdentity := &domain.UserIdentity{
		ID:          1,
		UserID:      2,
		Number:      "123456789",
		Type:        "Citizen ID",
		Status:      0,
		ExpiryDate:  time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC),
		PlaceIssued: "Some city",
	}
	mockDB.On("Create", userIdentity).Return(&gorm.DB{Error: nil})

	err := repo.CreateUserIdentity(userIdentity)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestUpdateUserIdentity(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewUserIdentityRepository(&mockDB.DB)

	userIdentity := &domain.UserIdentity{
		ID:          1,
		UserID:      2,
		Number:      "123456987",
		Type:        "Citizen ID",
		Status:      0,
		ExpiryDate:  time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC),
		PlaceIssued: "Some city",
	}
	mockDB.On("Save", userIdentity).Return(&gorm.DB{Error: nil})

	err := repo.UpdateUserIdentity(userIdentity)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestFindUserIdentityByID(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewUserIdentityRepository(&mockDB.DB)

	userIdentity := &domain.UserIdentity{
		ID:          1,
		UserID:      2,
		Number:      "123456987",
		Type:        "Citizen ID",
		Status:      0,
		ExpiryDate:  time.Date(2025, 12, 12, 0, 0, 0, 0, time.UTC),
		PlaceIssued: "Some city",
	}
	mockDB.On("First", &userIdentity, 1).Return(&gorm.DB{Error: nil}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(**domain.UserIdentity)
		*arg = userIdentity
	})

	result, err := repo.FindUserIdentityByID(1)

	assert.NoError(t, err)
	assert.Equal(t, userIdentity, result)
	mockDB.AssertExpectations(t)
}
