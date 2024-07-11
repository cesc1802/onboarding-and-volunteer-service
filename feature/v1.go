package feature

import (
	"log"
	"net/http"
	"os"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/transport"
	"github.com/cesc1802/share-module/system"
	"github.com/gin-gonic/gin"
)

func RegisterHandlerV1(mono system.Service) {
	router := mono.Router()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"data": "success",
		})
	})
	v1 := router.Group("/api/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/login", func(c *gin.Context) {

		})

		auth.POST("/register", func(c *gin.Context) {

		})
	}

	//Kết nối database
	dsn := os.Getenv("DB_URL")
	if err := storage.InitDB(dsn); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	//Đăng ký các route dành cho applicant
	applicantGroup := v1.Group("/applicant")
	{
		applicantGroup.POST("/", transport.SignupUser)
		applicantGroup.PUT("/", transport.UpdateUser)
		applicantGroup.DELETE("/:id", transport.DeleteUser)
		applicantGroup.POST("/application", transport.SubmitApplicationForm)
	}

}
