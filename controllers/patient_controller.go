package controllers

import (
	"net/http"
	"strconv"

	"healthcare-app/healthcare-backend/config"
	"healthcare-app/healthcare-backend/models"

	"github.com/gin-gonic/gin"
)

// CreatePatient creates a new patient
// @Summary Create a new patient
// @Description Create a new patient and store in the database
// @Accept  json
// @Produce  json
// @Param patient body models.Patient true "Patient Data"
// @Success 201 {object} models.Patient "Patient Created"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /patients [post]
func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, patient)
}

// GetPatients godoc
// @Summary      Get all patients
// @Description  Retrieve a paginated list of patients
// @Tags         patients
// @Accept       json
// @Produce      json
// @Param        page  query int false "Page number"
// @Param        limit query int false "Items per page"
// @Success      200   {array} models.Patient
// @Router       /patients [get]
func GetPatients(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var patients []models.Patient
	if err := config.DB.Preload("Appointments").
		Limit(limit).
		Offset(offset).
		Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, patients)
}

func GetPatientByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	var patient models.Patient
	if err := config.DB.Preload("Appointments").First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}
