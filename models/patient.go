// models/patient.go
package models

import "gorm.io/gorm"

// Patient represents the patient data
// @Description Patient data, including name, email, and appointments.
// @Model
type Patient struct {
	gorm.Model
	// Name of the patient
	// @example John Doe
	Name string `json:"name" binding:"required" swagger:"desc(Patient's full name)"`
	// Email of the patient
	// @example john.doe@example.com
	Email string `json:"email" binding:"required,email" swagger:"desc(Patient's email address)"`
	// Appointments associated with the patient
	Appointments []Appointment `json:"appointments,omitempty" gorm:"foreignKey:PatientID" swagger:"desc(Patient's appointments)"`
}
