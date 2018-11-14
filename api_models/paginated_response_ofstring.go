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

// PaginatedResponseOfstring PaginatedResponseOfstring
// swagger:model PaginatedResponseOfstring
type PaginatedResponseOfstring struct {

	// elements
	// Required: true
	Elements []string `json:"elements"`

	// The number of returned elements
	// Required: true
	Limit *int32 `json:"limit"`

	// Starting with the n-th element
	// Required: true
	Offset *int32 `json:"offset"`

	// The total number of elements
	Total int32 `json:"total,omitempty"`
}

// Validate validates this paginated response ofstring
func (m *PaginatedResponseOfstring) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateElements(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffset(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaginatedResponseOfstring) validateElements(formats strfmt.Registry) error {

	if err := validate.Required("elements", "body", m.Elements); err != nil {
		return err
	}

	return nil
}

func (m *PaginatedResponseOfstring) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *PaginatedResponseOfstring) validateOffset(formats strfmt.Registry) error {

	if err := validate.Required("offset", "body", m.Offset); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaginatedResponseOfstring) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginatedResponseOfstring) UnmarshalBinary(b []byte) error {
	var res PaginatedResponseOfstring
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
