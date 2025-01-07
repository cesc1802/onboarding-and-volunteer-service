package dto

// RoleCreateDTO represents the data transfer object for creating a role.
type RoleCreateDTO struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required,oneof=pending approved rejected "`
}

// RoleUpdateDTO represents the data transfer object for updating a role.
type RoleUpdateDTO struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required,oneof=pending approved rejected "`
}

// RoleResponseDTO represents the data transfer object for returning role data.
type RoleResponseDTO struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
}
