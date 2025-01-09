package storage

import (
	"errors"
	"log"
	"strings"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"gorm.io/gorm"
)

type AdminRepositoryInterface interface {
	GetListPendingRequest() ([]*domain.Request, string)
	GetPendingRequestByID(id int) (*domain.Request, string)
	GetListAllRequest() ([]*domain.Request, string)
	GetRequestByID(id int) (*domain.Request, string)
	ApproveRequest(id int, verifier_id int) string
	RejectRequest(id int, verifier_id int) string
	AddRejectNotes(id int, notes string) string
	DeleteRequest(id int) string
}

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) GetListPendingRequest() ([]*domain.Request, string) {
	var listRequest []*domain.Request
	result := r.db.Where("reject_notes IS NULL OR verifier_id IS NULL").Find(&listRequest)
	if result.Error != nil {
		return nil, result.Error.Error()
	}
	if len(listRequest) == 0 {
		return nil, "No request pending"
	}
	return listRequest, ""
}

func (r *AdminRepository) GetPendingRequestByID(id int) (*domain.Request, string) {
	log.Printf("Attempting to fetch request with ID: %d", id)
	var request domain.Request
	result := r.db.Where("id = ?", id).First(&request)
	if result.Error != nil {
		log.Printf("Error occurred while fetching request with ID %d: %v", id, result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, "No pending request found with the given ID"
		}
		return nil, result.Error.Error()
	}
	return &request, ""
}

func (r *AdminRepository) GetListAllRequest() ([]*domain.Request, string) {
	var listRequest []*domain.Request
	result := r.db.Find(&listRequest)
	if result.Error != nil {
		return nil, result.Error.Error()
	}
	if len(listRequest) == 0 {
		return nil, "No request found"
	}
	return listRequest, ""
}

func (r *AdminRepository) GetRequestByID(id int) (*domain.Request, string) {
	var request domain.Request
	result := r.db.Where("id = ?", id).First(&request)
	if result.Error != nil {

		return nil, result.Error.Error()
	}
	return &request, ""
}

func (r *AdminRepository) ApproveRequest(id int, verifierID int) string {
	// Get request type
	request, err := r.GetRequestByID(id)
	if err != "" {
		return err
	}
	if request.Status != 0 {
		return "Request already processed"
	}
	userID := request.UserID
	if strings.TrimSpace(request.Type) == "registration" {
		result := r.db.Model(&domain.Request{}).Where("id = ?", id).Update("status", 1).Update("verifier_id", verifierID)
		if result.Error != nil {
			return result.Error.Error()
		}
		// Change user role to 1 (applicant)
		s, done := updateRoleId(r, userID, 1)
		if done {
			return s
		}
		return "Approve request success"
	} else if strings.TrimSpace(request.Type) == "verification" {
		result := r.db.Model(&domain.Request{}).Where("id = ?", id).Update("status", 1).Update("verifier_id", verifierID)
		if result.Error != nil {
			return result.Error.Error()
		}
		// Change user role to 2 (volunteer)
		s, done := updateRoleId(r, userID, 2)
		if done {
			return s
		}
		// Insert into volunteer_details
		volunteerDetail := domain.VolunteerDetail{
			UserID: userID,
			Status: 1,
		}
		result = r.db.Create(&volunteerDetail)
		if result.Error != nil {
			return result.Error.Error()
		}
		return "Approve request success"
	}
	return "Invalid request type"
}

func (r *AdminRepository) RejectRequest(id int, verifierID int) string {
	result := r.db.Model(&domain.Request{}).Where("id = ?", id).Update("status", 2).Update("verifier_id", verifierID)
	if result.Error != nil {
		return result.Error.Error()
	}
	return "Reject request success"
}

func (r *AdminRepository) AddRejectNotes(id int, notes string) string {
	result := r.db.Model(&domain.Request{}).Where("id = ?", id).Update("reject_notes", notes)
	if result.Error != nil {
		return result.Error.Error()
	}
	return "Add reject notes success"
}

func (r *AdminRepository) DeleteRequest(id int) string {
	result := r.db.Where("id = ?", id).Delete(&domain.Request{})
	if result.Error != nil {
		return result.Error.Error()
	}
	return "Delete request success"
}

func (r *AdminRepository) GetRequestByRequestID(requestID int) (*domain.Request, string) {
	var request domain.Request
	r.db.First(&request, requestID)
	if request.ID == 0 {
		return nil, "Request not found"
	}
	return &request, ""
}

func updateRoleId(r *AdminRepository, userID uint, roleId int) (string, bool) {
	result := r.db.Model(&domain.User{}).Where("id = ?", userID).Update("role_id", roleId)
	if result.Error != nil {
		return result.Error.Error(), true
	}
	return "", false
}
