package storage

import "github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"

//CRUD với các user là applicant
func CreateUser(user *domain.User) error {
	return DB.Create(user).Error
}

func UpdateUser(user *domain.User) error {
	return DB.Save(user).Error
}

func DeleteUser(userID uint) error {
	return DB.Delete(&domain.User{}, userID).Error
}
