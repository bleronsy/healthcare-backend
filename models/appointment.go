package models

import "gorm.io/gorm"

// Appointment represents an appointment for a patient
// @Description Appointment data with date, time, notes, and the associated patient.
// @Model
type Appointment struct {
	gorm.Model
	// Date of the appointment
	// @example "2025-04-29"
	Date string `json:"date" binding:"required" swagger:"desc(Appointment date)"`
	// Time of the appointment
	// @example "09:00 AM"
	Time string `json:"time" binding:"required" swagger:"desc(Appointment time)"`
	// Additional notes for the appointment
	// @example "Patient has a checkup."
	Notes string `json:"notes" swagger:"desc(Additional appointment notes)"`
	// PatientID references the Patient the appointment is for
	PatientID uint `json:"patient_id" swagger:"desc(Patient ID the appointment is associated with)"`
	// The Patient associated with this appointment
	Patient Patient `json:"patient" gorm:"foreignKey:PatientID" swagger:"desc(Patient details for the appointment)"`
}
