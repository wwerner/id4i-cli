// Code generated by go-swagger; DO NOT EDIT.

package api_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddAPIKeyPrivilegeRequest AddApiKeyPrivilegeRequest
// swagger:model AddApiKeyPrivilegeRequest
type AddAPIKeyPrivilegeRequest struct {

	// privilege
	// Required: true
	Privilege *string `json:"privilege"`
}

// Validate validates this add Api key privilege request
func (m *AddAPIKeyPrivilegeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivilege(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddAPIKeyPrivilegeRequest) validatePrivilege(formats strfmt.Registry) error {

	if err := validate.Required("privilege", "body", m.Privilege); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddAPIKeyPrivilegeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddAPIKeyPrivilegeRequest) UnmarshalBinary(b []byte) error {
	var res AddAPIKeyPrivilegeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}