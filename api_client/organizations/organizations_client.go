// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new organizations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for organizations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddPartnerOrganization adds partner

Adding a partner organization. If the given organization is already a partner the result will be state 200 too.
*/
func (a *Client) AddPartnerOrganization(params *AddPartnerOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*AddPartnerOrganizationOK, *AddPartnerOrganizationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPartnerOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addPartnerOrganization",
		Method:             "PUT",
		PathPattern:        "/api/v1/organizations/{organizationId}/partner",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddPartnerOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *AddPartnerOrganizationOK:
		return value, nil, nil
	case *AddPartnerOrganizationAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
CreateOrganization creates organization

Creating a new organization.
*/
func (a *Client) CreateOrganization(params *CreateOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOrganizationOK, *CreateOrganizationCreated, *CreateOrganizationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOrganization",
		Method:             "POST",
		PathPattern:        "/api/v1/organizations",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *CreateOrganizationOK:
		return value, nil, nil, nil
	case *CreateOrganizationCreated:
		return nil, value, nil, nil
	case *CreateOrganizationAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
DeleteOrganization deletes organization
*/
func (a *Client) DeleteOrganization(params *DeleteOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrganization",
		Method:             "DELETE",
		PathPattern:        "/api/v1/organizations/{organizationId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrganizationOK), nil

}

/*
DeleteOrganizationBillingAddress removes billing address
*/
func (a *Client) DeleteOrganizationBillingAddress(params *DeleteOrganizationBillingAddressParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationBillingAddressOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationBillingAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrganizationBillingAddress",
		Method:             "DELETE",
		PathPattern:        "/api/v1/organizations/{organizationId}/addresses/billing",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOrganizationBillingAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrganizationBillingAddressOK), nil

}

/*
DeleteOrganizationLogo deletes organization logo
*/
func (a *Client) DeleteOrganizationLogo(params *DeleteOrganizationLogoParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOrganizationLogoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationLogoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteOrganizationLogo",
		Method:             "DELETE",
		PathPattern:        "/api/v1/organizations/{organizationId}/logo",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOrganizationLogoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOrganizationLogoOK), nil

}

