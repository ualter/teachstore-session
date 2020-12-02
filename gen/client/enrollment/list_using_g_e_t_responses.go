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

// ListUsingGETReader is a Reader for the ListUsingGET structure.
type ListUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListUsingGETBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListUsingGETOK creates a ListUsingGETOK with default headers values
func NewListUsingGETOK() *ListUsingGETOK {
	return &ListUsingGETOK{}
}

/*ListUsingGETOK handles this case with default header values.

Request succeeded
*/
type ListUsingGETOK struct {
	Payload []*models.EnrollmentView
}

func (o *ListUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api/enrollments][%d] listUsingGETOK  %+v", 200, o.Payload)
}

func (o *ListUsingGETOK) GetPayload() []*models.EnrollmentView {
	return o.Payload
}

func (o *ListUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUsingGETBadRequest creates a ListUsingGETBadRequest with default headers values
func NewListUsingGETBadRequest() *ListUsingGETBadRequest {
	return &ListUsingGETBadRequest{}
}

/*ListUsingGETBadRequest handles this case with default header values.

Bad Request
*/
type ListUsingGETBadRequest struct {
}

func (o *ListUsingGETBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/enrollments][%d] listUsingGETBadRequest ", 400)
}

func (o *ListUsingGETBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUsingGETUnauthorized creates a ListUsingGETUnauthorized with default headers values
func NewListUsingGETUnauthorized() *ListUsingGETUnauthorized {
	return &ListUsingGETUnauthorized{}
}

/*ListUsingGETUnauthorized handles this case with default header values.

Not authorized
*/
type ListUsingGETUnauthorized struct {
}

func (o *ListUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/enrollments][%d] listUsingGETUnauthorized ", 401)
}

func (o *ListUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUsingGETForbidden creates a ListUsingGETForbidden with default headers values
func NewListUsingGETForbidden() *ListUsingGETForbidden {
	return &ListUsingGETForbidden{}
}

/*ListUsingGETForbidden handles this case with default header values.

Access forbidden
*/
type ListUsingGETForbidden struct {
}

func (o *ListUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/enrollments][%d] listUsingGETForbidden ", 403)
}

func (o *ListUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUsingGETNotFound creates a ListUsingGETNotFound with default headers values
func NewListUsingGETNotFound() *ListUsingGETNotFound {
	return &ListUsingGETNotFound{}
}

/*ListUsingGETNotFound handles this case with default header values.

Not found
*/
type ListUsingGETNotFound struct {
}

func (o *ListUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/enrollments][%d] listUsingGETNotFound ", 404)
}

func (o *ListUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListUsingGETInternalServerError creates a ListUsingGETInternalServerError with default headers values
func NewListUsingGETInternalServerError() *ListUsingGETInternalServerError {
	return &ListUsingGETInternalServerError{}
}

/*ListUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListUsingGETInternalServerError struct {
}

func (o *ListUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/enrollments][%d] listUsingGETInternalServerError ", 500)
}

func (o *ListUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
