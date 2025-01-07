package usecase

import (
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRoleRepository is a mock implementation of the RoleRepositoryInterface.
type MockRoleRepository struct {
	mock.Mock
}

// Create is a mock method for creating a role
func (m *MockRoleRepository) Create(role *domain.Role) error {
	args := m.Called(role)
	return args.Error(0)
}

// GetByID is a mock method for getting a role by ID.
func (m *MockRoleRepository) GetByID(id uint) (*domain.Role, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Role), args.Error(1)
}

// Update is a mock method for updating a role
func (m *MockRoleRepository) Update(role *domain.Role) error {
	args := m.Called(role)
	return args.Error(0)
}

// Delete is a mock method for deleting a role
func (m *MockRoleRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateRole(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	usecase := NewRoleUsecase(mockRepo)

	input := dto.RoleCreateDTO{
		Name:   "Admin",
		Status: "123",
	}

	role := &domain.Role{
		Name:   input.Name,
		Status: input.Status,
	}

	mockRepo.On("Create", role).Return(nil)

	err := usecase.CreateRole(input)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Create", role)
}

func TestGetRoleByID(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	usecase := NewRoleUsecase(mockRepo)

	role := &domain.Role{

		Name:   "Admin",
		Status: "456",
	}

	mockRepo.On("GetByID", uint(1)).Return(role, nil)

	result, err := usecase.GetRoleByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, role.Name, result.Name)
	assert.Equal(t, role.Status, result.Status)
	mockRepo.AssertCalled(t, "GetByID", uint(1))
}

func TestUpdateRole(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	usecase := NewRoleUsecase(mockRepo)

	role := &domain.Role{
		Name:   "Admin",
		Status: "789",
	}

	input := dto.RoleUpdateDTO{
		Name:   "Admin Updated",
		Status: "666",
	}

	mockRepo.On("GetByID", uint(1)).Return(role, nil)
	mockRepo.On("Update", role).Return(nil)

	err := usecase.UpdateRole(1, input)
	assert.NoError(t, err)
	assert.Equal(t, input.Name, role.Name)
	assert.Equal(t, input.Status, role.Status)
	mockRepo.AssertCalled(t, "GetByID", uint(1))
	mockRepo.AssertCalled(t, "Update", role)
}

func TestDeleteRole(t *testing.T) {
	mockRepo := new(MockRoleRepository)
	usecase := NewRoleUsecase(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := usecase.DeleteRole(1)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Delete", uint(1))
}
