package transport

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVolunteerUsecase struct {
	mock.Mock
}

func (m *MockVolunteerUsecase) CreateVolunteer(input dto.VolunteerCreateDTO) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockVolunteerUsecase) UpdateVolunteer(id int, input dto.VolunteerUpdateDTO) error {
	args := m.Called(id, input)
	return args.Error(0)
}

func (m *MockVolunteerUsecase) DeleteVolunteer(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockVolunteerUsecase) FindVolunteerByID(id int) (*dto.VolunteerResponseDTO, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.VolunteerResponseDTO), args.Error(1)
}

func TestCreateVolunteer(t *testing.T) {
	// Setup
	mockUsecase := new(MockVolunteerUsecase)
	handler := NewVolunteerHandler(mockUsecase)
	router := gin.Default()
	router.POST("/api/v1/volunteer", handler.CreateVolunteer)

	// Mock behavior
	mockUsecase.On("CreateVolunteer", mock.Anything).Return(nil)

	// Correct payload
	payload := `{
		"user_id": 1,
		"department_id": 2,
		"status": 1
	}`
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/volunteer", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	// Recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	// Use JSONEq to compare JSON response body correctly
	expectedResponse := `{"message": "User created successfully"}`
	assert.JSONEq(t, expectedResponse, w.Body.String(), "Response body does not match the expected value.")
	mockUsecase.AssertExpectations(t)
}

func TestUpdateVolunteer(t *testing.T) {
	// Setup
	mockUsecase := new(MockVolunteerUsecase)
	handler := NewVolunteerHandler(mockUsecase)
	router := gin.Default()
	router.PUT("/api/v1/volunteer/:id", handler.UpdateVolunteer)

	// Mock behavior
	mockUsecase.On("UpdateVolunteer", 1, mock.Anything).Return(nil)

	// Request
	payload := `{"user_id": 1,
		         "department_id": 2,
		         "status": 2}`
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/volunteer/1", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	// Recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Volunteer updated successfully")
	mockUsecase.AssertExpectations(t)
}

func TestDeleteVolunteer(t *testing.T) {
	// Setup
	mockUsecase := new(MockVolunteerUsecase)
	handler := NewVolunteerHandler(mockUsecase)
	router := gin.Default()
	router.DELETE("/api/v1/volunteer/:id", handler.DeleteVolunteer)

	// Mock behavior
	mockUsecase.On("DeleteVolunteer", 1).Return(nil)

	// Request
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/volunteer/1", nil)

	// Recorder
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Volunteer deleted successfully")
	mockUsecase.AssertExpectations(t)
}

func TestFindVolunteerByID(t *testing.T) {
	// Setup
	mockUsecase := new(MockVolunteerUsecase)
	handler := NewVolunteerHandler(mockUsecase)

	// Mock behavior for a successful response
	mockUsecase.On("FindVolunteerByID", 1).Return(&dto.VolunteerResponseDTO{
		ID:           1,
		UserID:       123,
		DepartmentID: 456,
		Status:       123,
	}, nil)

	r := gin.Default()
	r.GET("/api/v1/volunteer/:id", handler.FindVolunteerByID)

	// Request
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/volunteer/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)

	// Check that the response contains the expected fields and values (raw response check)
	assert.Contains(t, w.Body.String(), `"id":1`)
	assert.Contains(t, w.Body.String(), `"user_id":123`)
	assert.Contains(t, w.Body.String(), `"department_id":456`)
	assert.Contains(t, w.Body.String(), `"status":123`)

	// Decode the response body into a struct and validate fields
	var response dto.VolunteerResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, 1, response.ID)
	assert.Equal(t, 123, response.UserID)
	assert.Equal(t, 456, response.DepartmentID)
	assert.Equal(t, 123, response.Status)

	mockUsecase.AssertExpectations(t)
}
