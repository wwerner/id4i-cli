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

// NewRemovePartnerOrganizationParams creates a new RemovePartnerOrganizationParams object
// with the default values initialized.
func NewRemovePartnerOrganizationParams() *RemovePartnerOrganizationParams {
	var ()
	return &RemovePartnerOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemovePartnerOrganizationParamsWithTimeout creates a new RemovePartnerOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemovePartnerOrganizationParamsWithTimeout(timeout time.Duration) *RemovePartnerOrganizationParams {
	var ()
	return &RemovePartnerOrganizationParams{

		timeout: timeout,
	}
}

// NewRemovePartnerOrganizationParamsWithContext creates a new RemovePartnerOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemovePartnerOrganizationParamsWithContext(ctx context.Context) *RemovePartnerOrganizationParams {
	var ()
	return &RemovePartnerOrganizationParams{

		Context: ctx,
	}
}

// NewRemovePartnerOrganizationParamsWithHTTPClient creates a new RemovePartnerOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemovePartnerOrganizationParamsWithHTTPClient(client *http.Client) *RemovePartnerOrganizationParams {
	var ()
	return &RemovePartnerOrganizationParams{
		HTTPClient: client,
	}
}

/*RemovePartnerOrganizationParams contains all the parameters to send to the API endpoint
for the remove partner organization operation typically these are written to a http.Request
*/
type RemovePartnerOrganizationParams struct {

	/*OrganizationID
	  The namespace of the organization

	*/
	OrganizationID string
	/*Request
	  request

	*/
	Request *api_models.RemovePartnerRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove partner organization params
func (o *RemovePartnerOrganizationParams) WithTimeout(timeout time.Duration) *RemovePartnerOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove partner organization params
func (o *RemovePartnerOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove partner organization params
func (o *RemovePartnerOrganizationParams) WithContext(ctx context.Context) *RemovePartnerOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove partner organization params
func (o *RemovePartnerOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove partner organization params
func (o *RemovePartnerOrganizationParams) WithHTTPClient(client *http.Client) *RemovePartnerOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove partner organization params
func (o *RemovePartnerOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the remove partner organization params
func (o *RemovePartnerOrganizationParams) WithOrganizationID(organizationID string) *RemovePartnerOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the remove partner organization params
func (o *RemovePartnerOrganizationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithRequest adds the request to the remove partner organization params
func (o *RemovePartnerOrganizationParams) WithRequest(request *api_models.RemovePartnerRequest) *RemovePartnerOrganizationParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the remove partner organization params
func (o *RemovePartnerOrganizationParams) SetRequest(request *api_models.RemovePartnerRequest) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *RemovePartnerOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
