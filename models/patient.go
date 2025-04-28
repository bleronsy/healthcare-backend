package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Appointments []Appointment `gorm:"foreignKey:PatientID"`
}
