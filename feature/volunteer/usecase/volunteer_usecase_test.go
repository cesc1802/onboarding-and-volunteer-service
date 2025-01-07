package usecase

import (
	"errors"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockVolunteerRepository is a mock implementation of the VolunteerRepositoryInterface
type MockVolunteerRepository struct {
	mock.Mock
}

func (m *MockVolunteerRepository) CreateVolunteer(volunteer *domain.Volunteer) error {
	args := m.Called(volunteer)
	return args.Error(0)
}

func (m *MockVolunteerRepository) UpdateVolunteer(volunteer *domain.Volunteer) error {
	args := m.Called(volunteer)
	return args.Error(0)
}

func (m *MockVolunteerRepository) DeleteVolunteer(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockVolunteerRepository) FindVolunteerByID(id int) (*domain.Volunteer, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Volunteer), args.Error(1)
}

func TestCreateVolunteer(t *testing.T) {
	mockRepo := new(MockVolunteerRepository)
	usecase := NewVolunteerUsecase(mockRepo)

	input := dto.VolunteerCreateDTO{
		UserID:       1,
		DepartmentID: 2,
		Status:       0,
	}

	mockRepo.On("CreateVolunteer", mock.Anything).Return(nil)

	err := usecase.CreateVolunteer(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateVolunteer(t *testing.T) {
	mockRepo := new(MockVolunteerRepository)
	usecase := NewVolunteerUsecase(mockRepo)

	input := dto.VolunteerUpdateDTO{
		DepartmentID: 4,
		Status:       1,
	}
	volunteer := &domain.Volunteer{
		ID:           1,
		UserID:       1,
		DepartmentID: 2,
		Status:       0,
	}

	mockRepo.On("FindVolunteerByID", 1).Return(volunteer, nil)
	mockRepo.On("UpdateVolunteer", volunteer).Return(nil)

	err := usecase.UpdateVolunteer(1, input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteVolunteer(t *testing.T) {
	mockRepo := new(MockVolunteerRepository)
	usecase := NewVolunteerUsecase(mockRepo)

	mockRepo.On("DeleteVolunteer", 1).Return(nil)

	err := usecase.DeleteVolunteer(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindVolunteerByID(t *testing.T) {
	mockRepo := new(MockVolunteerRepository)
	usecase := NewVolunteerUsecase(mockRepo)

	volunteer := &domain.Volunteer{
		ID:           1,
		UserID:       1,
		DepartmentID: 2,
		Status:       0,
	}

	mockRepo.On("FindVolunteerByID", 1).Return(volunteer, nil)

	result, err := usecase.FindVolunteerByID(1)

	assert.NoError(t, err)
	assert.Equal(t, &dto.VolunteerResponseDTO{
		ID:           1,
		UserID:       1,
		DepartmentID: 2,
		Status:       0,
	}, result)
	mockRepo.AssertExpectations(t)
}

func TestFindVolunteerByID_NotFound(t *testing.T) {
	repo := new(MockVolunteerRepository)
	usecase := NewVolunteerUsecase(repo)

	id := 1

	// Return nil explicitly typed as *domain.Volunteer and an error
	expectedError := errors.New("volunteer not found")
	repo.On("FindVolunteerByID", id).Return((*domain.Volunteer)(nil), expectedError)

	// Call the usecase method
	response, err := usecase.FindVolunteerByID(id)

	// Verify the mock method was called with the correct arguments
	repo.AssertCalled(t, "FindVolunteerByID", id)

	// Assertions
	assert.Nil(t, response)                          // Ensure the response is nil
	assert.EqualError(t, err, expectedError.Error()) // Ensure the error message matches
}
