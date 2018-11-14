// Code generated by go-swagger; DO NOT EDIT.

package transfer

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

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// NewPrepareParams creates a new PrepareParams object
// with the default values initialized.
func NewPrepareParams() *PrepareParams {
	var ()
	return &PrepareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPrepareParamsWithTimeout creates a new PrepareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPrepareParamsWithTimeout(timeout time.Duration) *PrepareParams {
	var ()
	return &PrepareParams{

		timeout: timeout,
	}
}

// NewPrepareParamsWithContext creates a new PrepareParams object
// with the default values initialized, and the ability to set a context for a request
func NewPrepareParamsWithContext(ctx context.Context) *PrepareParams {
	var ()
	return &PrepareParams{

		Context: ctx,
	}
}

// NewPrepareParamsWithHTTPClient creates a new PrepareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPrepareParamsWithHTTPClient(client *http.Client) *PrepareParams {
	var ()
	return &PrepareParams{
		HTTPClient: client,
	}
}

/*PrepareParams contains all the parameters to send to the API endpoint
for the prepare operation typically these are written to a http.Request
*/
type PrepareParams struct {

	/*ID4N
	  The ID4N to prepare for transfer

	*/
	ID4N string
	/*Request
	  Transfer preparation status

	*/
	Request *api_models.TransferSendInfo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the prepare params
func (o *PrepareParams) WithTimeout(timeout time.Duration) *PrepareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the prepare params
func (o *PrepareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the prepare params
func (o *PrepareParams) WithContext(ctx context.Context) *PrepareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the prepare params
func (o *PrepareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the prepare params
func (o *PrepareParams) WithHTTPClient(client *http.Client) *PrepareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the prepare params
func (o *PrepareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID4N adds the id4n to the prepare params
func (o *PrepareParams) WithID4N(id4n string) *PrepareParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the prepare params
func (o *PrepareParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WithRequest adds the request to the prepare params
func (o *PrepareParams) WithRequest(request *api_models.TransferSendInfo) *PrepareParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the prepare params
func (o *PrepareParams) SetRequest(request *api_models.TransferSendInfo) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PrepareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id4n
	if err := r.SetPathParam("id4n", o.ID4N); err != nil {
		return err
	}

	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
