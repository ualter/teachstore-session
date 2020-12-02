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
	"github.com/go-openapi/swag"
)

// NewFindByIDUsingGETParams creates a new FindByIDUsingGETParams object
// with the default values initialized.
func NewFindByIDUsingGETParams() *FindByIDUsingGETParams {
	var ()
	return &FindByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindByIDUsingGETParamsWithTimeout creates a new FindByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindByIDUsingGETParamsWithTimeout(timeout time.Duration) *FindByIDUsingGETParams {
	var ()
	return &FindByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewFindByIDUsingGETParamsWithContext creates a new FindByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindByIDUsingGETParamsWithContext(ctx context.Context) *FindByIDUsingGETParams {
	var ()
	return &FindByIDUsingGETParams{

		Context: ctx,
	}
}

// NewFindByIDUsingGETParamsWithHTTPClient creates a new FindByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindByIDUsingGETParamsWithHTTPClient(client *http.Client) *FindByIDUsingGETParams {
	var ()
	return &FindByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*FindByIDUsingGETParams contains all the parameters to send to the API endpoint
for the find by Id using g e t operation typically these are written to a http.Request
*/
type FindByIDUsingGETParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find by Id using g e t params
func (o *FindByIDUsingGETParams) WithTimeout(timeout time.Duration) *FindByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find by Id using g e t params
func (o *FindByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find by Id using g e t params
func (o *FindByIDUsingGETParams) WithContext(ctx context.Context) *FindByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find by Id using g e t params
func (o *FindByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find by Id using g e t params
func (o *FindByIDUsingGETParams) WithHTTPClient(client *http.Client) *FindByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find by Id using g e t params
func (o *FindByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find by Id using g e t params
func (o *FindByIDUsingGETParams) WithID(id int64) *FindByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find by Id using g e t params
func (o *FindByIDUsingGETParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
