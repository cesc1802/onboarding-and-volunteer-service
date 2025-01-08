package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/storage"
)

// RoleRepository defines the methods that any repository implementation must provide.
type RoleUsecaseInterface interface {
	CreateRole(role dto.RoleCreateDTO) error
	GetRoleByID(id uint) (*domain.Role, error)
	UpdateRole(id uint, input dto.RoleUpdateDTO) error
	DeleteRole(id uint) error
}

// RoleUsecase handles the business logic for roles.
type RoleUsecase struct {
	Rolerepo storage.RoleRepositoryInterface
}

// NewRoleUsecase creates a new instance of RoleUsecase.
func NewRoleUsecase(Rolerepo storage.RoleRepositoryInterface) *RoleUsecase {
	return &RoleUsecase{Rolerepo: Rolerepo}
}

// CreateRole creates a new role using the provided DTO.
func (u *RoleUsecase) CreateRole(input dto.RoleCreateDTO) error {
	role := &domain.Role{
		Name: input.Name,
	}
	return u.Rolerepo.Create(role)
}

// GetRoleByID retrieves a role by its ID.
func (u *RoleUsecase) GetRoleByID(id uint) (*domain.Role, error) {
	return u.Rolerepo.GetByID(id)
}

// UpdateRole updates a role using the provided DTO.
func (u *RoleUsecase) UpdateRole(id uint, input dto.RoleUpdateDTO) error {
	role, err := u.Rolerepo.GetByID(id)
	if err != nil {
		return err
	}
	role.Name = input.Name
	return u.Rolerepo.Update(role)
}

// DeleteRole deletes a role by its ID.
func (u *RoleUsecase) DeleteRole(id uint) error {
	return u.Rolerepo.Delete(id)
}
