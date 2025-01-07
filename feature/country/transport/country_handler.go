package transport

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/usecase"
	"github.com/gin-gonic/gin"
)

// CountryHandler handles the HTTP requests for countries.
type CountryHandler struct {
	usecase usecase.CountryUsecaseInterface
}

// NewCountryHandler creates a new instance of CountryHandler.
func NewCountryHandler(usecase usecase.CountryUsecaseInterface) *CountryHandler {
	return &CountryHandler{usecase: usecase}
}

// CreateCountry handles the HTTP POST request to create a new country.
// CreateCountry godoc
// @Summary Create a new country
// @Description Create a new country
// @Accept json
// @Produce json
// @Tags country
// @Param country body dto.CountryCreateDTO true "Country data"
// @Success 201 {object} domain.Country
// @Router /api/v1/countries [post]
func (h *CountryHandler) CreateCountry(c *gin.Context) {
	fmt.Println("CreateCountry handler called") // Log for debugging
	var input dto.CountryCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.CreateCountry(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create country successfully",
		"name":    input.Name,
		"status":  input.Status,
	})
}

// GetCountryByID handles the HTTP GET request to retrieve a country by its ID.
// GetCountryByID godoc
// @Summary Get country by ID
// @Description Get country by ID
// @Produce json
// @Tags country
// @Param id path int true "Country ID"
// @Success 200 {object} domain.Country
// @Router /api/v1/countries/{id} [get]
func (h *CountryHandler) GetCountryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	country, err := h.usecase.GetCountryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
		return
	}

	c.JSON(http.StatusOK, country)
}

// UpdateCountry handles the HTTP PUT request to update a country.
// UpdateCountry godoc
// @Summary Update country
// @Description Update country
// @Accept json
// @Produce json
// @Tags country
// @Param id path int true "Country ID"
// @Param country body dto.CountryUpdateDTO true "Country data"
// @Success 200 {object} domain.Country
// @Router /api/v1/countries/{id} [put]
func (h *CountryHandler) UpdateCountry(c *gin.Context) {
	var input dto.CountryUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.usecase.UpdateCountry(uint(id), input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Country updated successfully",
		"name":    input.Name,
		"status":  input.Status})
}

// DeleteCountry handles the HTTP DELETE request to delete a country.
// DeleteCountry godoc
// @Summary Delete country
// @Description Delete country
// @Tags country
// @Param id path int true "Country ID"
// @Success 204
// @Router /api/v1/countries/{id} [delete]
func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	err = h.usecase.DeleteCountry(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Country successfully",
	})
}
