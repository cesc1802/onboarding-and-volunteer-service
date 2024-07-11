package usecase

import (
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

// Function submit application form, tạo ra 1 record trong bảng request
func SubmitApplicationForm(appFormDTO dto.ApplicationFormDTO) error {
	//Function parse thời gian về đúng format mình cần dùng
	dob, err := time.Parse("2006-01-02", appFormDTO.DateOfBirth)
	if err != nil {
		return err
	}

	request := domain.Request{
		UserID: appFormDTO.UserID,
		Type:   "application form",
		Status: 0, // Dat 0 lam default
	}

	// Tạo ra request trong Database
	if err := storage.CreateRequest(&request); err != nil {
		return err
	}

	// Update thông tin applicant nếu cần thiết
	user := domain.User{
		ID:      appFormDTO.UserID,
		Name:    appFormDTO.Name,
		Surname: appFormDTO.Surname,
		Gender:  appFormDTO.Sex,
		DOB:     dob,
	}

	return storage.UpdateUser(&user)
}
