// Code generated by go-swagger; DO NOT EDIT.

package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// RetrieveItemReader is a Reader for the RetrieveItem structure.
type RetrieveItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrieveItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewRetrieveItemAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRetrieveItemUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRetrieveItemForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRetrieveItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewRetrieveItemMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewRetrieveItemNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewRetrieveItemUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRetrieveItemInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveItemOK creates a RetrieveItemOK with default headers values
func NewRetrieveItemOK() *RetrieveItemOK {
	return &RetrieveItemOK{}
}

/*RetrieveItemOK handles this case with default header values.

OK
*/
type RetrieveItemOK struct {
	Payload *api_models.PaginatedHistoryItemResponse
}

func (o *RetrieveItemOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/history/{id4n}/{organizationId}/{sequenceId}][%d] retrieveItemOK  %+v", 200, o.Payload)
}

func (o *RetrieveItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PaginatedHistoryItemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveItemAccepted creates a RetrieveItemAccepted with default headers values
func NewRetrieveItemAccepted() *RetrieveItemAccepted {
	return &RetrieveItemAccepted{}
}

/*RetrieveItemAccepted handles this case with default header values.

Accepted
*/
type RetrieveItemAccepted struct {
}

func (o *RetrieveItemAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/history/{id4n}/{organizationId}/{sequenceId}][%d] retrieveItemAccepted ", 202)
}

func (o *RetrieveItemAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrieveItemUnauthorized creates a RetrieveItemUnauthorized with default headers values
func NewRetrieveItemUnauthorized() *RetrieveItemUnauthorized {
	return &RetrieveItemUnauthorized{}
}

/*RetrieveItemUnauthorized handles this case with default header values.

Unauthorized
*/
type RetrieveItemUnauthorized struct {
	Payload *api_models.APIError
}

func (o *RetrieveItemUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/history/{id4n}/{organizationId}/{sequenceId}][%d] retrieveItemUnauthorized  %+v", 401, o.Payload)
}

func (o *RetrieveItemUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveItemForbidden creates a RetrieveItemForbidden with default headers values
func NewRetrieveItemForbidden() *RetrieveItemForbidden {
	return &RetrieveItemForbidden{}
}

/*RetrieveItemForbidden handles this case with default header values.

Forbidden
*/
type RetrieveItemForbidden struct {
	Payload *api_models.APIError
}

func (o *RetrieveItemForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/history/{id4n}/{organizationId}/{sequenceId}][%d] retrieveItemForbidden  %+v", 403, o.Payload)
}

func (o *RetrieveItemForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveItemNotFound creates a RetrieveItemNotFound with default headers values
func NewRetrieveItemNotFound() *RetrieveItemNotFound {
	return &RetrieveItemNotFound{}
}

/*RetrieveItemNotFound handles this case with default header values.

Not Found
*/
type RetrieveItemNotFound struct {
	Payload *api_models.APIError
}

func (o *RetrieveItemNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/history/{id4n}/{organizationId}/{sequenceId}][%d] retrieveItemNotFound  %+v", 404, o.Payload)
}

func (o *RetrieveItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveItemMethodNotAllowed creates a RetrieveItemMethodNotAllowed with default headers values
func NewRetrieveItemMethodNotAllowed() *RetrieveItemMethodNotAllowed {
	return &RetrieveItemMethodNotAllowed{}
}

/*RetrieveItemMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type RetrieveItemMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *RetrieveItemMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/history/{id4n}/{organizationId}/{sequenceId}][%d] retrieveItemMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *RetrieveItemMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveItemNotAcceptable creates a RetrieveItemNotAcceptable with default headers values
func NewRetrieveItemNotAcceptable() *RetrieveItemNotAcceptable {
	return &RetrieveItemNotAcceptable{}
}

/*RetrieveItemNotAcceptable handles this case with default header values.

Not Acceptable
*/
type RetrieveItemNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *RetrieveItemNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/history/{id4n}/{organizationId}/{sequenceId}][%d] retrieveItemNotAcceptable  %+v", 406, o.Payload)
}

func (o *RetrieveItemNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveItemUnsupportedMediaType creates a RetrieveItemUnsupportedMediaType with default headers values
func NewRetrieveItemUnsupportedMediaType() *RetrieveItemUnsupportedMediaType {
	return &RetrieveItemUnsupportedMediaType{}
}

/*RetrieveItemUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type RetrieveItemUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *RetrieveItemUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/history/{id4n}/{organizationId}/{sequenceId}][%d] retrieveItemUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *RetrieveItemUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveItemInternalServerError creates a RetrieveItemInternalServerError with default headers values
func NewRetrieveItemInternalServerError() *RetrieveItemInternalServerError {
	return &RetrieveItemInternalServerError{}
}

/*RetrieveItemInternalServerError handles this case with default header values.

Internal Server Error
*/
type RetrieveItemInternalServerError struct {
	Payload *api_models.APIError
}

func (o *RetrieveItemInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/history/{id4n}/{organizationId}/{sequenceId}][%d] retrieveItemInternalServerError  %+v", 500, o.Payload)
}

func (o *RetrieveItemInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}