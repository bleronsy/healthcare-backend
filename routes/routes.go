package routes

import (
	"healthcare-app/healthcare-backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/patients", controllers.GetPatients)
	router.GET("/patients/:id", controllers.GetPatientByID)
	router.POST("/patients", controllers.CreatePatient)

	router.GET("/appointments", controllers.GetAppointments)
	router.GET("/appointments/patient/:id", controllers.GetAppointmentsForPatient)
	router.POST("/patients/:id/appointments", controllers.CreateAppointment)
}
