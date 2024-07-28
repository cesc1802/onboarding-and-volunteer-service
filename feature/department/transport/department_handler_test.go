// department_handler_test.go
package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDepartmentUsecase is a mock implementation of the DepartmentUsecase
type MockDepartmentUsecase struct {
	mock.Mock
}

// CreateDepartment is a mock method for creating a department.
func (m *MockDepartmentUsecase) CreateDepartment(input dto.DepartmentCreateDTO) error {
	args := m.Called(input)
	return args.Error(0)
}

// GetDepartmentByID is a mock method for retrieving a department by ID.
func (m *MockDepartmentUsecase) GetDepartmentByID(id uint) (*domain.Department, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Department), args.Error(1)
}

// UpdateDepartment is a mock method for updating a department.
func (m *MockDepartmentUsecase) UpdateDepartment(id uint, input dto.DepartmentUpdateDTO) error {
	args := m.Called(id, input)
	return args.Error(0)
}

// DeleteDepartment is a mock method for deleting a department.
func (m *MockDepartmentUsecase) DeleteDepartment(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

// TestCreateDepartment tests the CreateDepartment handler function.
func TestCreateDepartment(t *testing.T) {
	mockUsecase := new(MockDepartmentUsecase)    // Create a new mock use case.
	handler := NewDepartmentHandler(mockUsecase) // Create a new handler with the mock use case.

	// Set up Gin in test mode and define the route.
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/v1/departments", handler.CreateDepartment)

	// Define the input and expected response.
	input := dto.DepartmentCreateDTO{
		Name:    "HR",
		Address: "123 HR Street",
		Status:  123,
	}
	response := dto.DepartmentCreateDTO{
		Name:    "HR",
		Address: "123 HR Street",
		Status:  123,
	}

	// Set the expected call on the mock use case.
	mockUsecase.On("CreateDepartment", input).Return(response, nil)

	// Create a new HTTP request and record the response.
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("POST", "/api/v1/departments", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response status code and body.
	assert.Equal(t, http.StatusCreated, w.Code)
	var result dto.DepartmentResponseDTO
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, response, result)
	assert.JSONEq(t, `{"Name":"HR","Address":"123 HR Street","Status":123}`, w.Body.String())

	// Ensure the mock use case's expectations were met.
	mockUsecase.AssertExpectations(t)
}

// TestGetDepartmentByID tests the GetDepartmentByID handler function.
func TestGetDepartmentByID(t *testing.T) {
	mockUsecase := new(MockDepartmentUsecase)    // Create a new mock use case.
	handler := NewDepartmentHandler(mockUsecase) // Create a new handler with the mock use case.

	// Set up Gin in test mode and define the route.
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/v1/departments/:id", handler.GetDepartmentByID)

	// Define the expected response.
	response := dto.DepartmentResponseDTO{Name: "Test Department"}

	// Set the expected call on the mock use case.
	mockUsecase.On("GetDepartmentByID", uint(1)).Return(response, nil)

	// Create a new HTTP request and record the response.
	req, _ := http.NewRequest("GET", "/api/v1/departments/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response status code and body.
	assert.Equal(t, http.StatusOK, w.Code)
	var result dto.DepartmentResponseDTO
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, response, result)
	assert.JSONEq(t, `{"ID":1,"Name":"Finance","Address":"456 Finance Street","Status":456}`, w.Body.String())

	// Ensure the mock use case's expectations were met.
	mockUsecase.AssertExpectations(t)
}

// TestUpdateDepartment tests the UpdateDepartment handler function.
func TestUpdateDepartment(t *testing.T) {
	mockUsecase := new(MockDepartmentUsecase)    // Create a new mock use case.
	handler := NewDepartmentHandler(mockUsecase) // Create a new handler with the mock use case.

	// Set up Gin in test mode and define the route.
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/api/v1/departments/:id", handler.UpdateDepartment)

	// Define the input and expected result.
	input := dto.DepartmentUpdateDTO{
		Name:    "IT Updated",
		Address: "789 IT Street Updated",
		Status:  789,
	}

	// Set the expected call on the mock use case.
	mockUsecase.On("UpdateDepartment", uint(1), input).Return(nil, nil)

	// Create a new HTTP request and record the response.
	body, _ := json.Marshal(input)
	req, _ := http.NewRequest("PUT", "/api/v1/departments/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response status code and body.
	assert.Equal(t, http.StatusOK, w.Code)
	var result dto.DepartmentResponseDTO
	json.Unmarshal(w.Body.Bytes(), &result)
	assert.Equal(t, nil, result)
	assert.JSONEq(t, `{"ID":1,"Name":"IT Updated","Address":"789 IT Street Updated","Status":"Inactive"}`, w.Body.String())
	// Ensure the mock use case's expectations were met.
	mockUsecase.AssertExpectations(t)
}

// TestDeleteDepartment tests the DeleteDepartment handler function.
func TestDeleteDepartment(t *testing.T) {
	mockUsecase := new(MockDepartmentUsecase)    // Create a new mock use case.
	handler := NewDepartmentHandler(mockUsecase) // Create a new handler with the mock use case.

	// Set up Gin in test mode and define the route.
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/api/v1/departments/:id", handler.DeleteDepartment)

	// Set the expected call on the mock use case.
	mockUsecase.On("DeleteDepartment", uint(1)).Return(nil)

	// Create a new HTTP request and record the response.
	req, _ := http.NewRequest("DELETE", "/api/v1/departments/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the response status code and body.
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.Bytes())

	// Ensure the mock use case's expectations were met.
	mockUsecase.AssertExpectations(t)
}
