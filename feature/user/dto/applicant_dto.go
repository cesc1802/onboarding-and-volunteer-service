package dto

type UserSignupDTO struct {
	Email   string `json:"email" binding:"required,email"`
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}

type UserUpdateDTO struct {
	ID      uint   `json:"id" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}
