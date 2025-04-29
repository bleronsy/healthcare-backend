package controllers

import (
	"net/http"
	"strconv"

	"healthcare-app/healthcare-backend/config"
	"healthcare-app/healthcare-backend/models"

	"github.com/gin-gonic/gin"
)

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

func GetAppointmentsForPatient(c *gin.Context) {
	patientID := c.Param("id")
	var appointments []models.Appointment

	if err := config.DB.Where("patient_id = ?", patientID).Preload("Patient").Find(&appointments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}
