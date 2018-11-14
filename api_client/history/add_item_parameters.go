// Code generated by go-swagger; DO NOT EDIT.

package history

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

// NewAddItemParams creates a new AddItemParams object
// with the default values initialized.
func NewAddItemParams() *AddItemParams {
	var ()
	return &AddItemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddItemParamsWithTimeout creates a new AddItemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddItemParamsWithTimeout(timeout time.Duration) *AddItemParams {
	var ()
	return &AddItemParams{

		timeout: timeout,
	}
}

// NewAddItemParamsWithContext creates a new AddItemParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddItemParamsWithContext(ctx context.Context) *AddItemParams {
	var ()
	return &AddItemParams{

		Context: ctx,
	}
}

// NewAddItemParamsWithHTTPClient creates a new AddItemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddItemParamsWithHTTPClient(client *http.Client) *AddItemParams {
	var ()
	return &AddItemParams{
		HTTPClient: client,
	}
}

/*AddItemParams contains all the parameters to send to the API endpoint
for the add item operation typically these are written to a http.Request
*/
type AddItemParams struct {

	/*HistoryItem
	  The history item to publish

	*/
	HistoryItem *api_models.HistoryItem
	/*ID4N
	  GUID to retrieve the history for

	*/
	ID4N string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add item params
func (o *AddItemParams) WithTimeout(timeout time.Duration) *AddItemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add item params
func (o *AddItemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add item params
func (o *AddItemParams) WithContext(ctx context.Context) *AddItemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add item params
func (o *AddItemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add item params
func (o *AddItemParams) WithHTTPClient(client *http.Client) *AddItemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add item params
func (o *AddItemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHistoryItem adds the historyItem to the add item params
func (o *AddItemParams) WithHistoryItem(historyItem *api_models.HistoryItem) *AddItemParams {
	o.SetHistoryItem(historyItem)
	return o
}

// SetHistoryItem adds the historyItem to the add item params
func (o *AddItemParams) SetHistoryItem(historyItem *api_models.HistoryItem) {
	o.HistoryItem = historyItem
}

// WithID4N adds the id4n to the add item params
func (o *AddItemParams) WithID4N(id4n string) *AddItemParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the add item params
func (o *AddItemParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WriteToRequest writes these params to a swagger request
func (o *AddItemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.HistoryItem != nil {
		if err := r.SetBodyParam(o.HistoryItem); err != nil {
			return err
		}
	}

	// path param id4n
	if err := r.SetPathParam("id4n", o.ID4N); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
