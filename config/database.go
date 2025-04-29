package config

import (
	"log"

	"healthcare-app/healthcare-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("healthcare.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	err = DB.AutoMigrate(&models.Patient{}, &models.Appointment{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
