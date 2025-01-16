package domain

import "time"

type VolunteerRequest struct {
	ID        int       `gorm:"primaryKey"`
	UserID    int       `gorm:"not null"`
	Type      string    `gorm:"not null"`
	Status    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (VolunteerRequest) TableName() string {
	return "requests"
}
