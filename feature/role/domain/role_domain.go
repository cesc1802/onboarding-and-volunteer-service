package domain

import (
	"time"
)

// Role struct represents the role entity interacting with the database using GORM.
type Role struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Status    string    `gorm:"type:ENUM('pending', 'approved', 'rejected');not null" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
