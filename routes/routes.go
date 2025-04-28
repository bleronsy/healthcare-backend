package routes

import (
	"healthcare-app/healthcare-backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	patientRoutes := router.Group("/patients")
	{
		patientRoutes.POST("/", handlers.CreatePatient)
		patientRoutes.GET("/", handlers.GetAllPatients)
		patientRoutes.GET("/:id", handlers.GetPatientByID)
	}

	appointmentRoutes := router.Group("/appointments")
	{
		appointmentRoutes.POST("/", handlers.CreateAppointment)
		appointmentRoutes.GET("/", handlers.GetAllAppointments)
		appointmentRoutes.GET("/patient/:patient_id", handlers.GetAppointmentsByPatientID)
	}
}
