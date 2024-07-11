package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

// Tao 1 user lần đầu tiên
func SignupUser(userDTO dto.UserSignupDTO) error {
	user := domain.User{
		Email:   userDTO.Email,
		Name:    userDTO.Name,
		Surname: userDTO.Surname,
	}
	return storage.CreateUser(&user)
}

// Update user khi điền xong application form
func UpdateUser(userDTO dto.UserUpdateDTO) error {
	user := domain.User{
		ID:      userDTO.ID,
		Email:   userDTO.Email,
		Name:    userDTO.Name,
		Surname: userDTO.Surname,
	}
	return storage.UpdateUser(&user)
}

// Delete
func DeleteUser(userID uint) error {
	return storage.DeleteUser(userID)
}
