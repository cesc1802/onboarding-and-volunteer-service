package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

type AdminUsecaseInterface interface {
	GetListPendingRequest() (*dto.ListRequest, string)
	GetPendingRequestById(id int) (*dto.RequestResponse, string)
	GetListRequest() (*dto.ListRequest, string)
	GetRequestById(id int) (*dto.RequestResponse, string)
	ApproveRequest(id int, verifier_id int) string
	RejectRequest(id int, verifier_id int) string
	AddRejectNotes(id int, notes string) string
	DeleteRequest(id int) string
}

type AdminUsecase struct {
	repo storage.AdminRepositoryInterface
}

func NewAdminUsecase(repo storage.AdminRepositoryInterface) *AdminUsecase {
	return &AdminUsecase{repo: repo}
}
func (u *AdminUsecase) GetListPendingRequest() (*dto.ListRequest, string) {
	requests, msg := u.repo.GetListPendingRequest()
	if requests != nil {
		return &dto.ListRequest{
			Requests: requests,
		}, msg
	} else {
		msg = "No request found"
	}
	return nil, msg
}
func (u *AdminUsecase) GetPendingRequestById(id int) (*dto.RequestResponse, string) {
	request, msg := u.repo.GetPendingRequestByID(id)
	if request != nil {
		return &dto.RequestResponse{
			ID:          request.ID,
			UserID:      request.UserID,
			Type:        request.Type,
			Status:      request.Status,
			RejectNotes: request.RejectNotes,
			VerifierID:  request.VerifierID,
		}, msg
	} else {
		msg = "Request not found"
	}
	return nil, msg
}

func (u *AdminUsecase) GetListRequest() (*dto.ListRequest, string) {
	requests, msg := u.repo.GetListAllRequest()
	if requests != nil {
		return &dto.ListRequest{
			Requests: requests,
		}, msg
	} else {
		msg = "No request found"
	}
	return nil, msg
}
func (u *AdminUsecase) GetRequestById(id int) (*dto.RequestResponse, string) {
	request, msg := u.repo.GetRequestByID(id)
	if request != nil {
		return &dto.RequestResponse{
			ID:          request.ID,
			UserID:      request.UserID,
			Type:        request.Type,
			Status:      request.Status,
			RejectNotes: request.RejectNotes,
			VerifierID:  request.VerifierID,
		}, msg
	} else {
		msg = "Request not found"
	}
	return nil, msg
}

func (u *AdminUsecase) ApproveRequest(id int, verifier_id int) string {
	return u.repo.ApproveRequest(id, verifier_id)
}
func (u *AdminUsecase) RejectRequest(id int, verifier_id int) string {
	return u.repo.RejectRequest(id, verifier_id)
}
func (u *AdminUsecase) AddRejectNotes(id int, notes string) string {
	return u.repo.AddRejectNotes(id, notes)
}
func (u *AdminUsecase) DeleteRequest(id int) string {
	return u.repo.DeleteRequest(id)
}
