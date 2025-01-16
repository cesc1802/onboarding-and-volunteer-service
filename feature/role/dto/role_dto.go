package dto

// RoleCreateDTO represents the data transfer object for creating a role.
type RoleCreateDTO struct {
	Name string `json:"name" binding:"required"`
}

// RoleUpdateDTO represents the data transfer object for updating a role.
type RoleUpdateDTO struct {
	Name string `json:"name" binding:"required"`
}

// RoleResponseDTO represents the data transfer object for returning role data.
type RoleResponseDTO struct {
	Name string `json:"name" binding:"required"`
}
