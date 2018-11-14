// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// FindOrganizationAddressReader is a Reader for the FindOrganizationAddress structure.
type FindOrganizationAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindOrganizationAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindOrganizationAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewFindOrganizationAddressAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewFindOrganizationAddressUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewFindOrganizationAddressForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewFindOrganizationAddressNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewFindOrganizationAddressMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewFindOrganizationAddressNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewFindOrganizationAddressUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewFindOrganizationAddressInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFindOrganizationAddressOK creates a FindOrganizationAddressOK with default headers values
func NewFindOrganizationAddressOK() *FindOrganizationAddressOK {
	return &FindOrganizationAddressOK{}
}

/*FindOrganizationAddressOK handles this case with default header values.

OK
*/
type FindOrganizationAddressOK struct {
	Payload *api_models.OrganizationAddress
}

func (o *FindOrganizationAddressOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/addresses/default][%d] findOrganizationAddressOK  %+v", 200, o.Payload)
}

func (o *FindOrganizationAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.OrganizationAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindOrganizationAddressAccepted creates a FindOrganizationAddressAccepted with default headers values
func NewFindOrganizationAddressAccepted() *FindOrganizationAddressAccepted {
	return &FindOrganizationAddressAccepted{}
}

/*FindOrganizationAddressAccepted handles this case with default header values.

Accepted
*/
type FindOrganizationAddressAccepted struct {
}

func (o *FindOrganizationAddressAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/addresses/default][%d] findOrganizationAddressAccepted ", 202)
}

func (o *FindOrganizationAddressAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindOrganizationAddressUnauthorized creates a FindOrganizationAddressUnauthorized with default headers values
func NewFindOrganizationAddressUnauthorized() *FindOrganizationAddressUnauthorized {
	return &FindOrganizationAddressUnauthorized{}
}

/*FindOrganizationAddressUnauthorized handles this case with default header values.

Unauthorized
*/
type FindOrganizationAddressUnauthorized struct {
	Payload *api_models.APIError
}

func (o *FindOrganizationAddressUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/addresses/default][%d] findOrganizationAddressUnauthorized  %+v", 401, o.Payload)
}

func (o *FindOrganizationAddressUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindOrganizationAddressForbidden creates a FindOrganizationAddressForbidden with default headers values
func NewFindOrganizationAddressForbidden() *FindOrganizationAddressForbidden {
	return &FindOrganizationAddressForbidden{}
}

/*FindOrganizationAddressForbidden handles this case with default header values.

Forbidden
*/
type FindOrganizationAddressForbidden struct {
	Payload *api_models.APIError
}

func (o *FindOrganizationAddressForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/addresses/default][%d] findOrganizationAddressForbidden  %+v", 403, o.Payload)
}

func (o *FindOrganizationAddressForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindOrganizationAddressNotFound creates a FindOrganizationAddressNotFound with default headers values
func NewFindOrganizationAddressNotFound() *FindOrganizationAddressNotFound {
	return &FindOrganizationAddressNotFound{}
}

/*FindOrganizationAddressNotFound handles this case with default header values.

Not Found
*/
type FindOrganizationAddressNotFound struct {
	Payload *api_models.APIError
}

func (o *FindOrganizationAddressNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/addresses/default][%d] findOrganizationAddressNotFound  %+v", 404, o.Payload)
}

func (o *FindOrganizationAddressNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindOrganizationAddressMethodNotAllowed creates a FindOrganizationAddressMethodNotAllowed with default headers values
func NewFindOrganizationAddressMethodNotAllowed() *FindOrganizationAddressMethodNotAllowed {
	return &FindOrganizationAddressMethodNotAllowed{}
}

/*FindOrganizationAddressMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type FindOrganizationAddressMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *FindOrganizationAddressMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/addresses/default][%d] findOrganizationAddressMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *FindOrganizationAddressMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindOrganizationAddressNotAcceptable creates a FindOrganizationAddressNotAcceptable with default headers values
func NewFindOrganizationAddressNotAcceptable() *FindOrganizationAddressNotAcceptable {
	return &FindOrganizationAddressNotAcceptable{}
}

/*FindOrganizationAddressNotAcceptable handles this case with default header values.

Not Acceptable
*/
type FindOrganizationAddressNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *FindOrganizationAddressNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/addresses/default][%d] findOrganizationAddressNotAcceptable  %+v", 406, o.Payload)
}

func (o *FindOrganizationAddressNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindOrganizationAddressUnsupportedMediaType creates a FindOrganizationAddressUnsupportedMediaType with default headers values
func NewFindOrganizationAddressUnsupportedMediaType() *FindOrganizationAddressUnsupportedMediaType {
	return &FindOrganizationAddressUnsupportedMediaType{}
}

/*FindOrganizationAddressUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type FindOrganizationAddressUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *FindOrganizationAddressUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/addresses/default][%d] findOrganizationAddressUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *FindOrganizationAddressUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindOrganizationAddressInternalServerError creates a FindOrganizationAddressInternalServerError with default headers values
func NewFindOrganizationAddressInternalServerError() *FindOrganizationAddressInternalServerError {
	return &FindOrganizationAddressInternalServerError{}
}

/*FindOrganizationAddressInternalServerError handles this case with default header values.

Internal Server Error
*/
type FindOrganizationAddressInternalServerError struct {
	Payload *api_models.APIError
}

func (o *FindOrganizationAddressInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/organizations/{organizationId}/addresses/default][%d] findOrganizationAddressInternalServerError  %+v", 500, o.Payload)
}

func (o *FindOrganizationAddressInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
