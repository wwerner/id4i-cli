// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// ListAPIKeyPrivilegesReader is a Reader for the ListAPIKeyPrivileges structure.
type ListAPIKeyPrivilegesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAPIKeyPrivilegesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAPIKeyPrivilegesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewListAPIKeyPrivilegesAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListAPIKeyPrivilegesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListAPIKeyPrivilegesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListAPIKeyPrivilegesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListAPIKeyPrivilegesMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewListAPIKeyPrivilegesNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewListAPIKeyPrivilegesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListAPIKeyPrivilegesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAPIKeyPrivilegesOK creates a ListAPIKeyPrivilegesOK with default headers values
func NewListAPIKeyPrivilegesOK() *ListAPIKeyPrivilegesOK {
	return &ListAPIKeyPrivilegesOK{}
}

/*ListAPIKeyPrivilegesOK handles this case with default header values.

OK
*/
type ListAPIKeyPrivilegesOK struct {
	Payload *api_models.APIKeyPrivilegePaginatedResponse
}

func (o *ListAPIKeyPrivilegesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/{key}/privileges][%d] listApiKeyPrivilegesOK  %+v", 200, o.Payload)
}

func (o *ListAPIKeyPrivilegesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIKeyPrivilegePaginatedResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeyPrivilegesAccepted creates a ListAPIKeyPrivilegesAccepted with default headers values
func NewListAPIKeyPrivilegesAccepted() *ListAPIKeyPrivilegesAccepted {
	return &ListAPIKeyPrivilegesAccepted{}
}

/*ListAPIKeyPrivilegesAccepted handles this case with default header values.

Accepted
*/
type ListAPIKeyPrivilegesAccepted struct {
}

func (o *ListAPIKeyPrivilegesAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/{key}/privileges][%d] listApiKeyPrivilegesAccepted ", 202)
}

func (o *ListAPIKeyPrivilegesAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAPIKeyPrivilegesUnauthorized creates a ListAPIKeyPrivilegesUnauthorized with default headers values
func NewListAPIKeyPrivilegesUnauthorized() *ListAPIKeyPrivilegesUnauthorized {
	return &ListAPIKeyPrivilegesUnauthorized{}
}

/*ListAPIKeyPrivilegesUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAPIKeyPrivilegesUnauthorized struct {
	Payload *api_models.APIError
}

func (o *ListAPIKeyPrivilegesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/{key}/privileges][%d] listApiKeyPrivilegesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAPIKeyPrivilegesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeyPrivilegesForbidden creates a ListAPIKeyPrivilegesForbidden with default headers values
func NewListAPIKeyPrivilegesForbidden() *ListAPIKeyPrivilegesForbidden {
	return &ListAPIKeyPrivilegesForbidden{}
}

/*ListAPIKeyPrivilegesForbidden handles this case with default header values.

Forbidden
*/
type ListAPIKeyPrivilegesForbidden struct {
	Payload *api_models.APIError
}

func (o *ListAPIKeyPrivilegesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/{key}/privileges][%d] listApiKeyPrivilegesForbidden  %+v", 403, o.Payload)
}

func (o *ListAPIKeyPrivilegesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeyPrivilegesNotFound creates a ListAPIKeyPrivilegesNotFound with default headers values
func NewListAPIKeyPrivilegesNotFound() *ListAPIKeyPrivilegesNotFound {
	return &ListAPIKeyPrivilegesNotFound{}
}

/*ListAPIKeyPrivilegesNotFound handles this case with default header values.

Not Found
*/
type ListAPIKeyPrivilegesNotFound struct {
	Payload *api_models.APIError
}

func (o *ListAPIKeyPrivilegesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/{key}/privileges][%d] listApiKeyPrivilegesNotFound  %+v", 404, o.Payload)
}

func (o *ListAPIKeyPrivilegesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeyPrivilegesMethodNotAllowed creates a ListAPIKeyPrivilegesMethodNotAllowed with default headers values
func NewListAPIKeyPrivilegesMethodNotAllowed() *ListAPIKeyPrivilegesMethodNotAllowed {
	return &ListAPIKeyPrivilegesMethodNotAllowed{}
}

/*ListAPIKeyPrivilegesMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type ListAPIKeyPrivilegesMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *ListAPIKeyPrivilegesMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/{key}/privileges][%d] listApiKeyPrivilegesMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListAPIKeyPrivilegesMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeyPrivilegesNotAcceptable creates a ListAPIKeyPrivilegesNotAcceptable with default headers values
func NewListAPIKeyPrivilegesNotAcceptable() *ListAPIKeyPrivilegesNotAcceptable {
	return &ListAPIKeyPrivilegesNotAcceptable{}
}

/*ListAPIKeyPrivilegesNotAcceptable handles this case with default header values.

Not Acceptable
*/
type ListAPIKeyPrivilegesNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *ListAPIKeyPrivilegesNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/{key}/privileges][%d] listApiKeyPrivilegesNotAcceptable  %+v", 406, o.Payload)
}

func (o *ListAPIKeyPrivilegesNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeyPrivilegesUnsupportedMediaType creates a ListAPIKeyPrivilegesUnsupportedMediaType with default headers values
func NewListAPIKeyPrivilegesUnsupportedMediaType() *ListAPIKeyPrivilegesUnsupportedMediaType {
	return &ListAPIKeyPrivilegesUnsupportedMediaType{}
}

/*ListAPIKeyPrivilegesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type ListAPIKeyPrivilegesUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *ListAPIKeyPrivilegesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/{key}/privileges][%d] listApiKeyPrivilegesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListAPIKeyPrivilegesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAPIKeyPrivilegesInternalServerError creates a ListAPIKeyPrivilegesInternalServerError with default headers values
func NewListAPIKeyPrivilegesInternalServerError() *ListAPIKeyPrivilegesInternalServerError {
	return &ListAPIKeyPrivilegesInternalServerError{}
}

/*ListAPIKeyPrivilegesInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAPIKeyPrivilegesInternalServerError struct {
	Payload *api_models.APIError
}

func (o *ListAPIKeyPrivilegesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/apikeys/{key}/privileges][%d] listApiKeyPrivilegesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAPIKeyPrivilegesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