/*
FindOrganization finds organization by id namespace

Returns a single organization.
*/
func (a *Client) FindOrganization(params *FindOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*FindOrganizationOK, *FindOrganizationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findOrganization",
		Method:             "GET",
		PathPattern:        "/api/v1/organizations/{organizationId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *FindOrganizationOK:
		return value, nil, nil
	case *FindOrganizationAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
FindOrganizationAddress retrieves address
*/
func (a *Client) FindOrganizationAddress(params *FindOrganizationAddressParams, authInfo runtime.ClientAuthInfoWriter) (*FindOrganizationAddressOK, *FindOrganizationAddressAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindOrganizationAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findOrganizationAddress",
		Method:             "GET",
		PathPattern:        "/api/v1/organizations/{organizationId}/addresses/default",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindOrganizationAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *FindOrganizationAddressOK:
		return value, nil, nil
	case *FindOrganizationAddressAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
FindOrganizationBillingAddress retrieves billing address
*/
func (a *Client) FindOrganizationBillingAddress(params *FindOrganizationBillingAddressParams, authInfo runtime.ClientAuthInfoWriter) (*FindOrganizationBillingAddressOK, *FindOrganizationBillingAddressAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindOrganizationBillingAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "findOrganizationBillingAddress",
		Method:             "GET",
		PathPattern:        "/api/v1/organizations/{organizationId}/addresses/billing",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindOrganizationBillingAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *FindOrganizationBillingAddressOK:
		return value, nil, nil
	case *FindOrganizationBillingAddressAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GETOrganizationPrivileges lists my privileges

Listing all privileges of the current user/APIKey of the specified organization.
*/
func (a *Client) GETOrganizationPrivileges(params *GETOrganizationPrivilegesParams, authInfo runtime.ClientAuthInfoWriter) (*GETOrganizationPrivilegesOK, *GETOrganizationPrivilegesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETOrganizationPrivilegesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOrganizationPrivileges",
		Method:             "GET",
		PathPattern:        "/api/v1/organizations/{organizationId}/privileges",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETOrganizationPrivilegesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETOrganizationPrivilegesOK:
		return value, nil, nil
	case *GETOrganizationPrivilegesAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
GETPartnerOrganizations Gets partners of an organization

Listing partners in a paginated manner.
*/
func (a *Client) GETPartnerOrganizations(params *GETPartnerOrganizationsParams, authInfo runtime.ClientAuthInfoWriter) (*GETPartnerOrganizationsOK, *GETPartnerOrganizationsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGETPartnerOrganizationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPartnerOrganizations",
		Method:             "GET",
		PathPattern:        "/api/v1/organizations/{organizationId}/partner",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GETPartnerOrganizationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *GETPartnerOrganizationsOK:
		return value, nil, nil
	case *GETPartnerOrganizationsAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
ListCountries lists countries
*/
func (a *Client) ListCountries(params *ListCountriesParams, authInfo runtime.ClientAuthInfoWriter) (*ListCountriesOK, *ListCountriesAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCountriesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listCountries",
		Method:             "GET",
		PathPattern:        "/api/v1/countries",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListCountriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ListCountriesOK:
		return value, nil, nil
	case *ListCountriesAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
RemovePartnerOrganization removes partner

Removing a partner organization
*/
func (a *Client) RemovePartnerOrganization(params *RemovePartnerOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*RemovePartnerOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemovePartnerOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "removePartnerOrganization",
		Method:             "DELETE",
		PathPattern:        "/api/v1/organizations/{organizationId}/partner",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RemovePartnerOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RemovePartnerOrganizationOK), nil

}

/*
SetOrganizationLogo updates organization logo

Updating an organization logo using a multipart file upload.
*/
func (a *Client) SetOrganizationLogo(params *SetOrganizationLogoParams, authInfo runtime.ClientAuthInfoWriter) (*SetOrganizationLogoOK, *SetOrganizationLogoCreated, *SetOrganizationLogoAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetOrganizationLogoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setOrganizationLogo",
		Method:             "POST",
		PathPattern:        "/api/v1/organizations/{organizationId}/logo",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetOrganizationLogoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *SetOrganizationLogoOK:
		return value, nil, nil, nil
	case *SetOrganizationLogoCreated:
		return nil, value, nil, nil
	case *SetOrganizationLogoAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
UpdateOrganization updates organization
*/
func (a *Client) UpdateOrganization(params *UpdateOrganizationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationOK, *UpdateOrganizationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOrganization",
		Method:             "PUT",
		PathPattern:        "/api/v1/organizations/{organizationId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateOrganizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateOrganizationOK:
		return value, nil, nil
	case *UpdateOrganizationAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
UpdateOrganizationAddress stores address
*/
func (a *Client) UpdateOrganizationAddress(params *UpdateOrganizationAddressParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationAddressOK, *UpdateOrganizationAddressAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOrganizationAddress",
		Method:             "PUT",
		PathPattern:        "/api/v1/organizations/{organizationId}/addresses/default",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateOrganizationAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateOrganizationAddressOK:
		return value, nil, nil
	case *UpdateOrganizationAddressAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
UpdateOrganizationBillingAddress stores billing address
*/
func (a *Client) UpdateOrganizationBillingAddress(params *UpdateOrganizationBillingAddressParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateOrganizationBillingAddressOK, *UpdateOrganizationBillingAddressAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateOrganizationBillingAddressParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateOrganizationBillingAddress",
		Method:             "PUT",
		PathPattern:        "/api/v1/organizations/{organizationId}/addresses/billing",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateOrganizationBillingAddressReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *UpdateOrganizationBillingAddressOK:
		return value, nil, nil
	case *UpdateOrganizationBillingAddressAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
