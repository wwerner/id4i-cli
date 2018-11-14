// Code generated by go-swagger; DO NOT EDIT.

package public_services

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

// NewGoParams creates a new GoParams object
// with the default values initialized.
func NewGoParams() *GoParams {
	var ()
	return &GoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGoParamsWithTimeout creates a new GoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGoParamsWithTimeout(timeout time.Duration) *GoParams {
	var ()
	return &GoParams{

		timeout: timeout,
	}
}

// NewGoParamsWithContext creates a new GoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGoParamsWithContext(ctx context.Context) *GoParams {
	var ()
	return &GoParams{

		Context: ctx,
	}
}

// NewGoParamsWithHTTPClient creates a new GoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGoParamsWithHTTPClient(client *http.Client) *GoParams {
	var ()
	return &GoParams{
		HTTPClient: client,
	}
}

/*GoParams contains all the parameters to send to the API endpoint
for the go operation typically these are written to a http.Request
*/
type GoParams struct {

	/*GUID
	  guid

	*/
	GUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the go params
func (o *GoParams) WithTimeout(timeout time.Duration) *GoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the go params
func (o *GoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the go params
func (o *GoParams) WithContext(ctx context.Context) *GoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the go params
func (o *GoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the go params
func (o *GoParams) WithHTTPClient(client *http.Client) *GoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the go params
func (o *GoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGUID adds the guid to the go params
func (o *GoParams) WithGUID(guid string) *GoParams {
	o.SetGUID(guid)
	return o
}

// SetGUID adds the guid to the go params
func (o *GoParams) SetGUID(guid string) {
	o.GUID = guid
}

// WriteToRequest writes these params to a swagger request
func (o *GoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param guid
	if err := r.SetPathParam("guid", o.GUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
