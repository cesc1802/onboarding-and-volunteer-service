package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRoleUsecase is a mock implementation of the RoleUsecaseInterface.
type MockRoleUsecase struct {
	mock.Mock
}

func (m *MockRoleUsecase) CreateRole(input dto.RoleCreateDTO) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockRoleUsecase) GetRoleByID(id uint) (*domain.Role, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Role), args.Error(1)
}

func (m *MockRoleUsecase) UpdateRole(id uint, input dto.RoleUpdateDTO) error {
	args := m.Called(id, input)
	return args.Error(0)
}

func (m *MockRoleUsecase) DeleteRole(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestRoleHandler_CreateRole(t *testing.T) {
	mockUsecase := new(MockRoleUsecase)
	handler := NewRoleHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/v1/roles", handler.CreateRole)

	input := dto.RoleCreateDTO{
		Name:   "Admin",
		Status: "pending",
	}

	mockUsecase.On("CreateRole", input).Return(nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/roles", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockUsecase.AssertCalled(t, "CreateRole", input)
}

func TestRoleHandler_GetRoleByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(MockRoleUsecase)
	handler := NewRoleHandler(mockUsecase)

	router := gin.Default()
	router.GET("/api/v1/role/:id", handler.GetRoleByID)

	role := &domain.Role{
		Name:   "Admin",
		Status: "Pending",
	}

	mockUsecase.On("GetRoleByID", uint(1)).Return(role, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/role/1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), role.Name)
	mockUsecase.AssertCalled(t, "GetRoleByID", uint(1))
}

func TestRoleHandler_UpdateRole(t *testing.T) {
	mockUsecase := new(MockRoleUsecase)
	handler := NewRoleHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/api/v1/role/:id", handler.UpdateRole)

	input := dto.RoleUpdateDTO{
		Name:   "Admin Updated",
		Status: "Pending",
	}

	mockUsecase.On("UpdateRole", uint(1), input).Return(nil)

	body, _ := json.Marshal(input)
	req, _ := http.NewRequest(http.MethodPut, "/api/v1/role/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertCalled(t, "UpdateRole", uint(1), input)
}

func TestRoleHandler_DeleteRole(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(MockRoleUsecase)
	handler := NewRoleHandler(mockUsecase)

	router := gin.Default()
	router.DELETE("/api/v1/role/:id", handler.DeleteRole)

	mockUsecase.On("DeleteRole", uint(1)).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/role/1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
	mockUsecase.AssertCalled(t, "DeleteRole", uint(1))
}
