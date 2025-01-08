package dto

import "time"

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginUserResponse struct {
	ID                 int       `json:"id"`
	RoleID             int       `json:"role_id"`
	DepartmentID       int       `json:"department_id"`
	Email              string    `json:"email"`
	Name               string    `json:"name"`
	Surname            string    `json:"surname"`
	Gender             string    `json:"gender"`
	DOB                time.Time `json:"dob"`
	Mobile             string    `json:"mobile"`
	CountryID          int       `json:"country_id"`
	ResidentCountryID  int       `json:"resident_country_id"`
	Avatar             string    `json:"avatar"`
	VerificationStatus int       `json:"verification_status"`
	Status             int       `json:"status"`
}

type LoginUserTokenResponse struct {
	Token string `json:"token"`
}

type RegisterUserRequest struct {
	ID                 int    `json:"id"`
	Email              string `json:"email" binding:"required,email"`
	Name               string `json:"name" binding:"required"`
	Surname            string `json:"surname" binding:"required"`
	Password           string `json:"password,omitempty" binding:"required"`
	RePassword         string `json:"re_password,omitempty" binding:"required"`
	Gender             string `json:"gender" binding:"required,oneof=male female"`
	Dob                string `json:"dob"`
	Mobile             string `json:"mobile"`
	CountryID          int    `json:"country_id"`
	ResidentCountryID  int    `json:"resident_country_id"`
	Avatar             string `json:"avatar"`
	VerificationStatus int    `json:"verification_status"`
	Status             int    `json:"status" binding:"required"`
	RoleID             int    `json:"role_id" binding:"required"`
	DepartmentID       int    `json:"department_id" binding:"required"`
}
type RegisterUserResponse struct {
	Message string               `json:"message"`
	User    *RegisterUserRequest `json:"user,omitempty"`
}
