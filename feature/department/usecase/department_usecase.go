package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/storage"
)

// DepartmentUsecase defines the methods that any use case implementation must provide.
type DepartmentUsecaseInterface interface {
	CreateDepartment(input dto.DepartmentCreateDTO) error
	GetDepartmentByID(id uint) (*domain.Department, error)
	UpdateDepartment(id uint, input dto.DepartmentUpdateDTO) error
	DeleteDepartment(id uint) error
}

// DepartmentUsecase handles the business logic for departments.
type DepartmentUsecase struct {
	repo storage.DepartmentRepositoryInterface
}

// NewDepartmentUsecase creates a new instance of DepartmentUsecase.
func NewDepartmentUsecase(repo storage.DepartmentRepositoryInterface) *DepartmentUsecase {
	return &DepartmentUsecase{repo: repo}
}

// CreateDepartment creates a new department using the provided DTO.
func (u *DepartmentUsecase) CreateDepartment(input dto.DepartmentCreateDTO) error {
	department := &domain.Department{
		Name:    input.Name,
		Address: input.Address,
		Status:  input.Status,
	}
	return u.repo.Create(department)
}

// GetDepartmentByID retrieves a department by its ID.
func (u *DepartmentUsecase) GetDepartmentByID(id uint) (*domain.Department, error) {
	return u.repo.GetByID(id)
}

// UpdateDepartment updates a department using the provided DTO.
func (u *DepartmentUsecase) UpdateDepartment(id uint, input dto.DepartmentUpdateDTO) error {
	department, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}
	department.Name = input.Name
	department.Address = input.Address
	department.Status = input.Status
	return u.repo.Update(department)

}

// DeleteDepartment deletes a department by its ID.
func (u *DepartmentUsecase) DeleteDepartment(id uint) error {
	return u.repo.Delete(id)
}
