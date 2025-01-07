package dto

// DepartmentCreateDTO represents the data transfer object for creating a department.
type DepartmentCreateDTO struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"location" binding:"required"`
	Status  uint   `json:"status" binding:"required"`
}

// DepartmentUpdateDTO represents the data transfer object for updating a department.
type DepartmentUpdateDTO struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"location" binding:"required"`
	Status  uint   `json:"status" binding:"required"`
}

type DepartmentResponseDTO struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"location" binding:"required"`
	Status  uint   `json:"status" binding:"required"`
}
