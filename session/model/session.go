package model

import "time"

// Session defines the structure for an API of Session Class
// swagger:model
type Session struct {
	// the id for the session
	//
	// required: true
	ID int32 `json:"id"`

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
	Enrollments []*EnrollmentView `json:"enrollments"`
	//Addresses []Address `json:"addresses,omitempty"`

	// The student attends to this session
	//
	// required: true
	Attendance bool `json:"attendance"`
}
