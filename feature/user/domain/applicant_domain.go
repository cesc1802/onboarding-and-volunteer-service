package domain

import "time"

type ApplicantDomain struct {
	ID        int       `gorm:"primaryKey"`
	RoleID    *int      `gorm:"index"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Surname   string    `gorm:"not null"`
	Gender    string    `gorm:"size:6;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (ApplicantDomain) TableName() string {
	return "users"
}
