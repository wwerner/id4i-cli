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

// VerifyPasswordResetReader is a Reader for the VerifyPasswordReset structure.
type VerifyPasswordResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyPasswordResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVerifyPasswordResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewVerifyPasswordResetAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewVerifyPasswordResetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewVerifyPasswordResetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewVerifyPasswordResetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewVerifyPasswordResetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewVerifyPasswordResetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewVerifyPasswordResetNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewVerifyPasswordResetConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewVerifyPasswordResetUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewVerifyPasswordResetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVerifyPasswordResetOK creates a VerifyPasswordResetOK with default headers values
func NewVerifyPasswordResetOK() *VerifyPasswordResetOK {
	return &VerifyPasswordResetOK{}
}

/*VerifyPasswordResetOK handles this case with default header values.

OK
*/
type VerifyPasswordResetOK struct {
	Payload *api_models.SimpleMessageResponse
}

func (o *VerifyPasswordResetOK) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetOK  %+v", 200, o.Payload)
}

func (o *VerifyPasswordResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.SimpleMessageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyPasswordResetAccepted creates a VerifyPasswordResetAccepted with default headers values
func NewVerifyPasswordResetAccepted() *VerifyPasswordResetAccepted {
	return &VerifyPasswordResetAccepted{}
}

/*VerifyPasswordResetAccepted handles this case with default header values.

Accepted
*/
type VerifyPasswordResetAccepted struct {
}

func (o *VerifyPasswordResetAccepted) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetAccepted ", 202)
}

func (o *VerifyPasswordResetAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVerifyPasswordResetBadRequest creates a VerifyPasswordResetBadRequest with default headers values
func NewVerifyPasswordResetBadRequest() *VerifyPasswordResetBadRequest {
	return &VerifyPasswordResetBadRequest{}
}

/*VerifyPasswordResetBadRequest handles this case with default header values.

Bad Request
*/
type VerifyPasswordResetBadRequest struct {
	Payload *api_models.APIError
}

func (o *VerifyPasswordResetBadRequest) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetBadRequest  %+v", 400, o.Payload)
}

func (o *VerifyPasswordResetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyPasswordResetUnauthorized creates a VerifyPasswordResetUnauthorized with default headers values
func NewVerifyPasswordResetUnauthorized() *VerifyPasswordResetUnauthorized {
	return &VerifyPasswordResetUnauthorized{}
}

/*VerifyPasswordResetUnauthorized handles this case with default header values.

Unauthorized
*/
type VerifyPasswordResetUnauthorized struct {
	Payload *api_models.APIError
}

func (o *VerifyPasswordResetUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetUnauthorized  %+v", 401, o.Payload)
}

func (o *VerifyPasswordResetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyPasswordResetForbidden creates a VerifyPasswordResetForbidden with default headers values
func NewVerifyPasswordResetForbidden() *VerifyPasswordResetForbidden {
	return &VerifyPasswordResetForbidden{}
}

/*VerifyPasswordResetForbidden handles this case with default header values.

Forbidden
*/
type VerifyPasswordResetForbidden struct {
	Payload *api_models.APIError
}

func (o *VerifyPasswordResetForbidden) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetForbidden  %+v", 403, o.Payload)
}

func (o *VerifyPasswordResetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyPasswordResetNotFound creates a VerifyPasswordResetNotFound with default headers values
func NewVerifyPasswordResetNotFound() *VerifyPasswordResetNotFound {
	return &VerifyPasswordResetNotFound{}
}

/*VerifyPasswordResetNotFound handles this case with default header values.

Not Found
*/
type VerifyPasswordResetNotFound struct {
	Payload *api_models.APIError
}

func (o *VerifyPasswordResetNotFound) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetNotFound  %+v", 404, o.Payload)
}

func (o *VerifyPasswordResetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyPasswordResetMethodNotAllowed creates a VerifyPasswordResetMethodNotAllowed with default headers values
func NewVerifyPasswordResetMethodNotAllowed() *VerifyPasswordResetMethodNotAllowed {
	return &VerifyPasswordResetMethodNotAllowed{}
}

/*VerifyPasswordResetMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type VerifyPasswordResetMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *VerifyPasswordResetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *VerifyPasswordResetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyPasswordResetNotAcceptable creates a VerifyPasswordResetNotAcceptable with default headers values
func NewVerifyPasswordResetNotAcceptable() *VerifyPasswordResetNotAcceptable {
	return &VerifyPasswordResetNotAcceptable{}
}

/*VerifyPasswordResetNotAcceptable handles this case with default header values.

Not Acceptable
*/
type VerifyPasswordResetNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *VerifyPasswordResetNotAcceptable) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetNotAcceptable  %+v", 406, o.Payload)
}

func (o *VerifyPasswordResetNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyPasswordResetConflict creates a VerifyPasswordResetConflict with default headers values
func NewVerifyPasswordResetConflict() *VerifyPasswordResetConflict {
	return &VerifyPasswordResetConflict{}
}

/*VerifyPasswordResetConflict handles this case with default header values.

Conflict
*/
type VerifyPasswordResetConflict struct {
	Payload *api_models.APIError
}

func (o *VerifyPasswordResetConflict) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetConflict  %+v", 409, o.Payload)
}

func (o *VerifyPasswordResetConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyPasswordResetUnsupportedMediaType creates a VerifyPasswordResetUnsupportedMediaType with default headers values
func NewVerifyPasswordResetUnsupportedMediaType() *VerifyPasswordResetUnsupportedMediaType {
	return &VerifyPasswordResetUnsupportedMediaType{}
}

/*VerifyPasswordResetUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type VerifyPasswordResetUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *VerifyPasswordResetUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *VerifyPasswordResetUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVerifyPasswordResetInternalServerError creates a VerifyPasswordResetInternalServerError with default headers values
func NewVerifyPasswordResetInternalServerError() *VerifyPasswordResetInternalServerError {
	return &VerifyPasswordResetInternalServerError{}
}

/*VerifyPasswordResetInternalServerError handles this case with default header values.

Internal Server Error
*/
type VerifyPasswordResetInternalServerError struct {
	Payload *api_models.APIError
}

func (o *VerifyPasswordResetInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /account/password][%d] verifyPasswordResetInternalServerError  %+v", 500, o.Payload)
}

func (o *VerifyPasswordResetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
