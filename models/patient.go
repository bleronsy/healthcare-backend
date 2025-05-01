package models

// Patient represents the patient data
// @Description Patient data, including name, email, and appointments.
// @Model
type Patient struct {
	// ID of the patient (inherited from gorm.Model)
	// @example 1
	ID uint `json:"id"`

	// CreatedAt timestamp (inherited from gorm.Model)
	// @example "2025-05-01T12:00:00Z"
	CreatedAt string `json:"created_at"`

	// UpdatedAt timestamp (inherited from gorm.Model)
	// @example "2025-05-01T12:00:00Z"
	UpdatedAt string `json:"updated_at"`

	// DeletedAt timestamp (nullable, inherited from gorm.Model)
	// @example "2025-05-01T12:00:00Z"
	DeletedAt *string `json:"deleted_at,omitempty"`

	// Name of the patient
	// @example John Doe
	// @description Patient's full name
	// @required
	Name string `json:"name" binding:"required" example:"John Doe"`

	// Email of the patient
	// @example john.doe@example.com
	// @description Patient's email address
	// @required
	Email string `json:"email" binding:"required,email" example:"john.doe@example.com"`

	// Appointments associated with the patient
	// @description Patient's appointments (if any)
	// @swagger:property
	// @type array
	// @items Appointment
	// @model
	Appointments []Appointment `json:"appointments,omitempty" gorm:"foreignKey:PatientID"`
}
