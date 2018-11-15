// Code generated by go-swagger; DO NOT EDIT.

package api_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Timestamp Timestamp
// swagger:model Timestamp
type Timestamp struct {

	// date
	Date int32 `json:"date,omitempty"`

	// day
	Day int32 `json:"day,omitempty"`

	// hours
	Hours int32 `json:"hours,omitempty"`

	// minutes
	Minutes int32 `json:"minutes,omitempty"`

	// month
	Month int32 `json:"month,omitempty"`

	// nanos
	Nanos int32 `json:"nanos,omitempty"`

	// seconds
	Seconds int32 `json:"seconds,omitempty"`

	// time
	Time int64 `json:"time,omitempty"`

	// timezone offset
	TimezoneOffset int32 `json:"timezoneOffset,omitempty"`

	// year
	Year int32 `json:"year,omitempty"`
}

// Validate validates this timestamp
func (m *Timestamp) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Timestamp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Timestamp) UnmarshalBinary(b []byte) error {
	var res Timestamp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
