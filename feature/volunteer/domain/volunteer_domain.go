package domain

import (
	"time"
)

type Volunteer struct {
	ID                 int       `gorm:"primaryKey"`
	UserID             int       `gorm:"unique;notnull"`
	DepartmentID       *int      `gorm:"notnull"`
	Dob                time.Time `gorm:"not null"`
	Mobile             string    `gorm:"not null"`
	CountryID          *int      `gorm:"index"`
	ResidentCountryID  *int      `gorm:"index"`
	Avatar             string    `gorm:"default:null"`
	VerificationStatus int       `gorm:"default:0"`
	Status             int       `gorm:"not null"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

func (Volunteer) TableName() string {
	return "volunteer_details"
}
