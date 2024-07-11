package domain

import (
	"time"
)

// lớp struct tương tác với database (GORM)
type Request struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	Type        string `gorm:"size:45;not null"`
	Status      uint   `gorm:"not null"`
	RejectNotes *string
	VerifierID  *uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
