// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// NewAddPartnerOrganizationParams creates a new AddPartnerOrganizationParams object
// with the default values initialized.
func NewAddPartnerOrganizationParams() *AddPartnerOrganizationParams {
	var ()
	return &AddPartnerOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddPartnerOrganizationParamsWithTimeout creates a new AddPartnerOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddPartnerOrganizationParamsWithTimeout(timeout time.Duration) *AddPartnerOrganizationParams {
	var ()
	return &AddPartnerOrganizationParams{

		timeout: timeout,
	}
}

// NewAddPartnerOrganizationParamsWithContext creates a new AddPartnerOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddPartnerOrganizationParamsWithContext(ctx context.Context) *AddPartnerOrganizationParams {
	var ()
	return &AddPartnerOrganizationParams{

		Context: ctx,
	}
}

// NewAddPartnerOrganizationParamsWithHTTPClient creates a new AddPartnerOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddPartnerOrganizationParamsWithHTTPClient(client *http.Client) *AddPartnerOrganizationParams {
	var ()
	return &AddPartnerOrganizationParams{
		HTTPClient: client,
	}
}

/*AddPartnerOrganizationParams contains all the parameters to send to the API endpoint
for the add partner organization operation typically these are written to a http.Request
*/
type AddPartnerOrganizationParams struct {

	/*OrganizationID
	  The namespace of the organization

	*/
	OrganizationID string
	/*Request
	  request

	*/
	Request *api_models.AddPartnerRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add partner organization params
func (o *AddPartnerOrganizationParams) WithTimeout(timeout time.Duration) *AddPartnerOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add partner organization params
func (o *AddPartnerOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add partner organization params
func (o *AddPartnerOrganizationParams) WithContext(ctx context.Context) *AddPartnerOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add partner organization params
func (o *AddPartnerOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add partner organization params
func (o *AddPartnerOrganizationParams) WithHTTPClient(client *http.Client) *AddPartnerOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add partner organization params
func (o *AddPartnerOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the add partner organization params
func (o *AddPartnerOrganizationParams) WithOrganizationID(organizationID string) *AddPartnerOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the add partner organization params
func (o *AddPartnerOrganizationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithRequest adds the request to the add partner organization params
func (o *AddPartnerOrganizationParams) WithRequest(request *api_models.AddPartnerRequest) *AddPartnerOrganizationParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the add partner organization params
func (o *AddPartnerOrganizationParams) SetRequest(request *api_models.AddPartnerRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *AddPartnerOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", o.OrganizationID); err != nil {
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