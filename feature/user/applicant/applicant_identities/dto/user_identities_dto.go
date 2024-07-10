package dto

import (
	"time"
)

type UserIdentity struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Number      string    `json:"number"`
	Type        string    `json:"type"`
	Status      int       `json:"status"`
	ExpiryDate  time.Time `json:"expiry_date"`
	PlaceIssued string    `json:"place_issued"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func main()
