// Code generated by go-swagger; DO NOT EDIT.

package enrollment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ualter/teachstore-session/gen/models"
)

// NewUpdateUsingPUTParams creates a new UpdateUsingPUTParams object
// with the default values initialized.
func NewUpdateUsingPUTParams() *UpdateUsingPUTParams {
	var ()
	return &UpdateUsingPUTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUsingPUTParamsWithTimeout creates a new UpdateUsingPUTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateUsingPUTParams {
	var ()
	return &UpdateUsingPUTParams{

		timeout: timeout,
	}
}

// NewUpdateUsingPUTParamsWithContext creates a new UpdateUsingPUTParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUsingPUTParamsWithContext(ctx context.Context) *UpdateUsingPUTParams {
	var ()
	return &UpdateUsingPUTParams{

		Context: ctx,
	}
}

// NewUpdateUsingPUTParamsWithHTTPClient creates a new UpdateUsingPUTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateUsingPUTParams {
	var ()
	return &UpdateUsingPUTParams{
		HTTPClient: client,
	}
}

/*UpdateUsingPUTParams contains all the parameters to send to the API endpoint
for the update using p u t operation typically these are written to a http.Request
*/
type UpdateUsingPUTParams struct {

	/*Enroll
	  enroll

	*/
	Enroll *models.Enrollment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update using p u t params
func (o *UpdateUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update using p u t params
func (o *UpdateUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update using p u t params
func (o *UpdateUsingPUTParams) WithContext(ctx context.Context) *UpdateUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update using p u t params
func (o *UpdateUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update using p u t params
func (o *UpdateUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update using p u t params
func (o *UpdateUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnroll adds the enroll to the update using p u t params
func (o *UpdateUsingPUTParams) WithEnroll(enroll *models.Enrollment) *UpdateUsingPUTParams {
	o.SetEnroll(enroll)
	return o
}

// SetEnroll adds the enroll to the update using p u t params
func (o *UpdateUsingPUTParams) SetEnroll(enroll *models.Enrollment) {
	o.Enroll = enroll
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Enroll != nil {
		if err := r.SetBodyParam(o.Enroll); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
