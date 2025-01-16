package usecase

import (
	"log"
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/storage"
)

type VolunteerUsecaseInterface interface {
	CreateVolunteer(input dto.VolunteerCreateDTO) error
	UpdateVolunteer(id int, input dto.VolunteerUpdateDTO) error
	DeleteVolunteer(id int) error
	FindVolunteerByID(id int) (*dto.VolunteerResponseDTO, error)
}

type VolunteerUsecase struct {
	VolunteerRepo storage.VolunteerRepositoryInterface
}

func NewVolunteerUsecase(volunteerRepo storage.VolunteerRepositoryInterface) *VolunteerUsecase {
	return &VolunteerUsecase{VolunteerRepo: volunteerRepo}
}

func (u *VolunteerUsecase) CreateVolunteer(input dto.VolunteerCreateDTO) error {
	dob, err := time.Parse("2006-01-02", input.Dob) // Expected format: YYYY-MM-DD
	if err != nil {
		log.Println("Invalid Dob format:", input.Dob)
		return nil
	}

	volunteer := &domain.Volunteer{
		UserID:             input.UserID,
		DepartmentID:       &input.DepartmentID,
		Dob:                dob,
		Mobile:             input.Mobile,
		CountryID:          &input.CountryID,
		ResidentCountryID:  &input.ResidentCountryID,
		Avatar:             input.Avatar,
		VerificationStatus: input.VerificationStatus,
		Status:             input.Status,
	}
	return u.VolunteerRepo.CreateVolunteer(volunteer)
}

func (u *VolunteerUsecase) UpdateVolunteer(id int, input dto.VolunteerUpdateDTO) error {
	volunteer, err := u.VolunteerRepo.FindVolunteerByID(id)
	if err != nil {
		return err
	}
	volunteer.DepartmentID = &input.DepartmentID
	volunteer.Status = input.Status

	return u.VolunteerRepo.UpdateVolunteer(volunteer)
}

func (u *VolunteerUsecase) DeleteVolunteer(id int) error {
	return u.VolunteerRepo.DeleteVolunteer(id)
}

func (u *VolunteerUsecase) FindVolunteerByID(id int) (*dto.VolunteerResponseDTO, error) {
	volunteer, err := u.VolunteerRepo.FindVolunteerByID(id)
	if err != nil {
		return nil, err
	}
	response := &dto.VolunteerResponseDTO{
		ID:                 volunteer.ID,
		UserID:             volunteer.UserID,
		DepartmentID:       *volunteer.DepartmentID,
		Dob:                volunteer.Dob.Format("2006-01-02"),
		Mobile:             volunteer.Mobile,
		Avatar:             volunteer.Avatar,
		CountryID:          *volunteer.CountryID,
		ResidentCountryID:  *volunteer.ResidentCountryID,
		VerificationStatus: volunteer.VerificationStatus,
		Status:             volunteer.Status,
	}
	return response, nil
}
