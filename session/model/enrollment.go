package model

// EnrollmentView EnrollmentView
//
// swagger:model EnrollmentView
type EnrollmentView struct {

	// course
	Course *CourseView `json:"course,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// register date
	RegisterDate string `json:"registerDate,omitempty"`

	// student
	Student *StudentView `json:"student,omitempty"`
}
