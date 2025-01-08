package storage

import (
	"errors"
	"log"
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenticationStore interface {
	GetUserByEmail(email string, password string) (*domain.User, string)
	RegisterUser(request *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error)
}

type AuthenticationRepository struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) *AuthenticationRepository {
	return &AuthenticationRepository{db: db}
}

func (r *AuthenticationRepository) GetUserByEmail(email string, password string) (*domain.User, string) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, "User not found"
		}
		return nil, err.Error()
	}
	if user.Status == 0 {
		return nil, "User is inactive"
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "Password is incorrect"
	}
	return &user, ""
}

func (r *AuthenticationRepository) RegisterUser(request *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error) {

	var existingUser domain.User
	if err := r.db.Where("email = ?", request.Email).First(&existingUser).Error; err == nil {
		log.Println("Email already registered:", request.Email)
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	dob, err := time.Parse("2006-01-02", request.Dob) // Expected format: YYYY-MM-DD
	if err != nil {
		log.Println("Invalid Dob format:", request.Dob)
		return nil, errors.New("invalid date of birth format, expected YYYY-MM-DD")
	}

	user := domain.User{
		Email:              request.Email,
		Name:               request.Name,
		Surname:            request.Surname,
		Password:           string(hashedPassword),
		Gender:             request.Gender,
		Dob:                dob,
		Mobile:             request.Mobile,
		CountryID:          &request.CountryID,
		ResidentCountryID:  &request.ResidentCountryID,
		Avatar:             request.Avatar,
		VerificationStatus: request.VerificationStatus,
		Status:             request.Status,
		RoleID:             &request.RoleID,
		DepartmentID:       &request.DepartmentID,
	}

	if err := r.db.Create(&user).Error; err != nil {
		log.Println("Error creating user:", err)
		return nil, err
	}
	log.Println("User created successfully:", user)

	response := &dto.RegisterUserResponse{
		Message: "User registered successfully",
		User: &dto.RegisterUserRequest{
			ID:                 user.ID,
			Email:              user.Email,
			Name:               user.Name,
			Password:           user.Password,
			Surname:            user.Surname,
			Gender:             user.Gender,
			Dob:                user.Dob.Format("11/01/2000"),
			Mobile:             user.Mobile,
			Avatar:             user.Avatar,
			CountryID:          *user.CountryID,
			ResidentCountryID:  *user.ResidentCountryID,
			VerificationStatus: user.VerificationStatus,
			Status:             user.Status,
			RoleID:             *user.RoleID,
			DepartmentID:       *user.DepartmentID,
		},
	}

	return response, nil
}
