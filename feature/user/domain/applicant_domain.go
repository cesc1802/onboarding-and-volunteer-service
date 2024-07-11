package domain

import (
	"time"
)

// lớp struct tương tác với database (GORM)

type User struct {
	ID                 uint `gorm:"primaryKey"`
	RoleID             uint `gorm:"not null"`
	DepartmentID       *uint
	Email              string    `gorm:"size:45;not null"`
	Password           string    `gorm:"type:text;not null"`
	Name               string    `gorm:"size:45;not null"`
	Surname            string    `gorm:"size:45;not null"`
	Gender             string    `gorm:"size:20;not null"`
	DOB                time.Time `gorm:"not null"`
	Mobile             string    `gorm:"size:15;not null"`
	CountryID          uint      `gorm:"not null"`
	ResidentCountryID  uint      `gorm:"not null"`
	Avatar             *string
	VerificationStatus uint `gorm:"default:0"`
	Status             uint `gorm:"not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
