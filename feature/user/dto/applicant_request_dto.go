package dto

type ApplicationFormDTO struct {
	UserID       uint   `json:"user_id" binding:"required"`
	DepartmentID uint   `json:"department_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Surname      string `json:"surname" binding:"required"`
	Sex          string `json:"sex" binding:"required"`
	DateOfBirth  string `json:"dob" binding:"required"`
}
