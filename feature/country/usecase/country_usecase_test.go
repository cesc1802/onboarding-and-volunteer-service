package usecase

import (
	"errors"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockCountryRepository is a mock implementation of CountryRepositoryInterface.
type MockCountryRepository struct {
	mock.Mock
}

// Create is a mock method for creating a country.
func (m *MockCountryRepository) Create(country *domain.Country) error {
	args := m.Called(country)
	return args.Error(0)
}

// GetCountryByID is a mock method for getting a country by ID.
func (m *MockCountryRepository) GetByID(id uint) (*domain.Country, error) {
	args := m.Called(id)
	country, ok := args.Get(0).(*domain.Country)
	if !ok {
		return nil, args.Error(1)
	}
	return country, args.Error(1)
}

// Update is a mock method for updating a country.
func (m *MockCountryRepository) Update(country *domain.Country) error {
	args := m.Called(country)
	return args.Error(0)
}

// Delete is a mock method for deleting a country.
func (m *MockCountryRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateCountry(t *testing.T) {
	mockRepo := new(MockCountryRepository)
	mockRepo.ExpectedCalls = nil
	usecase := NewCountryUsecase(mockRepo)

	input := dto.CountryCreateDTO{
		Name:   "TestCountry",
		Status: 1,
	}

	expectedCountry := &domain.Country{
		Name:   input.Name,
		Status: input.Status,
	}

	mockRepo.On("Create", mock.MatchedBy(func(country *domain.Country) bool {
		return country.Name == expectedCountry.Name && country.Status == expectedCountry.Status
	})).Return(nil)

	err := usecase.CreateCountry(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetCountryByID(t *testing.T) {
	mockRepo := new(MockCountryRepository)
	usecase := NewCountryUsecase(mockRepo)
	expectedCountryDomain := &domain.Country{
		Name:   "TestCountry",
		Status: 1,
	}

	expectedCountryDTO := &dto.CountryCreateDTO{
		Name:   "TestCountry",
		Status: 1,
	}

	mockRepo.On("GetByID", uint(1)).Return(expectedCountryDomain, nil)

	country, err := usecase.GetCountryByID(1)

	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedCountryDTO.Name, country.Name, "Expected country name to match")
	assert.Equal(t, expectedCountryDTO.Status, country.Status, "Expected country status to match")

	mockRepo.AssertExpectations(t)
}
func TestUpdateCountry(t *testing.T) {
	mockRepo := new(MockCountryRepository)
	mockRepo.ExpectedCalls = nil
	usecase := NewCountryUsecase(mockRepo)

	input := dto.CountryUpdateDTO{
		Name:   "UpdatedCountry",
		Status: 2,
	}

	existingCountry := &domain.Country{
		ID:     1,
		Name:   "TestCountry",
		Status: 1,
	}

	updatedCountry := &domain.Country{
		ID:     1,
		Name:   input.Name,
		Status: input.Status,
	}

	mockRepo.On("GetByID", uint(1)).Return(existingCountry, nil)
	mockRepo.On("Update", updatedCountry).Return(nil)

	err := usecase.UpdateCountry(1, input)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCountry(t *testing.T) {
	mockRepo := new(MockCountryRepository)
	mockRepo.ExpectedCalls = nil
	usecase := NewCountryUsecase(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := usecase.DeleteCountry(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetCountryByID_NotFound(t *testing.T) {
	mockRepo := new(MockCountryRepository)
	mockRepo.ExpectedCalls = nil
	usecase := NewCountryUsecase(mockRepo)

	mockRepo.On("GetByID", uint(1)).Return(nil, errors.New("not found"))

	country, err := usecase.GetCountryByID(1)

	assert.Error(t, err)
	assert.Nil(t, country)
	mockRepo.AssertExpectations(t)
}
