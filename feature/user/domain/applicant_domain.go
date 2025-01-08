package domain

import "time"

type ApplicantDomain struct {
	ID                 int       `gorm:"primaryKey"`
	RoleID             int       `gorm:"not null"`
	DepartmentID       int       `gorm:"not null"`
	Email              string    `gorm:"not null"`
	Password           string    `gorm:"not null"`
	Name               string    `gorm:"not null"`
	Surname            string    `gorm:"not null"`
	Gender             string    `gorm:"not null"`
	DOB                time.Time `gorm:"not null"`
	Mobile             string    `gorm:"not null"`
	CountryID          int       `gorm:"not null"`
	ResidentCountryID  int       `gorm:"not null"`
	Avatar             string    `gorm:"default:null"`
	VerificationStatus int       `gorm:"default:0"`
	Status             int       `gorm:"not null"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

func (ApplicantDomain) TableName() string {
	return "users"
}
