// department_usecase_test.go
package usecase

import (
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDepartmentRepository is a mock implementation of the DepartmentRepositoryInterface
type MockDepartmentRepository struct {
	mock.Mock
}

// Create is a mock method for creating a department
func (m *MockDepartmentRepository) Create(department *domain.Department) error {
	args := m.Called(department)
	return args.Error(0)
}

// GetByID is a mock method for getting a department by ID.
func (m *MockDepartmentRepository) GetByID(id uint) (*domain.Department, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Department), args.Error(1)
}

// Update is a mock method for updating a department
func (m *MockDepartmentRepository) Update(department *domain.Department) error {
	args := m.Called(department)
	return args.Error(0)
}

// Delete is a mock method for deleting a department
func (m *MockDepartmentRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// TestCreateDepartment tests the CreateDepartment method in the use case.
func TestCreateDepartment(t *testing.T) {
	mockRepo := new(MockDepartmentRepository) // Create a new mock repository.
	usecase := NewDepartmentUsecase(mockRepo) // Create a new use case with the mock repository.

	input := dto.DepartmentCreateDTO{
		Name:    "HR",
		Address: "123 HR Street",
		Status:  123,
	}

	expectedDepartment := &domain.Department{
		Name:    input.Name,
		Address: input.Address,
		Status:  input.Status,
	}

	mockRepo.On("Create", expectedDepartment).Return(nil) // Set expectation on the mock.

	err := usecase.CreateDepartment(input) // Call the use case method.
	assert.NoError(t, err)                 // Assert no error.
	mockRepo.AssertExpectations(t)         // Verify the mock's expectations.
}

// TestGetDepartmentByID tests the GetDepartmentByID method in the use case.
func TestGetDepartmentByID(t *testing.T) {
	mockRepo := new(MockDepartmentRepository)
	usecase := NewDepartmentUsecase(mockRepo)

	expectedDepartment := &domain.Department{

		Name:    "Finance",
		Address: "456 Finance Street",
		Status:  456,
	}

	mockRepo.On("GetByID", uint(1)).Return(expectedDepartment, nil)

	department, err := usecase.GetDepartmentByID(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedDepartment, department)

	mockRepo.AssertExpectations(t)
}

// TestUpdateDepartment tests the UpdateDepartment method in the use case.
func TestUpdateDepartment(t *testing.T) {
	mockRepo := new(MockDepartmentRepository)
	usecase := NewDepartmentUsecase(mockRepo)

	input := dto.DepartmentUpdateDTO{
		Name:    "IT Updated",
		Address: "789 IT Street Updated",
		Status:  789,
	}

	existingDepartment := &domain.Department{
		Name:    "IT",
		Address: "789 IT Street",
		Status:  789,
	}

	updatedDepartment := &domain.Department{
		Name:    input.Name,
		Address: input.Address,
		Status:  input.Status,
	}

	mockRepo.On("GetByID", uint(1)).Return(existingDepartment, nil)
	mockRepo.On("Update", updatedDepartment).Return(nil)

	err := usecase.UpdateDepartment(1, input)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

// TestDeleteDepartment tests the DeleteDepartment method in the use case.
func TestDeleteDepartment(t *testing.T) {
	mockRepo := new(MockDepartmentRepository)
	usecase := NewDepartmentUsecase(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := usecase.DeleteDepartment(1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}
