package transport

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/usecase"
	"github.com/gin-gonic/gin"
)

// RoleHandler handles the HTTP requests for roles.
type RoleHandler struct {
	usecase usecase.RoleUsecaseInterface
}

// NewRoleHandler creates a new instance of RoleHandler.
func NewRoleHandler(usecase usecase.RoleUsecaseInterface) *RoleHandler {
	return &RoleHandler{usecase: usecase}
}

// CreateRole handles the HTTP POST request to create a new role.
// CreateRole godoc
// @Summary Create role
// @Description Create role
// @Produce json
// @Tags role
// @Param request body dto.RoleCreateDTO true "Create Role Request"
// @Success 201 {object} domain.Role
// @Router /api/v1/role/ [post]
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var input dto.RoleCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Creating role:", input) // Add print to debug
	err := h.usecase.CreateRole(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "role created successfully",
		"name":    input.Name,
		"status":  input.Status,
	})
}

// GetRoleByID handles the HTTP GET request to retrieve a role by its ID.
// GetRoleByID godoc
// @Summary Get role by ID
// @Description Get role by ID
// @Produce json
// @Tags role
// @Param id path int true "Role ID"
// @Success 200 {object} domain.Role
// @Router /api/v1/role/{id} [get]
func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := h.usecase.GetRoleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// UpdateRole handles the HTTP PUT request to update a role.
// UpdateRole godoc
// @Summary Update role
// @Description Update role
// @Produce json
// @Tags role
// @Param id path int true "Role ID"
// @Param request body dto.RoleUpdateDTO true "Update Role Request"
// @Success 200 {object} domain.Role
// @Router /api/v1/role/{id} [put]
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	var input dto.RoleUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.UpdateRole(uint(id), input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "role updated successfully",
		"name":    input.Name,
		"status":  input.Status,
	})
}

// DeleteRole handles the HTTP DELETE request to delete a role.
// DeleteRole godoc
// @Summary Delete role
// @Description Delete role
// @Produce json
// @Tags role
// @Param id path int true "Role ID"
// @Success 204
// @Router /api/v1/role/{id} [delete]
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	err = h.usecase.DeleteRole(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "role updated successfully"})
}
