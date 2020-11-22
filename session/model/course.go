package model

// CourseView CourseView
//
// All details about the Airplane (DTO)
//
// swagger:model CourseView
type CourseView struct {

	// difficulty level
	DifficultyLevel int32 `json:"difficultyLevel,omitempty"`

	// end date
	EndDate string `json:"endDate,omitempty"`

	// enrollments
	Enrollments []*EnrollmentView `json:"enrollments"`

	// id
	ID int64 `json:"id,omitempty"`

	// seats
	Seats int32 `json:"seats,omitempty"`

	// start date
	StartDate string `json:"startDate,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}
