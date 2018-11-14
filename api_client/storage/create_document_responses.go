// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// CreateDocumentReader is a Reader for the CreateDocument structure.
type CreateDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewCreateDocumentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateDocumentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateDocumentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewCreateDocumentMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewCreateDocumentNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateDocumentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewCreateDocumentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateDocumentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDocumentOK creates a CreateDocumentOK with default headers values
func NewCreateDocumentOK() *CreateDocumentOK {
	return &CreateDocumentOK{}
}

/*CreateDocumentOK handles this case with default header values.

OK
*/
type CreateDocumentOK struct {
	Payload *api_models.Document
}

func (o *CreateDocumentOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentOK  %+v", 200, o.Payload)
}

func (o *CreateDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.Document)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDocumentAccepted creates a CreateDocumentAccepted with default headers values
func NewCreateDocumentAccepted() *CreateDocumentAccepted {
	return &CreateDocumentAccepted{}
}

/*CreateDocumentAccepted handles this case with default header values.

Accepted
*/
type CreateDocumentAccepted struct {
}

func (o *CreateDocumentAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentAccepted ", 202)
}

func (o *CreateDocumentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateDocumentBadRequest creates a CreateDocumentBadRequest with default headers values
func NewCreateDocumentBadRequest() *CreateDocumentBadRequest {
	return &CreateDocumentBadRequest{}
}

/*CreateDocumentBadRequest handles this case with default header values.

Bad Request
*/
type CreateDocumentBadRequest struct {
	Payload *api_models.APIError
}

func (o *CreateDocumentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDocumentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDocumentUnauthorized creates a CreateDocumentUnauthorized with default headers values
func NewCreateDocumentUnauthorized() *CreateDocumentUnauthorized {
	return &CreateDocumentUnauthorized{}
}

/*CreateDocumentUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateDocumentUnauthorized struct {
	Payload *api_models.APIError
}

func (o *CreateDocumentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateDocumentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDocumentForbidden creates a CreateDocumentForbidden with default headers values
func NewCreateDocumentForbidden() *CreateDocumentForbidden {
	return &CreateDocumentForbidden{}
}

/*CreateDocumentForbidden handles this case with default header values.

Forbidden
*/
type CreateDocumentForbidden struct {
	Payload *api_models.APIError
}

func (o *CreateDocumentForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentForbidden  %+v", 403, o.Payload)
}

func (o *CreateDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDocumentNotFound creates a CreateDocumentNotFound with default headers values
func NewCreateDocumentNotFound() *CreateDocumentNotFound {
	return &CreateDocumentNotFound{}
}

/*CreateDocumentNotFound handles this case with default header values.

Not Found
*/
type CreateDocumentNotFound struct {
	Payload *api_models.APIError
}

func (o *CreateDocumentNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentNotFound  %+v", 404, o.Payload)
}

func (o *CreateDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDocumentMethodNotAllowed creates a CreateDocumentMethodNotAllowed with default headers values
func NewCreateDocumentMethodNotAllowed() *CreateDocumentMethodNotAllowed {
	return &CreateDocumentMethodNotAllowed{}
}

/*CreateDocumentMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type CreateDocumentMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *CreateDocumentMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CreateDocumentMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDocumentNotAcceptable creates a CreateDocumentNotAcceptable with default headers values
func NewCreateDocumentNotAcceptable() *CreateDocumentNotAcceptable {
	return &CreateDocumentNotAcceptable{}
}

/*CreateDocumentNotAcceptable handles this case with default header values.

Not Acceptable
*/
type CreateDocumentNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *CreateDocumentNotAcceptable) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentNotAcceptable  %+v", 406, o.Payload)
}

func (o *CreateDocumentNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDocumentConflict creates a CreateDocumentConflict with default headers values
func NewCreateDocumentConflict() *CreateDocumentConflict {
	return &CreateDocumentConflict{}
}

/*CreateDocumentConflict handles this case with default header values.

Conflict
*/
type CreateDocumentConflict struct {
	Payload *api_models.APIError
}

func (o *CreateDocumentConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentConflict  %+v", 409, o.Payload)
}

func (o *CreateDocumentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDocumentUnsupportedMediaType creates a CreateDocumentUnsupportedMediaType with default headers values
func NewCreateDocumentUnsupportedMediaType() *CreateDocumentUnsupportedMediaType {
	return &CreateDocumentUnsupportedMediaType{}
}

/*CreateDocumentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type CreateDocumentUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *CreateDocumentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateDocumentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDocumentInternalServerError creates a CreateDocumentInternalServerError with default headers values
func NewCreateDocumentInternalServerError() *CreateDocumentInternalServerError {
	return &CreateDocumentInternalServerError{}
}

/*CreateDocumentInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateDocumentInternalServerError struct {
	Payload *api_models.APIError
}

func (o *CreateDocumentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v1/documents/{id4n}/{organizationId}][%d] createDocumentInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateDocumentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
