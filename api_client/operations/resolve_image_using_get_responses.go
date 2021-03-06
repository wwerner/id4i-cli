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

// ResolveImageUsingGETReader is a Reader for the ResolveImageUsingGET structure.
type ResolveImageUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResolveImageUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewResolveImageUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewResolveImageUsingGETAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewResolveImageUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewResolveImageUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewResolveImageUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewResolveImageUsingGETMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewResolveImageUsingGETNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewResolveImageUsingGETUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewResolveImageUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResolveImageUsingGETOK creates a ResolveImageUsingGETOK with default headers values
func NewResolveImageUsingGETOK() *ResolveImageUsingGETOK {
	return &ResolveImageUsingGETOK{}
}

/*ResolveImageUsingGETOK handles this case with default header values.

OK
*/
type ResolveImageUsingGETOK struct {
	Payload strfmt.Base64
}

func (o *ResolveImageUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/image/{imageID}][%d] resolveImageUsingGetOK  %+v", 200, o.Payload)
}

func (o *ResolveImageUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResolveImageUsingGETAccepted creates a ResolveImageUsingGETAccepted with default headers values
func NewResolveImageUsingGETAccepted() *ResolveImageUsingGETAccepted {
	return &ResolveImageUsingGETAccepted{}
}

/*ResolveImageUsingGETAccepted handles this case with default header values.

Accepted
*/
type ResolveImageUsingGETAccepted struct {
}

func (o *ResolveImageUsingGETAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/image/{imageID}][%d] resolveImageUsingGetAccepted ", 202)
}

func (o *ResolveImageUsingGETAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResolveImageUsingGETUnauthorized creates a ResolveImageUsingGETUnauthorized with default headers values
func NewResolveImageUsingGETUnauthorized() *ResolveImageUsingGETUnauthorized {
	return &ResolveImageUsingGETUnauthorized{}
}

/*ResolveImageUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type ResolveImageUsingGETUnauthorized struct {
	Payload *api_models.APIError
}

func (o *ResolveImageUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/image/{imageID}][%d] resolveImageUsingGetUnauthorized  %+v", 401, o.Payload)
}

func (o *ResolveImageUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResolveImageUsingGETForbidden creates a ResolveImageUsingGETForbidden with default headers values
func NewResolveImageUsingGETForbidden() *ResolveImageUsingGETForbidden {
	return &ResolveImageUsingGETForbidden{}
}

/*ResolveImageUsingGETForbidden handles this case with default header values.

Forbidden
*/
type ResolveImageUsingGETForbidden struct {
	Payload *api_models.APIError
}

func (o *ResolveImageUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/image/{imageID}][%d] resolveImageUsingGetForbidden  %+v", 403, o.Payload)
}

func (o *ResolveImageUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResolveImageUsingGETNotFound creates a ResolveImageUsingGETNotFound with default headers values
func NewResolveImageUsingGETNotFound() *ResolveImageUsingGETNotFound {
	return &ResolveImageUsingGETNotFound{}
}

/*ResolveImageUsingGETNotFound handles this case with default header values.

Not Found
*/
type ResolveImageUsingGETNotFound struct {
	Payload *api_models.APIError
}

func (o *ResolveImageUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/image/{imageID}][%d] resolveImageUsingGetNotFound  %+v", 404, o.Payload)
}

func (o *ResolveImageUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResolveImageUsingGETMethodNotAllowed creates a ResolveImageUsingGETMethodNotAllowed with default headers values
func NewResolveImageUsingGETMethodNotAllowed() *ResolveImageUsingGETMethodNotAllowed {
	return &ResolveImageUsingGETMethodNotAllowed{}
}

/*ResolveImageUsingGETMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type ResolveImageUsingGETMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *ResolveImageUsingGETMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/image/{imageID}][%d] resolveImageUsingGetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ResolveImageUsingGETMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResolveImageUsingGETNotAcceptable creates a ResolveImageUsingGETNotAcceptable with default headers values
func NewResolveImageUsingGETNotAcceptable() *ResolveImageUsingGETNotAcceptable {
	return &ResolveImageUsingGETNotAcceptable{}
}

/*ResolveImageUsingGETNotAcceptable handles this case with default header values.

Not Acceptable
*/
type ResolveImageUsingGETNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *ResolveImageUsingGETNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/image/{imageID}][%d] resolveImageUsingGetNotAcceptable  %+v", 406, o.Payload)
}

func (o *ResolveImageUsingGETNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResolveImageUsingGETUnsupportedMediaType creates a ResolveImageUsingGETUnsupportedMediaType with default headers values
func NewResolveImageUsingGETUnsupportedMediaType() *ResolveImageUsingGETUnsupportedMediaType {
	return &ResolveImageUsingGETUnsupportedMediaType{}
}

/*ResolveImageUsingGETUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type ResolveImageUsingGETUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *ResolveImageUsingGETUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/image/{imageID}][%d] resolveImageUsingGetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ResolveImageUsingGETUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResolveImageUsingGETInternalServerError creates a ResolveImageUsingGETInternalServerError with default headers values
func NewResolveImageUsingGETInternalServerError() *ResolveImageUsingGETInternalServerError {
	return &ResolveImageUsingGETInternalServerError{}
}

/*ResolveImageUsingGETInternalServerError handles this case with default header values.

Internal Server Error
*/
type ResolveImageUsingGETInternalServerError struct {
	Payload *api_models.APIError
}

func (o *ResolveImageUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/image/{imageID}][%d] resolveImageUsingGetInternalServerError  %+v", 500, o.Payload)
}

func (o *ResolveImageUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
