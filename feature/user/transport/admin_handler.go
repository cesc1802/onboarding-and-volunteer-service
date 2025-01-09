package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	usecase usecase.AdminUsecaseInterface
}

func NewAuthenticationHandler(usecase usecase.AdminUsecaseInterface) *AdminHandler {
	return &AdminHandler{usecase: usecase}
}

// GetListPendingRequest godoc
// @Summary Get list pending request
// @Description Get list pending request
// @Produce json
// @Tags admin
// @Security bearerToken
// @Success 200 {object} dto.ListRequest{}
// @Router /api/v1/admin/list-pending-request [get]
func (h *AdminHandler) GetListPendingRequest(c *gin.Context) {
	resp, msg := h.usecase.GetListPendingRequest()
	if msg != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetPendingRequestById godoc
// @Summary Get pending request by ID
// @Description Get pending request by ID
// @Produce json
// @Tags admin
// @Param id path int true "Request ID"
// @Success 200 {object} dto.RequestResponse{}
// @Security bearerToken
// @Router /api/v1/admin/pending-request/{id} [get]
func (h *AdminHandler) GetPendingRequestById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	resp, msg := h.usecase.GetPendingRequestById(id)
	if msg != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetListRequest godoc
// @Summary Get list request
// @Description Get list request
// @Produce json
// @Tags admin
// @Security bearerToken
// @Success 200 {object} dto.ListRequest{}
// @Router /api/v1/admin/list-request [get]
func (h *AdminHandler) GetListRequest(c *gin.Context) {
	resp, msg := h.usecase.GetListRequest()
	if msg != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetRequestById godoc
// @Summary Get request by ID
// @Description Get request by ID
// @Produce json
// @Tags admin
// @Param id path int true "Request ID"
// @Success 200 {object} dto.RequestResponse{}
// @Security bearerToken
// @Router /api/v1/admin/request/{id} [get]
func (h *AdminHandler) GetRequestById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	resp, msg := h.usecase.GetRequestById(id)
	if msg != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// ApproveRequest godoc
// @Summary Approve request
// @Description Approve request
// @Produce json
// @Tags admin
// @Param id path int true "Request ID"
// @Success 200 string message
// @Security bearerToken
// @Router /api/v1/admin/approve-request/{id} [post]
func (h *AdminHandler) ApproveRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	msg := h.usecase.ApproveRequest(id, userId.(int))
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

// RejectRequest godoc
// @Summary Reject request
// @Description Reject request
// @Produce json
// @Tags admin
// @Param id path int true "Request ID"
// @Success 200 string message
// @Security bearerToken
// @Router /api/v1/admin/reject-request/{id} [post]
func (h *AdminHandler) RejectRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	msg := h.usecase.RejectRequest(id, userId.(int))
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

// AddRejectNotes godoc
// @Summary Add reject notes
// @Description Add reject notes
// @Produce json
// @Tags admin
// @Param id path int true "Request ID"
// @Param notes body dto.AddRejectNoteRequest true "Add Reject Note Request"
// @Success 200 string message
// @Security bearerToken
// @Router /api/v1/admin/add-reject-notes/{id} [post]
func (h *AdminHandler) AddRejectNotes(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	var req dto.AddRejectNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msg := h.usecase.AddRejectNotes(id, req.Notes)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

// DeleteRequest godoc
// @Summary Delete request
// @Description Delete request
// @Produce json
// @Tags admin
// @Param id path int true "Request ID"
// @Success 200 string message
// @Security bearerToken
// @Router /api/v1/admin/delete-request/{id} [delete]
func (h *AdminHandler) DeleteRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	msg := h.usecase.DeleteRequest(id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}
