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

// CompleteRegistrationReader is a Reader for the CompleteRegistration structure.
type CompleteRegistrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteRegistrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCompleteRegistrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewCompleteRegistrationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCompleteRegistrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCompleteRegistrationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCompleteRegistrationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCompleteRegistrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewCompleteRegistrationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewCompleteRegistrationNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCompleteRegistrationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewCompleteRegistrationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCompleteRegistrationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompleteRegistrationOK creates a CompleteRegistrationOK with default headers values
func NewCompleteRegistrationOK() *CompleteRegistrationOK {
	return &CompleteRegistrationOK{}
}

/*CompleteRegistrationOK handles this case with default header values.

OK
*/
type CompleteRegistrationOK struct {
}

func (o *CompleteRegistrationOK) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationOK ", 200)
}

func (o *CompleteRegistrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteRegistrationAccepted creates a CompleteRegistrationAccepted with default headers values
func NewCompleteRegistrationAccepted() *CompleteRegistrationAccepted {
	return &CompleteRegistrationAccepted{}
}

/*CompleteRegistrationAccepted handles this case with default header values.

Accepted
*/
type CompleteRegistrationAccepted struct {
}

func (o *CompleteRegistrationAccepted) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationAccepted ", 202)
}

func (o *CompleteRegistrationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteRegistrationBadRequest creates a CompleteRegistrationBadRequest with default headers values
func NewCompleteRegistrationBadRequest() *CompleteRegistrationBadRequest {
	return &CompleteRegistrationBadRequest{}
}

/*CompleteRegistrationBadRequest handles this case with default header values.

Bad Request
*/
type CompleteRegistrationBadRequest struct {
	Payload *api_models.APIError
}

func (o *CompleteRegistrationBadRequest) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationBadRequest  %+v", 400, o.Payload)
}

func (o *CompleteRegistrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteRegistrationUnauthorized creates a CompleteRegistrationUnauthorized with default headers values
func NewCompleteRegistrationUnauthorized() *CompleteRegistrationUnauthorized {
	return &CompleteRegistrationUnauthorized{}
}

/*CompleteRegistrationUnauthorized handles this case with default header values.

Unauthorized
*/
type CompleteRegistrationUnauthorized struct {
	Payload *api_models.APIError
}

func (o *CompleteRegistrationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationUnauthorized  %+v", 401, o.Payload)
}

func (o *CompleteRegistrationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteRegistrationForbidden creates a CompleteRegistrationForbidden with default headers values
func NewCompleteRegistrationForbidden() *CompleteRegistrationForbidden {
	return &CompleteRegistrationForbidden{}
}

/*CompleteRegistrationForbidden handles this case with default header values.

Forbidden
*/
type CompleteRegistrationForbidden struct {
	Payload *api_models.APIError
}

func (o *CompleteRegistrationForbidden) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationForbidden  %+v", 403, o.Payload)
}

func (o *CompleteRegistrationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteRegistrationNotFound creates a CompleteRegistrationNotFound with default headers values
func NewCompleteRegistrationNotFound() *CompleteRegistrationNotFound {
	return &CompleteRegistrationNotFound{}
}

/*CompleteRegistrationNotFound handles this case with default header values.

Not Found
*/
type CompleteRegistrationNotFound struct {
	Payload *api_models.APIError
}

func (o *CompleteRegistrationNotFound) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationNotFound  %+v", 404, o.Payload)
}

func (o *CompleteRegistrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteRegistrationMethodNotAllowed creates a CompleteRegistrationMethodNotAllowed with default headers values
func NewCompleteRegistrationMethodNotAllowed() *CompleteRegistrationMethodNotAllowed {
	return &CompleteRegistrationMethodNotAllowed{}
}

/*CompleteRegistrationMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type CompleteRegistrationMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *CompleteRegistrationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CompleteRegistrationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteRegistrationNotAcceptable creates a CompleteRegistrationNotAcceptable with default headers values
func NewCompleteRegistrationNotAcceptable() *CompleteRegistrationNotAcceptable {
	return &CompleteRegistrationNotAcceptable{}
}

/*CompleteRegistrationNotAcceptable handles this case with default header values.

Not Acceptable
*/
type CompleteRegistrationNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *CompleteRegistrationNotAcceptable) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationNotAcceptable  %+v", 406, o.Payload)
}

func (o *CompleteRegistrationNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteRegistrationConflict creates a CompleteRegistrationConflict with default headers values
func NewCompleteRegistrationConflict() *CompleteRegistrationConflict {
	return &CompleteRegistrationConflict{}
}

/*CompleteRegistrationConflict handles this case with default header values.

Conflict
*/
type CompleteRegistrationConflict struct {
	Payload *api_models.APIError
}

func (o *CompleteRegistrationConflict) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationConflict  %+v", 409, o.Payload)
}

func (o *CompleteRegistrationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteRegistrationUnsupportedMediaType creates a CompleteRegistrationUnsupportedMediaType with default headers values
func NewCompleteRegistrationUnsupportedMediaType() *CompleteRegistrationUnsupportedMediaType {
	return &CompleteRegistrationUnsupportedMediaType{}
}

/*CompleteRegistrationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type CompleteRegistrationUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *CompleteRegistrationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CompleteRegistrationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteRegistrationInternalServerError creates a CompleteRegistrationInternalServerError with default headers values
func NewCompleteRegistrationInternalServerError() *CompleteRegistrationInternalServerError {
	return &CompleteRegistrationInternalServerError{}
}

/*CompleteRegistrationInternalServerError handles this case with default header values.

Internal Server Error
*/
type CompleteRegistrationInternalServerError struct {
	Payload *api_models.APIError
}

func (o *CompleteRegistrationInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /account/registration][%d] completeRegistrationInternalServerError  %+v", 500, o.Payload)
}

func (o *CompleteRegistrationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
