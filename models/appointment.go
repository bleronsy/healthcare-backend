package models

// Appointment represents an appointment for a patient
// @Description Appointment data, including date, time, and reason for the appointment.
// @Model
type Appointment struct {
	// ID of the appointment
	// @example 1
	ID uint `json:"id"`

	// Date of the appointment
	// @example 2025-05-01
	Date string `json:"date" binding:"required" example:"2025-05-01"`

	// Time of the appointment
	// @example 10:00 AM
	Time string `json:"time" binding:"required" example:"10:00 AM"`

	// Reason for the appointment
	// @example Routine Checkup
	Reason string `json:"reason" example:"Routine Checkup"`

	// Notes for the appointment
	// @example "Patient reports feeling unwell."
	Notes string `json:"notes" example:"Patient reports feeling unwell."`

	// Foreign key to the patient this appointment belongs to
	PatientID uint `json:"patient_id"`
}
