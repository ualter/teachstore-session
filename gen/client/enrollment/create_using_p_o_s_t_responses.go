// Code generated by go-swagger; DO NOT EDIT.

package enrollment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ualter/teachstore-session/gen/models"
)

// CreateUsingPOSTReader is a Reader for the CreateUsingPOST structure.
type CreateUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUsingPOSTCreated creates a CreateUsingPOSTCreated with default headers values
func NewCreateUsingPOSTCreated() *CreateUsingPOSTCreated {
	return &CreateUsingPOSTCreated{}
}

/*CreateUsingPOSTCreated handles this case with default header values.

Resource created
*/
type CreateUsingPOSTCreated struct {
	Payload *models.Enrollment
}

func (o *CreateUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /api/enrollments][%d] createUsingPOSTCreated  %+v", 201, o.Payload)
}

func (o *CreateUsingPOSTCreated) GetPayload() *models.Enrollment {
	return o.Payload
}

func (o *CreateUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Enrollment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUsingPOSTBadRequest creates a CreateUsingPOSTBadRequest with default headers values
func NewCreateUsingPOSTBadRequest() *CreateUsingPOSTBadRequest {
	return &CreateUsingPOSTBadRequest{}
}

/*CreateUsingPOSTBadRequest handles this case with default header values.

Bad Request
*/
type CreateUsingPOSTBadRequest struct {
}

func (o *CreateUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/enrollments][%d] createUsingPOSTBadRequest ", 400)
}

func (o *CreateUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOSTUnauthorized creates a CreateUsingPOSTUnauthorized with default headers values
func NewCreateUsingPOSTUnauthorized() *CreateUsingPOSTUnauthorized {
	return &CreateUsingPOSTUnauthorized{}
}

/*CreateUsingPOSTUnauthorized handles this case with default header values.

Not authorized
*/
type CreateUsingPOSTUnauthorized struct {
}

func (o *CreateUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/enrollments][%d] createUsingPOSTUnauthorized ", 401)
}

func (o *CreateUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOSTForbidden creates a CreateUsingPOSTForbidden with default headers values
func NewCreateUsingPOSTForbidden() *CreateUsingPOSTForbidden {
	return &CreateUsingPOSTForbidden{}
}

/*CreateUsingPOSTForbidden handles this case with default header values.

Access forbidden
*/
type CreateUsingPOSTForbidden struct {
}

func (o *CreateUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /api/enrollments][%d] createUsingPOSTForbidden ", 403)
}

func (o *CreateUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOSTNotFound creates a CreateUsingPOSTNotFound with default headers values
func NewCreateUsingPOSTNotFound() *CreateUsingPOSTNotFound {
	return &CreateUsingPOSTNotFound{}
}

/*CreateUsingPOSTNotFound handles this case with default header values.

Not found
*/
type CreateUsingPOSTNotFound struct {
}

func (o *CreateUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /api/enrollments][%d] createUsingPOSTNotFound ", 404)
}

func (o *CreateUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOSTInternalServerError creates a CreateUsingPOSTInternalServerError with default headers values
func NewCreateUsingPOSTInternalServerError() *CreateUsingPOSTInternalServerError {
	return &CreateUsingPOSTInternalServerError{}
}

/*CreateUsingPOSTInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateUsingPOSTInternalServerError struct {
}

func (o *CreateUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/enrollments][%d] createUsingPOSTInternalServerError ", 500)
}

func (o *CreateUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}