package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"gorm.io/gorm"
	"strings"
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
	result := r.db.Where("status = ?", 0).Find(&listRequest)
	if result.Error != nil {
		return nil, result.Error.Error()
	}
	if len(listRequest) == 0 {
		return nil, "No request pending"
	}
	return listRequest, ""
}

func (r *AdminRepository) GetPendingRequestByID(id int) (*domain.Request, string) {
	var request domain.Request
	result := r.db.Where("id = ? and status = 0", id).First(&request)
	if result.Error != nil {
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

// ApproveRequest change status of request to 1 (approved)
// change verifier_id to admin id
// if requestType is registration, change user role to 1 (applicant)
// else if requestType is verification, change user role to 2 (volunteer) and change verification status to 1 (active)
// and insert this user to volunteer_details table
func (r *AdminRepository) ApproveRequest(id int, verifier_id int) string {
	// get request type
	request := r.getRequestByRequestID(id)
	if request == nil {
		return "Request not found"
	}
	if request.Status != 0 {
		return "Request already processed"
	}
	userID := request.UserID
	if strings.TrimSpace(request.Type) == "registration" {
		result := r.db.Model(&domain.Request{}).Where("id = ?", id).Update("status", 1).Update("verifier_id", verifier_id)
		if result.Error != nil {
			return result.Error.Error()
		}
		// change user role to 1 (applicant)
		s, done := updateRoleId(result, r, userID, 1)
		if done {
			return s
		}
		return "Approve request success"
	} else if strings.TrimSpace(request.Type) == "verification" {
		result := r.db.Model(&domain.Request{}).Where("id = ?", id).Update("status", 1).Update("verifier_id", verifier_id)
		if result.Error != nil {
			return result.Error.Error()
		}
		// change user role to 2 (volunteer)
		s, done := updateRoleId(result, r, userID, 2)
		if done {
			return s
		}
		// insert to volunteer_details
		departmentID := r.getDeptIdFromUser(userID)
		volunteerDetail := domain.VolunteerDetail{
			UserID:       userID,
			DepartmentID: *departmentID,
			Status:       1,
		}
		result = r.db.Create(&volunteerDetail)
		if result.Error != nil {
			return result.Error.Error()
		}
		return "Approve request success"
	}
	return "Invalid request type"
}
func (r *AdminRepository) RejectRequest(id int, verifier_id int) string {
	result := r.db.Model(&domain.Request{}).Where("id = ?", id).Update("status", 2).Update("verifier_id", verifier_id)
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

func (r *AdminRepository) getRequestByRequestID(requestID int) *domain.Request {
	var request *domain.Request
	r.db.First(&request, requestID)
	return request
}

func (r *AdminRepository) getDeptIdFromUser(id uint) *int {
	var user domain.User
	r.db.First(&user, id)
	return user.DepartmentID
}

func updateRoleId(result *gorm.DB, r *AdminRepository, userID uint, roleId int) (string, bool) {
	result = r.db.Model(&domain.User{}).Where("id = ?", userID).Update("role_id", roleId)
	if result.Error != nil {
		return result.Error.Error(), true
	}
	return "", false
}
