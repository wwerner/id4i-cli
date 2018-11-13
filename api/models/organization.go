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

// Organization Organization
//
// An organization
// swagger:model Organization
type Organization struct {

	// The id of the organization ( Deprecated: Use namespace instead. )
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// URL to a logo of the organization
	// Read Only: true
	LogoURL string `json:"logoURL,omitempty"`

	// The name of the organization
	// Required: true
	// Max Length: 254
	// Min Length: 3
	Name *string `json:"name"`

	// The namespace of the organization
	// Required: true
	// Max Length: 255
	// Min Length: 3
	Namespace *string `json:"namespace"`
}

// Validate validates this organization
func (m *Organization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Organization) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 254); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	if err := validate.MinLength("namespace", "body", string(*m.Namespace), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("namespace", "body", string(*m.Namespace), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Organization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Organization) UnmarshalBinary(b []byte) error {
	var res Organization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}