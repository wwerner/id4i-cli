// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// GETOrganizationsOfUserReader is a Reader for the GETOrganizationsOfUser structure.
type GETOrganizationsOfUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETOrganizationsOfUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETOrganizationsOfUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETOrganizationsOfUserAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETOrganizationsOfUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETOrganizationsOfUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETOrganizationsOfUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETOrganizationsOfUserMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETOrganizationsOfUserNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETOrganizationsOfUserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETOrganizationsOfUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETOrganizationsOfUserOK creates a GETOrganizationsOfUserOK with default headers values
func NewGETOrganizationsOfUserOK() *GETOrganizationsOfUserOK {
	return &GETOrganizationsOfUserOK{}
}

/*GETOrganizationsOfUserOK handles this case with default header values.

OK
*/
type GETOrganizationsOfUserOK struct {
	Payload *api_models.PaginatedOrganizationResponse
}

func (o *GETOrganizationsOfUserOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/organizations][%d] getOrganizationsOfUserOK  %+v", 200, o.Payload)
}

func (o *GETOrganizationsOfUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PaginatedOrganizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETOrganizationsOfUserAccepted creates a GETOrganizationsOfUserAccepted with default headers values
func NewGETOrganizationsOfUserAccepted() *GETOrganizationsOfUserAccepted {
	return &GETOrganizationsOfUserAccepted{}
}

/*GETOrganizationsOfUserAccepted handles this case with default header values.

Accepted
*/
type GETOrganizationsOfUserAccepted struct {
}

func (o *GETOrganizationsOfUserAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/organizations][%d] getOrganizationsOfUserAccepted ", 202)
}

func (o *GETOrganizationsOfUserAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETOrganizationsOfUserUnauthorized creates a GETOrganizationsOfUserUnauthorized with default headers values
func NewGETOrganizationsOfUserUnauthorized() *GETOrganizationsOfUserUnauthorized {
	return &GETOrganizationsOfUserUnauthorized{}
}

/*GETOrganizationsOfUserUnauthorized handles this case with default header values.

Unauthorized
*/
type GETOrganizationsOfUserUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETOrganizationsOfUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/organizations][%d] getOrganizationsOfUserUnauthorized  %+v", 401, o.Payload)
}

func (o *GETOrganizationsOfUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETOrganizationsOfUserForbidden creates a GETOrganizationsOfUserForbidden with default headers values
func NewGETOrganizationsOfUserForbidden() *GETOrganizationsOfUserForbidden {
	return &GETOrganizationsOfUserForbidden{}
}

/*GETOrganizationsOfUserForbidden handles this case with default header values.

Forbidden
*/
type GETOrganizationsOfUserForbidden struct {
	Payload *api_models.APIError
}

func (o *GETOrganizationsOfUserForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/organizations][%d] getOrganizationsOfUserForbidden  %+v", 403, o.Payload)
}

func (o *GETOrganizationsOfUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETOrganizationsOfUserNotFound creates a GETOrganizationsOfUserNotFound with default headers values
func NewGETOrganizationsOfUserNotFound() *GETOrganizationsOfUserNotFound {
	return &GETOrganizationsOfUserNotFound{}
}

/*GETOrganizationsOfUserNotFound handles this case with default header values.

Not Found
*/
type GETOrganizationsOfUserNotFound struct {
	Payload *api_models.APIError
}

func (o *GETOrganizationsOfUserNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/organizations][%d] getOrganizationsOfUserNotFound  %+v", 404, o.Payload)
}

func (o *GETOrganizationsOfUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETOrganizationsOfUserMethodNotAllowed creates a GETOrganizationsOfUserMethodNotAllowed with default headers values
func NewGETOrganizationsOfUserMethodNotAllowed() *GETOrganizationsOfUserMethodNotAllowed {
	return &GETOrganizationsOfUserMethodNotAllowed{}
}

/*GETOrganizationsOfUserMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETOrganizationsOfUserMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETOrganizationsOfUserMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/organizations][%d] getOrganizationsOfUserMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETOrganizationsOfUserMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETOrganizationsOfUserNotAcceptable creates a GETOrganizationsOfUserNotAcceptable with default headers values
func NewGETOrganizationsOfUserNotAcceptable() *GETOrganizationsOfUserNotAcceptable {
	return &GETOrganizationsOfUserNotAcceptable{}
}

/*GETOrganizationsOfUserNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETOrganizationsOfUserNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETOrganizationsOfUserNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/organizations][%d] getOrganizationsOfUserNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETOrganizationsOfUserNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETOrganizationsOfUserUnsupportedMediaType creates a GETOrganizationsOfUserUnsupportedMediaType with default headers values
func NewGETOrganizationsOfUserUnsupportedMediaType() *GETOrganizationsOfUserUnsupportedMediaType {
	return &GETOrganizationsOfUserUnsupportedMediaType{}
}

/*GETOrganizationsOfUserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETOrganizationsOfUserUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETOrganizationsOfUserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/organizations][%d] getOrganizationsOfUserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETOrganizationsOfUserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETOrganizationsOfUserInternalServerError creates a GETOrganizationsOfUserInternalServerError with default headers values
func NewGETOrganizationsOfUserInternalServerError() *GETOrganizationsOfUserInternalServerError {
	return &GETOrganizationsOfUserInternalServerError{}
}

/*GETOrganizationsOfUserInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETOrganizationsOfUserInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETOrganizationsOfUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/user/organizations][%d] getOrganizationsOfUserInternalServerError  %+v", 500, o.Payload)
}

func (o *GETOrganizationsOfUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
