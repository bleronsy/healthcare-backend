package models

import "gorm.io/gorm"

type Appointment struct {
	gorm.Model
	PatientID uint   `json:"patient_id"`
	Date      string `json:"date"`
	Reason    string `json:"reason"`
}
