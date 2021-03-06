// Code generated by go-swagger; DO NOT EDIT.

package api_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PublicImagePresentation PublicImagePresentation
// swagger:model PublicImagePresentation
type PublicImagePresentation struct {

	// The uri/url of the image
	// Read Only: true
	URI string `json:"uri,omitempty"`
}

// Validate validates this public image presentation
func (m *PublicImagePresentation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PublicImagePresentation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicImagePresentation) UnmarshalBinary(b []byte) error {
	var res PublicImagePresentation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
