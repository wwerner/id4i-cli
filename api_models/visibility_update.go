// Code generated by go-swagger; DO NOT EDIT.

package api_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// VisibilityUpdate VisibilityUpdate
// swagger:model VisibilityUpdate
type VisibilityUpdate struct {

	// Document is publicly readable (if ID4N is owned by the same organization)
	Public bool `json:"public,omitempty"`

	// Document is readable by these organizations (independend of ID4N ownership)
	SharedWithOrganizationIds []string `json:"sharedWithOrganizationIds"`
}

// Validate validates this visibility update
func (m *VisibilityUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VisibilityUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VisibilityUpdate) UnmarshalBinary(b []byte) error {
	var res VisibilityUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
