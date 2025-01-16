package transport

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserIdentityUsecase struct {
	mock.Mock
}

func (m *MockUserIdentityUsecase) CreateUserIdentity(request dto.CreateUserIdentityRequest) error {
	args := m.Called(request)
	return args.Error(0)
}

func (m *MockUserIdentityUsecase) UpdateUserIdentity(id int, request dto.UpdateUserIdentityRequest) error {
	args := m.Called(id, request)
	return args.Error(0)
}

func (m *MockUserIdentityUsecase) FindUserIdentityByID(id int) (*dto.UserIdentityResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.UserIdentityResponse), args.Error(1)
}

func TestCreateUserIdentity(t *testing.T) {
	mockUsecase := new(MockUserIdentityUsecase)
	handler := NewUserIdentityHandler(mockUsecase)

	// Define the expected request body (adjusted Status to integer)
	requestBody := `{"user_id": 1, "number": "09485112345", "type":"Citizen ID", "status": 1, "expiry_date":"2024-12-31", "place_issued": "New York"}`

	// Mock the usecase method
	mockUsecase.On("CreateUserIdentity", mock.AnythingOfType("dto.CreateUserIdentityRequest")).Return(nil)

	// Create a new Gin context and request
	r := gin.Default()
	r.POST("/api/v1/applicant-identity/", handler.CreateUserIdentity)

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/applicant-identity/", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the status code and response
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User identity created successfully")

	// Assert that the mock method was called
	mockUsecase.AssertExpectations(t)
}

func TestUpdateUserIdentity(t *testing.T) {
	mockUsecase := new(MockUserIdentityUsecase)
	handler := NewUserIdentityHandler(mockUsecase)

	requestBody := `{"identity_number": "67890", "full_name": "Jane Doe"}`

	mockUsecase.On("UpdateUserIdentity", 1, mock.AnythingOfType("dto.UpdateUserIdentityRequest")).Return(nil)

	r := gin.Default()
	r.PUT("/api/v1/applicant-identity/:id", handler.UpdateUserIdentity)

	req, _ := http.NewRequest(http.MethodPut, "/api/v1/applicant-identity/1", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User identity updated successfully")

	mockUsecase.AssertExpectations(t)
}

func TestFindUserIdentity(t *testing.T) {
	mockUsecase := new(MockUserIdentityUsecase)
	handler := NewUserIdentityHandler(mockUsecase)

	mockUsecase.On("FindUserIdentityByID", 1).Return(&dto.UserIdentityResponse{
		ID:          1,
		UserID:      1,
		Number:      "0987612345",
		Type:        "Citizen ID",
		Status:      1,
		ExpiryDate:  "2024-12-31",
		PlaceIssued: "New York",
	}, nil)

	r := gin.Default()
	r.GET("/api/v1/applicant-identity/:id", handler.FindUserIdentity)

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/applicant-identity/1", nil)
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Assert the status code and response
	assert.Equal(t, http.StatusOK, w.Code)

	// Check if the response body contains the expected values
	assert.Contains(t, w.Body.String(), `"id":1`)
	assert.Contains(t, w.Body.String(), `"user_id":1`)
	assert.Contains(t, w.Body.String(), `"number":"0987612345"`)
	assert.Contains(t, w.Body.String(), `"type":"Citizen ID"`)
	assert.Contains(t, w.Body.String(), `"status":1`)
	assert.Contains(t, w.Body.String(), `"expiry_date":"2024-12-31"`)
	assert.Contains(t, w.Body.String(), `"place_issued":"New York"`)

	// Assert that the mock method was called with the correct parameters
	mockUsecase.AssertExpectations(t)
}
