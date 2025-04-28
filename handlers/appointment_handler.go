package handlers

import (
	"net/http"

	"healthcare-app/healthcare-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAppointment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate patient exists
	var patient models.Patient
	if err := db.First(&patient, appointment.PatientID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid patient ID"})
		return
	}

	db.Create(&appointment)
	c.JSON(http.StatusOK, appointment)
}

func GetAllAppointments(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var appointments []models.Appointment
	db.Find(&appointments)
	c.JSON(http.StatusOK, appointments)
}

func GetAppointmentsByPatientID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	patientID := c.Param("patient_id")
	var appointments []models.Appointment
	db.Where("patient_id = ?", patientID).Find(&appointments)
	c.JSON(http.StatusOK, appointments)
}
