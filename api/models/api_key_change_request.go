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

// APIKeyChangeRequest ApiKeyChangeRequest
// swagger:model ApiKeyChangeRequest
type APIKeyChangeRequest struct {

	// active
	Active bool `json:"active,omitempty"`

	// new label
	// Required: true
	// Max Length: 50
	// Min Length: 5
	NewLabel *string `json:"newLabel"`
}

// Validate validates this Api key change request
func (m *APIKeyChangeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNewLabel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIKeyChangeRequest) validateNewLabel(formats strfmt.Registry) error {

	if err := validate.Required("newLabel", "body", m.NewLabel); err != nil {
		return err
	}

	if err := validate.MinLength("newLabel", "body", string(*m.NewLabel), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("newLabel", "body", string(*m.NewLabel), 50); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIKeyChangeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIKeyChangeRequest) UnmarshalBinary(b []byte) error {
	var res APIKeyChangeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}