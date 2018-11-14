// Code generated by go-swagger; DO NOT EDIT.

package public_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// ReadOrganizationInfoReader is a Reader for the ReadOrganizationInfo structure.
type ReadOrganizationInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadOrganizationInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReadOrganizationInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewReadOrganizationInfoAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewReadOrganizationInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewReadOrganizationInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewReadOrganizationInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewReadOrganizationInfoMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewReadOrganizationInfoNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewReadOrganizationInfoUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewReadOrganizationInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadOrganizationInfoOK creates a ReadOrganizationInfoOK with default headers values
func NewReadOrganizationInfoOK() *ReadOrganizationInfoOK {
	return &ReadOrganizationInfoOK{}
}

/*ReadOrganizationInfoOK handles this case with default header values.

OK
*/
type ReadOrganizationInfoOK struct {
	Payload *api_models.Organization
}

func (o *ReadOrganizationInfoOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/organizations/{organizationId}][%d] readOrganizationInfoOK  %+v", 200, o.Payload)
}

func (o *ReadOrganizationInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOrganizationInfoAccepted creates a ReadOrganizationInfoAccepted with default headers values
func NewReadOrganizationInfoAccepted() *ReadOrganizationInfoAccepted {
	return &ReadOrganizationInfoAccepted{}
}

/*ReadOrganizationInfoAccepted handles this case with default header values.

Accepted
*/
type ReadOrganizationInfoAccepted struct {
}

func (o *ReadOrganizationInfoAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/organizations/{organizationId}][%d] readOrganizationInfoAccepted ", 202)
}

func (o *ReadOrganizationInfoAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewReadOrganizationInfoUnauthorized creates a ReadOrganizationInfoUnauthorized with default headers values
func NewReadOrganizationInfoUnauthorized() *ReadOrganizationInfoUnauthorized {
	return &ReadOrganizationInfoUnauthorized{}
}

/*ReadOrganizationInfoUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadOrganizationInfoUnauthorized struct {
	Payload *api_models.APIError
}

func (o *ReadOrganizationInfoUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/organizations/{organizationId}][%d] readOrganizationInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadOrganizationInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOrganizationInfoForbidden creates a ReadOrganizationInfoForbidden with default headers values
func NewReadOrganizationInfoForbidden() *ReadOrganizationInfoForbidden {
	return &ReadOrganizationInfoForbidden{}
}

/*ReadOrganizationInfoForbidden handles this case with default header values.

Forbidden
*/
type ReadOrganizationInfoForbidden struct {
	Payload *api_models.APIError
}

func (o *ReadOrganizationInfoForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/organizations/{organizationId}][%d] readOrganizationInfoForbidden  %+v", 403, o.Payload)
}

func (o *ReadOrganizationInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOrganizationInfoNotFound creates a ReadOrganizationInfoNotFound with default headers values
func NewReadOrganizationInfoNotFound() *ReadOrganizationInfoNotFound {
	return &ReadOrganizationInfoNotFound{}
}

/*ReadOrganizationInfoNotFound handles this case with default header values.

Not Found
*/
type ReadOrganizationInfoNotFound struct {
	Payload *api_models.APIError
}

func (o *ReadOrganizationInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/organizations/{organizationId}][%d] readOrganizationInfoNotFound  %+v", 404, o.Payload)
}

func (o *ReadOrganizationInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOrganizationInfoMethodNotAllowed creates a ReadOrganizationInfoMethodNotAllowed with default headers values
func NewReadOrganizationInfoMethodNotAllowed() *ReadOrganizationInfoMethodNotAllowed {
	return &ReadOrganizationInfoMethodNotAllowed{}
}

/*ReadOrganizationInfoMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type ReadOrganizationInfoMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *ReadOrganizationInfoMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/organizations/{organizationId}][%d] readOrganizationInfoMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ReadOrganizationInfoMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOrganizationInfoNotAcceptable creates a ReadOrganizationInfoNotAcceptable with default headers values
func NewReadOrganizationInfoNotAcceptable() *ReadOrganizationInfoNotAcceptable {
	return &ReadOrganizationInfoNotAcceptable{}
}

/*ReadOrganizationInfoNotAcceptable handles this case with default header values.

Not Acceptable
*/
type ReadOrganizationInfoNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *ReadOrganizationInfoNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/organizations/{organizationId}][%d] readOrganizationInfoNotAcceptable  %+v", 406, o.Payload)
}

func (o *ReadOrganizationInfoNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOrganizationInfoUnsupportedMediaType creates a ReadOrganizationInfoUnsupportedMediaType with default headers values
func NewReadOrganizationInfoUnsupportedMediaType() *ReadOrganizationInfoUnsupportedMediaType {
	return &ReadOrganizationInfoUnsupportedMediaType{}
}

/*ReadOrganizationInfoUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type ReadOrganizationInfoUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *ReadOrganizationInfoUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/organizations/{organizationId}][%d] readOrganizationInfoUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ReadOrganizationInfoUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadOrganizationInfoInternalServerError creates a ReadOrganizationInfoInternalServerError with default headers values
func NewReadOrganizationInfoInternalServerError() *ReadOrganizationInfoInternalServerError {
	return &ReadOrganizationInfoInternalServerError{}
}

/*ReadOrganizationInfoInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadOrganizationInfoInternalServerError struct {
	Payload *api_models.APIError
}

func (o *ReadOrganizationInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/organizations/{organizationId}][%d] readOrganizationInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadOrganizationInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
