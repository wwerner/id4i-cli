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

// GETCollectionsReader is a Reader for the GETCollections structure.
type GETCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GETCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGETCollectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewGETCollectionsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGETCollectionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGETCollectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGETCollectionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewGETCollectionsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGETCollectionsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewGETCollectionsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGETCollectionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGETCollectionsOK creates a GETCollectionsOK with default headers values
func NewGETCollectionsOK() *GETCollectionsOK {
	return &GETCollectionsOK{}
}

/*GETCollectionsOK handles this case with default header values.

OK
*/
type GETCollectionsOK struct {
	Payload *api_models.PaginatedResponseOfGUIDCollection
}

func (o *GETCollectionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}/collections][%d] getCollectionsOK  %+v", 200, o.Payload)
}

func (o *GETCollectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PaginatedResponseOfGUIDCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCollectionsAccepted creates a GETCollectionsAccepted with default headers values
func NewGETCollectionsAccepted() *GETCollectionsAccepted {
	return &GETCollectionsAccepted{}
}

/*GETCollectionsAccepted handles this case with default header values.

Accepted
*/
type GETCollectionsAccepted struct {
}

func (o *GETCollectionsAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}/collections][%d] getCollectionsAccepted ", 202)
}

func (o *GETCollectionsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGETCollectionsUnauthorized creates a GETCollectionsUnauthorized with default headers values
func NewGETCollectionsUnauthorized() *GETCollectionsUnauthorized {
	return &GETCollectionsUnauthorized{}
}

/*GETCollectionsUnauthorized handles this case with default header values.

Unauthorized
*/
type GETCollectionsUnauthorized struct {
	Payload *api_models.APIError
}

func (o *GETCollectionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}/collections][%d] getCollectionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GETCollectionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCollectionsForbidden creates a GETCollectionsForbidden with default headers values
func NewGETCollectionsForbidden() *GETCollectionsForbidden {
	return &GETCollectionsForbidden{}
}

/*GETCollectionsForbidden handles this case with default header values.

Forbidden
*/
type GETCollectionsForbidden struct {
	Payload *api_models.APIError
}

func (o *GETCollectionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}/collections][%d] getCollectionsForbidden  %+v", 403, o.Payload)
}

func (o *GETCollectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCollectionsNotFound creates a GETCollectionsNotFound with default headers values
func NewGETCollectionsNotFound() *GETCollectionsNotFound {
	return &GETCollectionsNotFound{}
}

/*GETCollectionsNotFound handles this case with default header values.

Not Found
*/
type GETCollectionsNotFound struct {
	Payload *api_models.APIError
}

func (o *GETCollectionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}/collections][%d] getCollectionsNotFound  %+v", 404, o.Payload)
}

func (o *GETCollectionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCollectionsMethodNotAllowed creates a GETCollectionsMethodNotAllowed with default headers values
func NewGETCollectionsMethodNotAllowed() *GETCollectionsMethodNotAllowed {
	return &GETCollectionsMethodNotAllowed{}
}

/*GETCollectionsMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type GETCollectionsMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *GETCollectionsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}/collections][%d] getCollectionsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GETCollectionsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCollectionsNotAcceptable creates a GETCollectionsNotAcceptable with default headers values
func NewGETCollectionsNotAcceptable() *GETCollectionsNotAcceptable {
	return &GETCollectionsNotAcceptable{}
}

/*GETCollectionsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type GETCollectionsNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *GETCollectionsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}/collections][%d] getCollectionsNotAcceptable  %+v", 406, o.Payload)
}

func (o *GETCollectionsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCollectionsUnsupportedMediaType creates a GETCollectionsUnsupportedMediaType with default headers values
func NewGETCollectionsUnsupportedMediaType() *GETCollectionsUnsupportedMediaType {
	return &GETCollectionsUnsupportedMediaType{}
}

/*GETCollectionsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type GETCollectionsUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *GETCollectionsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}/collections][%d] getCollectionsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GETCollectionsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGETCollectionsInternalServerError creates a GETCollectionsInternalServerError with default headers values
func NewGETCollectionsInternalServerError() *GETCollectionsInternalServerError {
	return &GETCollectionsInternalServerError{}
}

/*GETCollectionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GETCollectionsInternalServerError struct {
	Payload *api_models.APIError
}

func (o *GETCollectionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/id4ns/{id4n}/collections][%d] getCollectionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GETCollectionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
