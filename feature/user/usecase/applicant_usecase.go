package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

type ApplicantUsecaseInterface interface {
	CreateApplicant(request dto.ApplicantCreateDTO) error
	UpdateApplicant(id int, request dto.ApplicantUpdateDTO) error
	DeleteApplicant(id int) error
	FindApplicantByID(id int) (*dto.ApplicantResponseDTO, error)
}

type ApplicantUsecase struct {
	ApplicantRepo storage.ApplicantRepositoryInterface
}

func NewApplicantUsecase(userRepo storage.ApplicantRepositoryInterface) *ApplicantUsecase {
	return &ApplicantUsecase{ApplicantRepo: userRepo}
}

func (u *ApplicantUsecase) CreateApplicant(request dto.ApplicantCreateDTO) error {
	user := &domain.ApplicantDomain{
		Email:   request.Email,
		Name:    request.Name,
		Surname: request.Surname,
	}
	return u.ApplicantRepo.CreateApplicant(user)
}

func (u *ApplicantUsecase) UpdateApplicant(id int, request dto.ApplicantUpdateDTO) error {
	user, err := u.ApplicantRepo.FindApplicantByID(id)
	if err != nil {
		return err
	}
	user.Email = request.Email
	user.Name = request.Name
	user.Surname = request.Surname
	user.Gender = request.Gender
	user.RoleID = &request.RoleID
	return u.ApplicantRepo.UpdateApplicant(user)
}

func (u *ApplicantUsecase) DeleteApplicant(id int) error {
	return u.ApplicantRepo.DeleteApplicant(id)
}

func (u *ApplicantUsecase) FindApplicantByID(id int) (*dto.ApplicantResponseDTO, error) {
	user, err := u.ApplicantRepo.FindApplicantByID(id)
	if err != nil {
		return nil, err
	}

	response := &dto.ApplicantResponseDTO{
		ID:      user.ID,
		RoleID:  *user.RoleID,
		Email:   user.Email,
		Name:    user.Name,
		Surname: user.Surname,
		Gender:  user.Gender,
	}
	return response, nil
}
