// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// GETPositionsForOrganizationReader is a Reader for the GETPositionsForOrganization structure.
type GETPositionsForOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETPositionsForOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETPositionsForOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETPositionsForOrganizationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETPositionsForOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETPositionsForOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETPositionsForOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETPositionsForOrganizationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETPositionsForOrganizationNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETPositionsForOrganizationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETPositionsForOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETPositionsForOrganizationOK creates a GETPositionsForOrganizationOK with default headers values
func NewGETPositionsForOrganizationOK() *GETPositionsForOrganizationOK {
	return &GETPositionsForOrganizationOK{}
}

/*GETPositionsForOrganizationOK handles this case with default header values.

OK
*/
type GETPositionsForOrganizationOK struct {
	Payload []*api_models.BillingPosition
}

func (o *GETPositionsForOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/billing/{organizationId}/positions][%d] getPositionsForOrganizationOK  %+v", 200, o.Payload)
}

func (o *GETPositionsForOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPositionsForOrganizationAccepted creates a GETPositionsForOrganizationAccepted with default headers values
func NewGETPositionsForOrganizationAccepted() *GETPositionsForOrganizationAccepted {
	return &GETPositionsForOrganizationAccepted{}
}

/*GETPositionsForOrganizationAccepted handles this case with default header values.

Accepted
*/
type GETPositionsForOrganizationAccepted struct {
}

func (o *GETPositionsForOrganizationAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/billing/{organizationId}/positions][%d] getPositionsForOrganizationAccepted ", 202)
}

func (o *GETPositionsForOrganizationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETPositionsForOrganizationUnauthorized creates a GETPositionsForOrganizationUnauthorized with default headers values
func NewGETPositionsForOrganizationUnauthorized() *GETPositionsForOrganizationUnauthorized {
	return &GETPositionsForOrganizationUnauthorized{}
}

/*GETPositionsForOrganizationUnauthorized handles this case with default header values.

Unauthorized
*/
type GETPositionsForOrganizationUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETPositionsForOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/billing/{organizationId}/positions][%d] getPositionsForOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *GETPositionsForOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPositionsForOrganizationForbidden creates a GETPositionsForOrganizationForbidden with default headers values
func NewGETPositionsForOrganizationForbidden() *GETPositionsForOrganizationForbidden {
	return &GETPositionsForOrganizationForbidden{}
}

/*GETPositionsForOrganizationForbidden handles this case with default header values.

Forbidden
*/
type GETPositionsForOrganizationForbidden struct {
	Payload *api_models.APIError
}

func (o *GETPositionsForOrganizationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/billing/{organizationId}/positions][%d] getPositionsForOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *GETPositionsForOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPositionsForOrganizationNotFound creates a GETPositionsForOrganizationNotFound with default headers values
func NewGETPositionsForOrganizationNotFound() *GETPositionsForOrganizationNotFound {
	return &GETPositionsForOrganizationNotFound{}
}

/*GETPositionsForOrganizationNotFound handles this case with default header values.

Not Found
*/
type GETPositionsForOrganizationNotFound struct {
	Payload *api_models.APIError
}

func (o *GETPositionsForOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/billing/{organizationId}/positions][%d] getPositionsForOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *GETPositionsForOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPositionsForOrganizationMethodNotAllowed creates a GETPositionsForOrganizationMethodNotAllowed with default headers values
func NewGETPositionsForOrganizationMethodNotAllowed() *GETPositionsForOrganizationMethodNotAllowed {
	return &GETPositionsForOrganizationMethodNotAllowed{}
}

/*GETPositionsForOrganizationMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETPositionsForOrganizationMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETPositionsForOrganizationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/billing/{organizationId}/positions][%d] getPositionsForOrganizationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETPositionsForOrganizationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPositionsForOrganizationNotAcceptable creates a GETPositionsForOrganizationNotAcceptable with default headers values
func NewGETPositionsForOrganizationNotAcceptable() *GETPositionsForOrganizationNotAcceptable {
	return &GETPositionsForOrganizationNotAcceptable{}
}

/*GETPositionsForOrganizationNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETPositionsForOrganizationNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETPositionsForOrganizationNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/billing/{organizationId}/positions][%d] getPositionsForOrganizationNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETPositionsForOrganizationNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPositionsForOrganizationUnsupportedMediaType creates a GETPositionsForOrganizationUnsupportedMediaType with default headers values
func NewGETPositionsForOrganizationUnsupportedMediaType() *GETPositionsForOrganizationUnsupportedMediaType {
	return &GETPositionsForOrganizationUnsupportedMediaType{}
}

/*GETPositionsForOrganizationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETPositionsForOrganizationUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETPositionsForOrganizationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/billing/{organizationId}/positions][%d] getPositionsForOrganizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETPositionsForOrganizationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETPositionsForOrganizationInternalServerError creates a GETPositionsForOrganizationInternalServerError with default headers values
func NewGETPositionsForOrganizationInternalServerError() *GETPositionsForOrganizationInternalServerError {
	return &GETPositionsForOrganizationInternalServerError{}
}

/*GETPositionsForOrganizationInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETPositionsForOrganizationInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETPositionsForOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/billing/{organizationId}/positions][%d] getPositionsForOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *GETPositionsForOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
