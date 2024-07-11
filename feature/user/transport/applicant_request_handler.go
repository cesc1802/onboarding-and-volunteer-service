package transport

import (
	"net/http"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"
	"github.com/gin-gonic/gin"
)

// submit thông tin của request vào table request khi mà người dùng điền xong application form
func SubmitApplicationForm(c *gin.Context) {

	var appFormDTO dto.ApplicationFormDTO

	if err := c.ShouldBindJSON(&appFormDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.SubmitApplicationForm(appFormDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Application form submitted successfully"})
}
