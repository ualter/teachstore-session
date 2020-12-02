package model

import (
	"time"

	api_models "github.com/ualter/teachstore-session/gen/models"
)

// Session defines the structure for an API of Session Class
// swagger:model
type Session struct {
	// the id for the session
	//
	// required: true
	ID int64 `json:"id"`

	// the name of the session (usually related with the Course)
	//
	// required: true
	// max length: 255
	Name string `json:"name"`

	// the date of the session
	//
	// required: true
	// max length: 255
	Date time.Time `json:"date"`

	// Enrollments

	//Enrollments []*api_models.EnrollmentView `json:"enrollments"`
	Enrollment *api_models.EnrollmentView `json:"enrollment"`

	// The student attends to this session
	//
	// required: true
	Attendance bool `json:"attendance"`
}
