package model

// StudentView StudentView
//
// All details about the Airplane (DTO)
//
// swagger:model StudentView
type StudentView struct {

	// birth date
	BirthDate string `json:"birthDate,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}
