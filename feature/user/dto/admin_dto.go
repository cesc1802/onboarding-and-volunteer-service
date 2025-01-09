package dto

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
)

type PendingRequest struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Type   string `json:"type"`
}

type RequestResponse struct {
	ID          int    `json:"id"`
	UserID      uint   `json:"user_id"`
	Type        string `json:"type"`
	Status      int    `json:"status"`
	RejectNotes string `json:"reject_notes"`
	VerifierID  int    `json:"verifier_id"`
}

type ListRequest struct {
	Requests []*domain.Request `json:"requests"`
}

type AddRejectNoteRequest struct {
	Notes string `json:"notes"`
}
