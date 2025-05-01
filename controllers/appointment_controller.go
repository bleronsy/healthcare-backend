package controllers

import (
	"net/http"
	"strconv"

	"healthcare-app/healthcare-backend/config"
	"healthcare-app/healthcare-backend/models"

	"github.com/gin-gonic/gin"
)

// CreateAppointment godoc
// @Summary      Create a new appointment for a patient
// @Description  Create a new appointment and link it to an existing patient
// @Tags         appointments
// @Accept       json
// @Produce      json
// @Param        id path int true "Patient ID"
// @Param        appointment body models.Appointment true "Appointment Data"
// @Success      201 {object} models.Appointment "Appointment Created"
// @Failure      400 {object} map[string]string "Bad Request"
// @Failure      404 {object} map[string]string "Patient Not Found"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /patients/{id}/appointments [post]
func CreateAppointment(c *gin.Context) {
	patientID := c.Param("id")

	var patient models.Patient
	if err := config.DB.First(&patient, patientID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment.PatientID = patient.ID

	if err := config.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, appointment)
}

// GetAppointments godoc
// @Summary      Get all appointments
// @Description  Retrieve a paginated list of appointments
// @Tags         appointments
// @Accept       json
// @Produce      json
// @Param        page  query int false "Page number"
// @Param        limit query int false "Items per page"
// @Success      200 {array} models.Appointment
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /appointments [get]
func GetAppointments(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	var appointments []models.Appointment
	if err := config.DB.Preload("Patient").
		Limit(limit).
		Offset(offset).
		Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointments)
}

// GetAppointmentsForPatient godoc
// @Summary      Get all appointments for a specific patient
// @Description  Retrieve all appointments for a specific patient by patient ID
// @Tags         appointments
// @Accept       json
// @Produce      json
// @Param        id path int true "Patient ID"
// @Success      200 {array} models.Appointment
// @Failure      404 {object} map[string]string "Patient Not Found"
// @Failure      500 {object} map[string]string "Internal Server Error"
// @Router       /patients/{id}/appointments [get]
func GetAppointmentsForPatient(c *gin.Context) {
	patientID := c.Param("id")
	var appointments []models.Appointment

	if err := config.DB.Where("patient_id = ?", patientID).Preload("Patient").Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}
