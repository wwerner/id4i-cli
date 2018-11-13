// Code generated by go-swagger; DO NOT EDIT.

package collections

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

// NewUpdateCollectionParams creates a new UpdateCollectionParams object
// with the default values initialized.
func NewUpdateCollectionParams() *UpdateCollectionParams {
	var ()
	return &UpdateCollectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateCollectionParamsWithTimeout creates a new UpdateCollectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateCollectionParamsWithTimeout(timeout time.Duration) *UpdateCollectionParams {
	var ()
	return &UpdateCollectionParams{

		timeout: timeout,
	}
}

// NewUpdateCollectionParamsWithContext creates a new UpdateCollectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateCollectionParamsWithContext(ctx context.Context) *UpdateCollectionParams {
	var ()
	return &UpdateCollectionParams{

		Context: ctx,
	}
}

// NewUpdateCollectionParamsWithHTTPClient creates a new UpdateCollectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateCollectionParamsWithHTTPClient(client *http.Client) *UpdateCollectionParams {
	var ()
	return &UpdateCollectionParams{
		HTTPClient: client,
	}
}

/*UpdateCollectionParams contains all the parameters to send to the API endpoint
for the update collection operation typically these are written to a http.Request
*/
type UpdateCollectionParams struct {

	/*ID4N
	  id4n

	*/
	ID4N string
	/*Request
	  request

	*/
	Request *api_models.GUIDCollection

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update collection params
func (o *UpdateCollectionParams) WithTimeout(timeout time.Duration) *UpdateCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update collection params
func (o *UpdateCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update collection params
func (o *UpdateCollectionParams) WithContext(ctx context.Context) *UpdateCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update collection params
func (o *UpdateCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update collection params
func (o *UpdateCollectionParams) WithHTTPClient(client *http.Client) *UpdateCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update collection params
func (o *UpdateCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID4N adds the id4n to the update collection params
func (o *UpdateCollectionParams) WithID4N(id4n string) *UpdateCollectionParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the update collection params
func (o *UpdateCollectionParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WithRequest adds the request to the update collection params
func (o *UpdateCollectionParams) WithRequest(request *api_models.GUIDCollection) *UpdateCollectionParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update collection params
func (o *UpdateCollectionParams) SetRequest(request *api_models.GUIDCollection) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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