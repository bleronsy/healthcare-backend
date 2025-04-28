package handlers

import (
	"net/http"

	"healthcare-app/healthcare-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&patient)
	c.JSON(http.StatusOK, patient)
}

func GetAllPatients(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var patients []models.Patient
	db.Find(&patients)
	c.JSON(http.StatusOK, patients)
}

func GetPatientByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var patient models.Patient
	if err := db.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}
