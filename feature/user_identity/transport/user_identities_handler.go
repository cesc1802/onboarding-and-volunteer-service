package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/usecase"
	"github.com/gin-gonic/gin"
)

type UserIdentityHandler struct {
	UserIdentityUsecase usecase.UserIdentityUsecaseInterface
}

func NewUserIdentityHandler(userIdentityUsecase usecase.UserIdentityUsecaseInterface) *UserIdentityHandler {
	return &UserIdentityHandler{UserIdentityUsecase: userIdentityUsecase}
}

// CreateUserIdentity godoc
// @Summary Create user identity
// @Description Create user identity
// @Produce json
// @Tags user_identity
// @Param request body dto.CreateUserIdentityRequest true "Create User Identity Request"
// @Success 201 {string} message "User identity created successfully"
// @Router /api/v1/applicant-identity/ [post]
func (h *UserIdentityHandler) CreateUserIdentity(c *gin.Context) {
	var request dto.CreateUserIdentityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.UserIdentityUsecase.CreateUserIdentity(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User identity created successfully"})
}

// UpdateUserIdentity godoc
// @Summary Update user identity
// @Description Update user identity
// @Produce json
// @Tags user_identity
// @Param id path int true "Identity ID"
// @Param request body dto.UpdateUserIdentityRequest true "Update User Identity Request"
// @Success 200 {string} message "User identity updated successfully"
// @Router /api/v1/applicant-identity/{id} [put]
func (h *UserIdentityHandler) UpdateUserIdentity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid identity ID"})
		return
	}

	var request dto.UpdateUserIdentityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserIdentityUsecase.UpdateUserIdentity(id, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User identity updated successfully"})
}

// FindUserIdentity godoc
// @Summary Find user identity
// @Description Find user identity
// @Produce json
// @Tags user_identity
// @Param id path int true "Identity ID"
// @Success 200 {object} dto.UserIdentityResponse
// @Router /api/v1/applicant-identity/{id} [get]
func (h *UserIdentityHandler) FindUserIdentity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid identity ID"})
		return
	}

	identity, err := h.UserIdentityUsecase.FindUserIdentityByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, identity)
}
