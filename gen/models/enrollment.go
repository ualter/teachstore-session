// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Enrollment Enrollment
//
// All details about the Enroll.
//
// swagger:model Enrollment
type Enrollment struct {

	// course
	Course *CourseView `json:"course,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// register date
	RegisterDate string `json:"registerDate,omitempty"`

	// student
	Student *StudentView `json:"student,omitempty"`
}

// Validate validates this enrollment
func (m *Enrollment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCourse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStudent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Enrollment) validateCourse(formats strfmt.Registry) error {

	if swag.IsZero(m.Course) { // not required
		return nil
	}

	if m.Course != nil {
		if err := m.Course.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("course")
			}
			return err
		}
	}

	return nil
}

func (m *Enrollment) validateStudent(formats strfmt.Registry) error {

	if swag.IsZero(m.Student) { // not required
		return nil
	}

	if m.Student != nil {
		if err := m.Student.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("student")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Enrollment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Enrollment) UnmarshalBinary(b []byte) error {
	var res Enrollment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
