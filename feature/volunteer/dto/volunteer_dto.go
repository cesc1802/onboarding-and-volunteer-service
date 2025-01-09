package dto

type VolunteerCreateDTO struct {
	UserID             int    `json:"user_id" binding:"required"`
	DepartmentID       int    `json:"department_id" binding:"required"`
	Status             int    `json:"status" binding:"required"`
	Dob                string `json:"dob"`
	Mobile             string `json:"mobile"`
	CountryID          int    `json:"country_id"`
	ResidentCountryID  int    `json:"resident_country_id"`
	Avatar             string `json:"avatar"`
	VerificationStatus int    `json:"verification_status"`
}

type VolunteerUpdateDTO struct {
	DepartmentID       int    `json:"department_id"`
	Dob                string `json:"dob"`
	Mobile             string `json:"mobile"`
	CountryID          int    `json:"country_id"`
	ResidentCountryID  int    `json:"resident_country_id"`
	Avatar             string `json:"avatar"`
	VerificationStatus int    `json:"verification_status"`
	Status             int    `json:"status"`
}

type VolunteerResponseDTO struct {
	ID                 int    `json:"id"`
	UserID             int    `json:"user_id"`
	DepartmentID       int    `json:"department_id"`
	Dob                string `json:"dob"`
	Mobile             string `json:"mobile"`
	CountryID          int    `json:"country_id"`
	ResidentCountryID  int    `json:"resident_country_id"`
	Avatar             string `json:"avatar"`
	VerificationStatus int    `json:"verification_status"`
	Status             int    `json:"status"`
}
