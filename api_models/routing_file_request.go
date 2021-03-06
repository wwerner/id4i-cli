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

// RoutingFileRequest RoutingFileRequest
// swagger:model RoutingFileRequest
type RoutingFileRequest struct {

	// organization Id
	OrganizationID string `json:"organizationId,omitempty"`

	// routing
	// Required: true
	Routing *RoutingFile `json:"routing"`
}

// Validate validates this routing file request
func (m *RoutingFileRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRouting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RoutingFileRequest) validateRouting(formats strfmt.Registry) error {

	if err := validate.Required("routing", "body", m.Routing); err != nil {
		return err
	}

	if m.Routing != nil {
		if err := m.Routing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RoutingFileRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutingFileRequest) UnmarshalBinary(b []byte) error {
	var res RoutingFileRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
