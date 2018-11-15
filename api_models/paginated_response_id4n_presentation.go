// Code generated by go-swagger; DO NOT EDIT.

package api_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PaginatedResponseID4NPresentation PaginatedResponse«Id4nPresentation»
// swagger:model PaginatedResponse«Id4nPresentation»
type PaginatedResponseID4NPresentation struct {

	// elements
	// Required: true
	Elements []*ID4NPresentation `json:"elements"`

	// The number of returned elements
	// Required: true
	Limit *int32 `json:"limit"`

	// Starting with the n-th element
	// Required: true
	Offset *int32 `json:"offset"`

	// The total number of elements
	Total int32 `json:"total,omitempty"`
}

// Validate validates this paginated response Id4n presentation
func (m *PaginatedResponseID4NPresentation) Validate(formats strfmt.Registry) error {
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

func (m *PaginatedResponseID4NPresentation) validateElements(formats strfmt.Registry) error {

	if err := validate.Required("elements", "body", m.Elements); err != nil {
		return err
	}

	for i := 0; i < len(m.Elements); i++ {
		if swag.IsZero(m.Elements[i]) { // not required
			continue
		}

		if m.Elements[i] != nil {
			if err := m.Elements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("elements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PaginatedResponseID4NPresentation) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *PaginatedResponseID4NPresentation) validateOffset(formats strfmt.Registry) error {

	if err := validate.Required("offset", "body", m.Offset); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaginatedResponseID4NPresentation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaginatedResponseID4NPresentation) UnmarshalBinary(b []byte) error {
	var res PaginatedResponseID4NPresentation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
