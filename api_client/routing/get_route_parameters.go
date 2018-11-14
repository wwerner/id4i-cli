// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGETRouteParams creates a new GETRouteParams object
// with the default values initialized.
func NewGETRouteParams() *GETRouteParams {
	var (
		interpolateDefault   = bool(true)
		privateRoutesDefault = bool(true)
		publicRoutesDefault  = bool(true)
	)
	return &GETRouteParams{
		Interpolate:   &interpolateDefault,
		PrivateRoutes: &privateRoutesDefault,
		PublicRoutes:  &publicRoutesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGETRouteParamsWithTimeout creates a new GETRouteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGETRouteParamsWithTimeout(timeout time.Duration) *GETRouteParams {
	var (
		interpolateDefault   = bool(true)
		privateRoutesDefault = bool(true)
		publicRoutesDefault  = bool(true)
	)
	return &GETRouteParams{
		Interpolate:   &interpolateDefault,
		PrivateRoutes: &privateRoutesDefault,
		PublicRoutes:  &publicRoutesDefault,

		timeout: timeout,
	}
}

// NewGETRouteParamsWithContext creates a new GETRouteParams object
// with the default values initialized, and the ability to set a context for a request
func NewGETRouteParamsWithContext(ctx context.Context) *GETRouteParams {
	var (
		interpolateDefault   = bool(true)
		privateRoutesDefault = bool(true)
		publicRoutesDefault  = bool(true)
	)
	return &GETRouteParams{
		Interpolate:   &interpolateDefault,
		PrivateRoutes: &privateRoutesDefault,
		PublicRoutes:  &publicRoutesDefault,

		Context: ctx,
	}
}

// NewGETRouteParamsWithHTTPClient creates a new GETRouteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGETRouteParamsWithHTTPClient(client *http.Client) *GETRouteParams {
	var (
		interpolateDefault   = bool(true)
		privateRoutesDefault = bool(true)
		publicRoutesDefault  = bool(true)
	)
	return &GETRouteParams{
		Interpolate:   &interpolateDefault,
		PrivateRoutes: &privateRoutesDefault,
		PublicRoutes:  &publicRoutesDefault,
		HTTPClient:    client,
	}
}

/*GETRouteParams contains all the parameters to send to the API endpoint
for the get route operation typically these are written to a http.Request
*/
type GETRouteParams struct {

	/*ID4N
	  id4n

	*/
	ID4N string
	/*Interpolate
	  interpolate

	*/
	Interpolate *bool
	/*PrivateRoutes
	  privateRoutes

	*/
	PrivateRoutes *bool
	/*PublicRoutes
	  publicRoutes

	*/
	PublicRoutes *bool
	/*Type
	  The type of route you want to have

	*/
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get route params
func (o *GETRouteParams) WithTimeout(timeout time.Duration) *GETRouteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get route params
func (o *GETRouteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get route params
func (o *GETRouteParams) WithContext(ctx context.Context) *GETRouteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get route params
func (o *GETRouteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get route params
func (o *GETRouteParams) WithHTTPClient(client *http.Client) *GETRouteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get route params
func (o *GETRouteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID4N adds the id4n to the get route params
func (o *GETRouteParams) WithID4N(id4n string) *GETRouteParams {
	o.SetID4N(id4n)
	return o
}

// SetID4N adds the id4n to the get route params
func (o *GETRouteParams) SetID4N(id4n string) {
	o.ID4N = id4n
}

// WithInterpolate adds the interpolate to the get route params
func (o *GETRouteParams) WithInterpolate(interpolate *bool) *GETRouteParams {
	o.SetInterpolate(interpolate)
	return o
}

// SetInterpolate adds the interpolate to the get route params
func (o *GETRouteParams) SetInterpolate(interpolate *bool) {
	o.Interpolate = interpolate
}

// WithPrivateRoutes adds the privateRoutes to the get route params
func (o *GETRouteParams) WithPrivateRoutes(privateRoutes *bool) *GETRouteParams {
	o.SetPrivateRoutes(privateRoutes)
	return o
}

// SetPrivateRoutes adds the privateRoutes to the get route params
func (o *GETRouteParams) SetPrivateRoutes(privateRoutes *bool) {
	o.PrivateRoutes = privateRoutes
}

// WithPublicRoutes adds the publicRoutes to the get route params
func (o *GETRouteParams) WithPublicRoutes(publicRoutes *bool) *GETRouteParams {
	o.SetPublicRoutes(publicRoutes)
	return o
}

// SetPublicRoutes adds the publicRoutes to the get route params
func (o *GETRouteParams) SetPublicRoutes(publicRoutes *bool) {
	o.PublicRoutes = publicRoutes
}

// WithType adds the typeVar to the get route params
func (o *GETRouteParams) WithType(typeVar string) *GETRouteParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get route params
func (o *GETRouteParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GETRouteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id4n
	if err := r.SetPathParam("id4n", o.ID4N); err != nil {
		return err
	}

	if o.Interpolate != nil {

		// query param interpolate
		var qrInterpolate bool
		if o.Interpolate != nil {
			qrInterpolate = *o.Interpolate
		}
		qInterpolate := swag.FormatBool(qrInterpolate)
		if qInterpolate != "" {
			if err := r.SetQueryParam("interpolate", qInterpolate); err != nil {
				return err
			}
		}

	}

	if o.PrivateRoutes != nil {

		// query param privateRoutes
		var qrPrivateRoutes bool
		if o.PrivateRoutes != nil {
			qrPrivateRoutes = *o.PrivateRoutes
		}
		qPrivateRoutes := swag.FormatBool(qrPrivateRoutes)
		if qPrivateRoutes != "" {
			if err := r.SetQueryParam("privateRoutes", qPrivateRoutes); err != nil {
				return err
			}
		}

	}

	if o.PublicRoutes != nil {

		// query param publicRoutes
		var qrPublicRoutes bool
		if o.PublicRoutes != nil {
			qrPublicRoutes = *o.PublicRoutes
		}
		qPublicRoutes := swag.FormatBool(qrPublicRoutes)
		if qPublicRoutes != "" {
			if err := r.SetQueryParam("publicRoutes", qPublicRoutes); err != nil {
				return err
			}
		}

	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
