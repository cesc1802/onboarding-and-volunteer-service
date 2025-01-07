package domain

import (
	"time"
)

type Volunteer struct {
	ID           int       `gorm:"primaryKey"`
	UserID       int       `gorm:"unique;notnull"`
	DepartmentID int       `gorm:"notnull"`
	Status       int       `gorm:"notnull"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (Volunteer) TableName() string {
	return "volunteer_details"
}
