package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/dto"
)

// MockCountryUsecase is a mock implementation of the CountryUsecaseInterface.
type MockCountryUsecase struct {
	mock.Mock
}

func (m *MockCountryUsecase) CreateCountry(input dto.CountryCreateDTO) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockCountryUsecase) GetCountryByID(id uint) (*dto.CountryResponseDTO, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.CountryResponseDTO), args.Error(1)
}

func (m *MockCountryUsecase) UpdateCountry(id uint, input dto.CountryUpdateDTO) error {
	args := m.Called(id, input)
	return args.Error(0)
}

func (m *MockCountryUsecase) DeleteCountry(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateCountry(t *testing.T) {
	mockUsecase := new(MockCountryUsecase)
	handler := NewCountryHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/v1/countries", handler.CreateCountry)

	input := dto.CountryCreateDTO{
		Name:   "Test Country",
		Status: 1,
	}

	response := dto.CountryCreateDTO{
		Name:   "Test Country",
		Status: 1,
	}
	mockUsecase.On("CreateCountry", input).Return(nil)
	mockUsecase.On("CreateCountry", response).Return(nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/countries", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var result map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)

	t.Logf("Response body: %s", w.Body.String())

	assert.Equal(t, "Test Country", result["name"])

	mockUsecase.AssertExpectations(t)
}

func TestGetCountryByID(t *testing.T) {
	mockUsecase := new(MockCountryUsecase)
	handler := NewCountryHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/countries/:id", handler.GetCountryByID)

	response := &dto.CountryResponseDTO{
		Name:   "Test Country",
		Status: 1,
	}

	mockUsecase.On("GetCountryByID", uint(1)).Return(response, nil).Once()

	req, _ := http.NewRequest(http.MethodGet, "/countries/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	t.Logf("Response status code: %d", w.Code)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status %d but got %d", http.StatusOK, w.Code)
	}

	var result dto.CountryResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("Error unmarshaling response: %v", err)
	}

	assert.Equal(t, response, &result)

	mockUsecase.AssertExpectations(t)
}

func TestUpdateCountry(t *testing.T) {
	mockUsecase := new(MockCountryUsecase)
	handler := NewCountryHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/countries/:id", handler.UpdateCountry)

	input := dto.CountryUpdateDTO{Name: "Updated Country"}

	mockUsecase.On("UpdateCountry", uint(1), input).Return(nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPut, "/countries/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestDeleteCountry(t *testing.T) {
	mockUsecase := new(MockCountryUsecase)
	handler := NewCountryHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/countries/:id", handler.DeleteCountry)

	mockUsecase.On("DeleteCountry", uint(1)).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/countries/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.Bytes())

	mockUsecase.AssertExpectations(t)
}
