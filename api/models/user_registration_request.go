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

// UserRegistrationRequest UserRegistrationRequest
// swagger:model UserRegistrationRequest
type UserRegistrationRequest struct {

	// email
	// Required: true
	Email *string `json:"email"`

	// password
	// Required: true
	// Max Length: 99
	// Min Length: 8
	Password *string `json:"password"`

	// username
	// Required: true
	// Pattern: [a-zA-Z0-9_.-]{6,50}
	Username *string `json:"username"`
}

// Validate validates this user registration request
func (m *UserRegistrationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserRegistrationRequest) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *UserRegistrationRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(*m.Password), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(*m.Password), 99); err != nil {
		return err
	}

	return nil
}

func (m *UserRegistrationRequest) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	if err := validate.Pattern("username", "body", string(*m.Username), `[a-zA-Z0-9_.-]{6,50}`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserRegistrationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserRegistrationRequest) UnmarshalBinary(b []byte) error {
	var res UserRegistrationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}