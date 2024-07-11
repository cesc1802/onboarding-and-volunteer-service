package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
)

// Tạo trong bảng request 1 record khi điền xong application form
func CreateRequest(request *domain.Request) error {
	return DB.Create(request).Error
}
