// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewReadFromMicrostorageParams creates a new ReadFromMicrostorageParams object
// with the default values initialized.
func NewReadFromMicrostorageParams() *ReadFromMicrostorageParams {
	var ()
	return &ReadFromMicrostorageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadFromMicrostorageParamsWithTimeout creates a new ReadFromMicrostorageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadFromMicrostorageParamsWithTimeout(timeout time.Duration) *ReadFromMicrostorageParams {
	var ()
	return &ReadFromMicrostorageParams{

		timeout: timeout,
	}
}

// NewReadFromMicrostorageParamsWithContext creates a new ReadFromMicrostorageParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadFromMicrostorageParamsWithContext(ctx context.Context) *ReadFromMicrostorageParams {
	var ()
	return &ReadFromMicrostorageParams{

		Context: ctx,
	}
}

// NewReadFromMicrostorageParamsWithHTTPClient creates a new ReadFromMicrostorageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadFromMicrostorageParamsWithHTTPClient(client *http.Client) *ReadFromMicrostorageParams {
	var ()
	return &ReadFromMicrostorageParams{
		HTTPClient: client,
	}
}

/*ReadFromMicrostorageParams contains all the parameters to send to the API endpoint
for the read from microstorage operation typically these are written to a http.Request
*/
type ReadFromMicrostorageParams struct {

	/*ID4N
	  id4n

	*/
	ID4N string
	/*Organization
	  organization

	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read from microstorage params
func (o *ReadFromMicrostorageParams) WithTimeout(timeout time.Duration) *ReadFromMicrostorageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read from microstorage params
func (o *ReadFromMicrostorageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read from microstorage params
func (o *ReadFromMicrostorageParams) WithContext(ctx context.Context) *ReadFromMicrostorageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read from microstorage params
func (o *ReadFromMicrostorageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read from microstorage params
func (o *ReadFromMicrostorageParams) WithHTTPClient(client *http.Client) *ReadFromMicrostorageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read from microstorage params
func (o *ReadFromMicrostorageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID4N adds the id4n to the read from microstorage params
func (o *ReadFromMicrostorageParams) WithID4N(id4n string) *ReadFromMicrostorageParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the read from microstorage params
func (o *ReadFromMicrostorageParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WithOrganization adds the organization to the read from microstorage params
func (o *ReadFromMicrostorageParams) WithOrganization(organization string) *ReadFromMicrostorageParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the read from microstorage params
func (o *ReadFromMicrostorageParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *ReadFromMicrostorageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id4n
	if err := r.SetPathParam("id4n", o.ID4N); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
