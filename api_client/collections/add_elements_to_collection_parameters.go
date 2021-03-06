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

// NewAddElementsToCollectionParams creates a new AddElementsToCollectionParams object
// with the default values initialized.
func NewAddElementsToCollectionParams() *AddElementsToCollectionParams {
	var ()
	return &AddElementsToCollectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddElementsToCollectionParamsWithTimeout creates a new AddElementsToCollectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddElementsToCollectionParamsWithTimeout(timeout time.Duration) *AddElementsToCollectionParams {
	var ()
	return &AddElementsToCollectionParams{

		timeout: timeout,
	}
}

// NewAddElementsToCollectionParamsWithContext creates a new AddElementsToCollectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddElementsToCollectionParamsWithContext(ctx context.Context) *AddElementsToCollectionParams {
	var ()
	return &AddElementsToCollectionParams{

		Context: ctx,
	}
}

// NewAddElementsToCollectionParamsWithHTTPClient creates a new AddElementsToCollectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddElementsToCollectionParamsWithHTTPClient(client *http.Client) *AddElementsToCollectionParams {
	var ()
	return &AddElementsToCollectionParams{
		HTTPClient: client,
	}
}

/*AddElementsToCollectionParams contains all the parameters to send to the API endpoint
for the add elements to collection operation typically these are written to a http.Request
*/
type AddElementsToCollectionParams struct {

	/*ID4N
	  id4n

	*/
	ID4N string
	/*ListOfGuids
	  listOfGuids

	*/
	ListOfGuids *api_models.ListOfId4ns

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add elements to collection params
func (o *AddElementsToCollectionParams) WithTimeout(timeout time.Duration) *AddElementsToCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add elements to collection params
func (o *AddElementsToCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add elements to collection params
func (o *AddElementsToCollectionParams) WithContext(ctx context.Context) *AddElementsToCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add elements to collection params
func (o *AddElementsToCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add elements to collection params
func (o *AddElementsToCollectionParams) WithHTTPClient(client *http.Client) *AddElementsToCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add elements to collection params
func (o *AddElementsToCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID4N adds the id4n to the add elements to collection params
func (o *AddElementsToCollectionParams) WithID4N(id4n string) *AddElementsToCollectionParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the add elements to collection params
func (o *AddElementsToCollectionParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WithListOfGuids adds the listOfGuids to the add elements to collection params
func (o *AddElementsToCollectionParams) WithListOfGuids(listOfGuids *api_models.ListOfId4ns) *AddElementsToCollectionParams {
	o.SetListOfGuids(listOfGuids)
	return o
}

// SetListOfGuids adds the listOfGuids to the add elements to collection params
func (o *AddElementsToCollectionParams) SetListOfGuids(listOfGuids *api_models.ListOfId4ns) {
	o.ListOfGuids = listOfGuids
}

// WriteToRequest writes these params to a swagger request
func (o *AddElementsToCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id4n
	if err := r.SetPathParam("id4n", o.ID4N); err != nil {
		return err
	}

	if o.ListOfGuids != nil {
		if err := r.SetBodyParam(o.ListOfGuids); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
