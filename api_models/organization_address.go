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

// OrganizationAddress OrganizationAddress
// swagger:model OrganizationAddress
type OrganizationAddress struct {

	// city
	// Required: true
	// Max Length: 99
	// Min Length: 2
	City *string `json:"city"`

	// company name
	// Max Length: 254
	// Min Length: 0
	CompanyName *string `json:"companyName,omitempty"`

	// The ISO 3166 two-letter country code
	// Required: true
	// Max Length: 2
	// Min Length: 0
	CountryCode *string `json:"countryCode"`

	// The country name
	// Read Only: true
	CountryName string `json:"countryName,omitempty"`

	// firstname
	// Required: true
	// Max Length: 255
	// Min Length: 0
	Firstname *string `json:"firstname"`

	// lastname
	// Required: true
	// Max Length: 255
	// Min Length: 0
	Lastname *string `json:"lastname"`

	// post code
	// Required: true
	// Max Length: 40
	// Min Length: 2
	PostCode *string `json:"postCode"`

	// street
	// Required: true
	// Max Length: 254
	// Min Length: 3
	Street *string `json:"street"`

	// The telephone number e.g.
	// Max Length: 99
	// Min Length: 0
	Telephone *string `json:"telephone,omitempty"`
}

// Validate validates this organization address
func (m *OrganizationAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompanyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountryCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelephone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrganizationAddress) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	if err := validate.MinLength("city", "body", string(*m.City), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("city", "body", string(*m.City), 99); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationAddress) validateCompanyName(formats strfmt.Registry) error {

	if swag.IsZero(m.CompanyName) { // not required
		return nil
	}

	if err := validate.MinLength("companyName", "body", string(*m.CompanyName), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("companyName", "body", string(*m.CompanyName), 254); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationAddress) validateCountryCode(formats strfmt.Registry) error {

	if err := validate.Required("countryCode", "body", m.CountryCode); err != nil {
		return err
	}

	if err := validate.MinLength("countryCode", "body", string(*m.CountryCode), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("countryCode", "body", string(*m.CountryCode), 2); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationAddress) validateFirstname(formats strfmt.Registry) error {

	if err := validate.Required("firstname", "body", m.Firstname); err != nil {
		return err
	}

	if err := validate.MinLength("firstname", "body", string(*m.Firstname), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("firstname", "body", string(*m.Firstname), 255); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationAddress) validateLastname(formats strfmt.Registry) error {

	if err := validate.Required("lastname", "body", m.Lastname); err != nil {
		return err
	}

	if err := validate.MinLength("lastname", "body", string(*m.Lastname), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("lastname", "body", string(*m.Lastname), 255); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationAddress) validatePostCode(formats strfmt.Registry) error {

	if err := validate.Required("postCode", "body", m.PostCode); err != nil {
		return err
	}

	if err := validate.MinLength("postCode", "body", string(*m.PostCode), 2); err != nil {
		return err
	}

	if err := validate.MaxLength("postCode", "body", string(*m.PostCode), 40); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationAddress) validateStreet(formats strfmt.Registry) error {

	if err := validate.Required("street", "body", m.Street); err != nil {
		return err
	}

	if err := validate.MinLength("street", "body", string(*m.Street), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("street", "body", string(*m.Street), 254); err != nil {
		return err
	}

	return nil
}

func (m *OrganizationAddress) validateTelephone(formats strfmt.Registry) error {

	if swag.IsZero(m.Telephone) { // not required
		return nil
	}

	if err := validate.MinLength("telephone", "body", string(*m.Telephone), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("telephone", "body", string(*m.Telephone), 99); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrganizationAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrganizationAddress) UnmarshalBinary(b []byte) error {
	var res OrganizationAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
