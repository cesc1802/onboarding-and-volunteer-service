package domain

import (
	"time"
)

// Department struct that interacts with databases (GORM)
type Department struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Address   string    `json:"location"`
	Status    uint      `gorm:"not null" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
