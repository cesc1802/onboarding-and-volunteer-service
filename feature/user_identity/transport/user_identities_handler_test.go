package transport

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserIdentityUsecase struct {
	mock.Mock
}

func (m *MockUserIdentityUsecase) CreateUserIdentity(input dto.CreateUserIdentityRequest) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockUserIdentityUsecase) UpdateUserIdentity(id int, input dto.UpdateUserIdentityRequest) error {
	args := m.Called(id, input)
	return args.Error(0)
}

func (m *MockUserIdentityUsecase) FindUserIdentityByID(id int) (*dto.UserIdentityResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.UserIdentityResponse), args.Error(1)
}

func TestCreateUserIdentity(t *testing.T) {
	mockUsecase := new(MockUserIdentityUsecase)
	handler := NewUserIdentityHandler(mockUsecase)
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/v1/user-identity", handler.CreateUserIdentity)

	t.Run("success", func(t *testing.T) {
		mockInput := dto.CreateUserIdentityRequest{
			UserID:      2,
			Number:      "123456789",
			Type:        "Citizen ID",
			Status:      0,
			ExpiryDate:  "12-12-2025",
			PlaceIssued: "Some city",
		}
		mockUsecase.On("CreateUserIdentity", mockInput).Return(nil)

		body := `{"UserID":2,"Number":"123456789","Type":"Citizen ID","Status":"Approved","Expiry Date":"12-12-2025","PlaceIssued":"Some city"}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/user-identity", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		body := `{"UserID":abc,"Number":"123456789","Type":"Citizen ID","Status":"Approved","Expiry Date":"12-12-2025","PlaceIssued":"Some city"}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/user-identity", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestUpdateUserIdentity(t *testing.T) {
	mockUsecase := new(MockUserIdentityUsecase)
	handler := NewUserIdentityHandler(mockUsecase)
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/api/v1/user-identity/:id", handler.UpdateUserIdentity)

	t.Run("success", func(t *testing.T) {
		mockInput := dto.UpdateUserIdentityRequest{
			UserID:      2,
			Number:      "123888789",
			Type:        "Passport",
			Status:      0,
			ExpiryDate:  "12-12-2025",
			PlaceIssued: "Another city",
		}
		mockUsecase.On("UpdateUserIdentity", 1, mockInput).Return(nil)

		body := `{"UserID":2,"Number":"123888789","Type":"Passport","Status":"Approved","Expiry Date":"12-12-2025","PlaceIssued":"Another city"}`
		req, err := http.NewRequest(http.MethodPut, "/api/v1/user-identity/1", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		body := `{"UserID":abc,"Number":"123888789","Type":"Passport","Status":"Approved","Expiry Date":"12-12-2025","PlaceIssued":"Another city"}`
		req, err := http.NewRequest(http.MethodPut, "/api/v1/user-identity/1", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestFindUserIdentity(t *testing.T) {
	mockUsecase := new(MockUserIdentityUsecase)
	handler := NewUserIdentityHandler(mockUsecase)
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/api/v1/user-identity/:id", handler.FindUserIdentity)

	t.Run("success", func(t *testing.T) {
		mockVolunteer := &dto.UserIdentityResponse{}
		mockUsecase.On("FindUserIdentity", 1).Return(mockVolunteer, nil)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/user-identity/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockUsecase.On("FindUserIdentity", 1).Return(nil, errors.New("user identity not found"))

		req, err := http.NewRequest(http.MethodGet, "/api/v1/user-identity/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
