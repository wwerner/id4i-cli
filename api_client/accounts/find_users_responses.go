// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// FindUsersReader is a Reader for the FindUsers structure.
type FindUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewFindUsersAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewFindUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewFindUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewFindUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewFindUsersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewFindUsersNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewFindUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewFindUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFindUsersOK creates a FindUsersOK with default headers values
func NewFindUsersOK() *FindUsersOK {
	return &FindUsersOK{}
}

/*FindUsersOK handles this case with default header values.

OK
*/
type FindUsersOK struct {
	Payload *api_models.PaginatedUserPresentationResponse
}

func (o *FindUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] findUsersOK  %+v", 200, o.Payload)
}

func (o *FindUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PaginatedUserPresentationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUsersAccepted creates a FindUsersAccepted with default headers values
func NewFindUsersAccepted() *FindUsersAccepted {
	return &FindUsersAccepted{}
}

/*FindUsersAccepted handles this case with default header values.

Accepted
*/
type FindUsersAccepted struct {
}

func (o *FindUsersAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] findUsersAccepted ", 202)
}

func (o *FindUsersAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFindUsersUnauthorized creates a FindUsersUnauthorized with default headers values
func NewFindUsersUnauthorized() *FindUsersUnauthorized {
	return &FindUsersUnauthorized{}
}

/*FindUsersUnauthorized handles this case with default header values.

Unauthorized
*/
type FindUsersUnauthorized struct {
	Payload *api_models.APIError
}

func (o *FindUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] findUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *FindUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUsersForbidden creates a FindUsersForbidden with default headers values
func NewFindUsersForbidden() *FindUsersForbidden {
	return &FindUsersForbidden{}
}

/*FindUsersForbidden handles this case with default header values.

Forbidden
*/
type FindUsersForbidden struct {
	Payload *api_models.APIError
}

func (o *FindUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] findUsersForbidden  %+v", 403, o.Payload)
}

func (o *FindUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUsersNotFound creates a FindUsersNotFound with default headers values
func NewFindUsersNotFound() *FindUsersNotFound {
	return &FindUsersNotFound{}
}

/*FindUsersNotFound handles this case with default header values.

Not Found
*/
type FindUsersNotFound struct {
	Payload *api_models.APIError
}

func (o *FindUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] findUsersNotFound  %+v", 404, o.Payload)
}

func (o *FindUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUsersMethodNotAllowed creates a FindUsersMethodNotAllowed with default headers values
func NewFindUsersMethodNotAllowed() *FindUsersMethodNotAllowed {
	return &FindUsersMethodNotAllowed{}
}

/*FindUsersMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type FindUsersMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *FindUsersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] findUsersMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *FindUsersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUsersNotAcceptable creates a FindUsersNotAcceptable with default headers values
func NewFindUsersNotAcceptable() *FindUsersNotAcceptable {
	return &FindUsersNotAcceptable{}
}

/*FindUsersNotAcceptable handles this case with default header values.

Not Acceptable
*/
type FindUsersNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *FindUsersNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] findUsersNotAcceptable  %+v", 406, o.Payload)
}

func (o *FindUsersNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUsersUnsupportedMediaType creates a FindUsersUnsupportedMediaType with default headers values
func NewFindUsersUnsupportedMediaType() *FindUsersUnsupportedMediaType {
	return &FindUsersUnsupportedMediaType{}
}

/*FindUsersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type FindUsersUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *FindUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] findUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *FindUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindUsersInternalServerError creates a FindUsersInternalServerError with default headers values
func NewFindUsersInternalServerError() *FindUsersInternalServerError {
	return &FindUsersInternalServerError{}
}

/*FindUsersInternalServerError handles this case with default header values.

Internal Server Error
*/
type FindUsersInternalServerError struct {
	Payload *api_models.APIError
}

func (o *FindUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/users][%d] findUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *FindUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}