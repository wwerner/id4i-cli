// Code generated by go-swagger; DO NOT EDIT.

package api_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ChangeLogEntry ChangeLogEntry
//
// A changelog entry
// swagger:model ChangeLogEntry
type ChangeLogEntry struct {

	// The unique id of the changelog entry
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The message as template or rendered as plain text
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The values of the properties in the message. May be nested as object with a value field
	// Read Only: true
	MessageProperties map[string]interface{} `json:"messageProperties,omitempty"`

	// The UTC unix timestamp when this change occurred
	// Read Only: true
	Timestamp int64 `json:"timestamp,omitempty"`
}

// Validate validates this change log entry
func (m *ChangeLogEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChangeLogEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeLogEntry) UnmarshalBinary(b []byte) error {
	var res ChangeLogEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
