// Code generated by go-swagger; DO NOT EDIT.

package guids

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// GETGuidsWithoutCollectionReader is a Reader for the GETGuidsWithoutCollection structure.
type GETGuidsWithoutCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETGuidsWithoutCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETGuidsWithoutCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETGuidsWithoutCollectionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETGuidsWithoutCollectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETGuidsWithoutCollectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETGuidsWithoutCollectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETGuidsWithoutCollectionMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETGuidsWithoutCollectionNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETGuidsWithoutCollectionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETGuidsWithoutCollectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETGuidsWithoutCollectionOK creates a GETGuidsWithoutCollectionOK with default headers values
func NewGETGuidsWithoutCollectionOK() *GETGuidsWithoutCollectionOK {
	return &GETGuidsWithoutCollectionOK{}
}

/*GETGuidsWithoutCollectionOK handles this case with default header values.

OK
*/
type GETGuidsWithoutCollectionOK struct {
	Payload *api_models.PaginatedResponseGUID
}

func (o *GETGuidsWithoutCollectionOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/withoutCollection][%d] getGuidsWithoutCollectionOK  %+v", 200, o.Payload)
}

func (o *GETGuidsWithoutCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PaginatedResponseGUID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGuidsWithoutCollectionAccepted creates a GETGuidsWithoutCollectionAccepted with default headers values
func NewGETGuidsWithoutCollectionAccepted() *GETGuidsWithoutCollectionAccepted {
	return &GETGuidsWithoutCollectionAccepted{}
}

/*GETGuidsWithoutCollectionAccepted handles this case with default header values.

Accepted
*/
type GETGuidsWithoutCollectionAccepted struct {
}

func (o *GETGuidsWithoutCollectionAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/withoutCollection][%d] getGuidsWithoutCollectionAccepted ", 202)
}

func (o *GETGuidsWithoutCollectionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETGuidsWithoutCollectionUnauthorized creates a GETGuidsWithoutCollectionUnauthorized with default headers values
func NewGETGuidsWithoutCollectionUnauthorized() *GETGuidsWithoutCollectionUnauthorized {
	return &GETGuidsWithoutCollectionUnauthorized{}
}

/*GETGuidsWithoutCollectionUnauthorized handles this case with default header values.

Unauthorized
*/
type GETGuidsWithoutCollectionUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETGuidsWithoutCollectionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/withoutCollection][%d] getGuidsWithoutCollectionUnauthorized  %+v", 401, o.Payload)
}

func (o *GETGuidsWithoutCollectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGuidsWithoutCollectionForbidden creates a GETGuidsWithoutCollectionForbidden with default headers values
func NewGETGuidsWithoutCollectionForbidden() *GETGuidsWithoutCollectionForbidden {
	return &GETGuidsWithoutCollectionForbidden{}
}

/*GETGuidsWithoutCollectionForbidden handles this case with default header values.

Forbidden
*/
type GETGuidsWithoutCollectionForbidden struct {
	Payload *api_models.APIError
}

func (o *GETGuidsWithoutCollectionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/withoutCollection][%d] getGuidsWithoutCollectionForbidden  %+v", 403, o.Payload)
}

func (o *GETGuidsWithoutCollectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGuidsWithoutCollectionNotFound creates a GETGuidsWithoutCollectionNotFound with default headers values
func NewGETGuidsWithoutCollectionNotFound() *GETGuidsWithoutCollectionNotFound {
	return &GETGuidsWithoutCollectionNotFound{}
}

/*GETGuidsWithoutCollectionNotFound handles this case with default header values.

Not Found
*/
type GETGuidsWithoutCollectionNotFound struct {
	Payload *api_models.APIError
}

func (o *GETGuidsWithoutCollectionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/withoutCollection][%d] getGuidsWithoutCollectionNotFound  %+v", 404, o.Payload)
}

func (o *GETGuidsWithoutCollectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGuidsWithoutCollectionMethodNotAllowed creates a GETGuidsWithoutCollectionMethodNotAllowed with default headers values
func NewGETGuidsWithoutCollectionMethodNotAllowed() *GETGuidsWithoutCollectionMethodNotAllowed {
	return &GETGuidsWithoutCollectionMethodNotAllowed{}
}

/*GETGuidsWithoutCollectionMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETGuidsWithoutCollectionMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETGuidsWithoutCollectionMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/withoutCollection][%d] getGuidsWithoutCollectionMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETGuidsWithoutCollectionMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGuidsWithoutCollectionNotAcceptable creates a GETGuidsWithoutCollectionNotAcceptable with default headers values
func NewGETGuidsWithoutCollectionNotAcceptable() *GETGuidsWithoutCollectionNotAcceptable {
	return &GETGuidsWithoutCollectionNotAcceptable{}
}

/*GETGuidsWithoutCollectionNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETGuidsWithoutCollectionNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETGuidsWithoutCollectionNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/withoutCollection][%d] getGuidsWithoutCollectionNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETGuidsWithoutCollectionNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGuidsWithoutCollectionUnsupportedMediaType creates a GETGuidsWithoutCollectionUnsupportedMediaType with default headers values
func NewGETGuidsWithoutCollectionUnsupportedMediaType() *GETGuidsWithoutCollectionUnsupportedMediaType {
	return &GETGuidsWithoutCollectionUnsupportedMediaType{}
}

/*GETGuidsWithoutCollectionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETGuidsWithoutCollectionUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETGuidsWithoutCollectionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/withoutCollection][%d] getGuidsWithoutCollectionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETGuidsWithoutCollectionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETGuidsWithoutCollectionInternalServerError creates a GETGuidsWithoutCollectionInternalServerError with default headers values
func NewGETGuidsWithoutCollectionInternalServerError() *GETGuidsWithoutCollectionInternalServerError {
	return &GETGuidsWithoutCollectionInternalServerError{}
}

/*GETGuidsWithoutCollectionInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETGuidsWithoutCollectionInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETGuidsWithoutCollectionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/guids/withoutCollection][%d] getGuidsWithoutCollectionInternalServerError  %+v", 500, o.Payload)
}

func (o *GETGuidsWithoutCollectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
