// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// CreateOrganizationReader is a Reader for the CreateOrganization structure.
type CreateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewCreateOrganizationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewCreateOrganizationAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewCreateOrganizationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewCreateOrganizationNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateOrganizationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewCreateOrganizationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOrganizationOK creates a CreateOrganizationOK with default headers values
func NewCreateOrganizationOK() *CreateOrganizationOK {
	return &CreateOrganizationOK{}
}

/*CreateOrganizationOK handles this case with default header values.

OK
*/
type CreateOrganizationOK struct {
	Payload *api_models.Organization
}

func (o *CreateOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationOK  %+v", 200, o.Payload)
}

func (o *CreateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationCreated creates a CreateOrganizationCreated with default headers values
func NewCreateOrganizationCreated() *CreateOrganizationCreated {
	return &CreateOrganizationCreated{}
}

/*CreateOrganizationCreated handles this case with default header values.

Created
*/
type CreateOrganizationCreated struct {
}

func (o *CreateOrganizationCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationCreated ", 201)
}

func (o *CreateOrganizationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrganizationAccepted creates a CreateOrganizationAccepted with default headers values
func NewCreateOrganizationAccepted() *CreateOrganizationAccepted {
	return &CreateOrganizationAccepted{}
}

/*CreateOrganizationAccepted handles this case with default header values.

Accepted
*/
type CreateOrganizationAccepted struct {
}

func (o *CreateOrganizationAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationAccepted ", 202)
}

func (o *CreateOrganizationAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrganizationBadRequest creates a CreateOrganizationBadRequest with default headers values
func NewCreateOrganizationBadRequest() *CreateOrganizationBadRequest {
	return &CreateOrganizationBadRequest{}
}

/*CreateOrganizationBadRequest handles this case with default header values.

Bad Request
*/
type CreateOrganizationBadRequest struct {
	Payload *api_models.APIError
}

func (o *CreateOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationUnauthorized creates a CreateOrganizationUnauthorized with default headers values
func NewCreateOrganizationUnauthorized() *CreateOrganizationUnauthorized {
	return &CreateOrganizationUnauthorized{}
}

/*CreateOrganizationUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateOrganizationUnauthorized struct {
	Payload *api_models.APIError
}

func (o *CreateOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationForbidden creates a CreateOrganizationForbidden with default headers values
func NewCreateOrganizationForbidden() *CreateOrganizationForbidden {
	return &CreateOrganizationForbidden{}
}

/*CreateOrganizationForbidden handles this case with default header values.

Forbidden
*/
type CreateOrganizationForbidden struct {
	Payload *api_models.APIError
}

func (o *CreateOrganizationForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *CreateOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationNotFound creates a CreateOrganizationNotFound with default headers values
func NewCreateOrganizationNotFound() *CreateOrganizationNotFound {
	return &CreateOrganizationNotFound{}
}

/*CreateOrganizationNotFound handles this case with default header values.

Not Found
*/
type CreateOrganizationNotFound struct {
	Payload *api_models.APIError
}

func (o *CreateOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *CreateOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationMethodNotAllowed creates a CreateOrganizationMethodNotAllowed with default headers values
func NewCreateOrganizationMethodNotAllowed() *CreateOrganizationMethodNotAllowed {
	return &CreateOrganizationMethodNotAllowed{}
}

/*CreateOrganizationMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type CreateOrganizationMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *CreateOrganizationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CreateOrganizationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationNotAcceptable creates a CreateOrganizationNotAcceptable with default headers values
func NewCreateOrganizationNotAcceptable() *CreateOrganizationNotAcceptable {
	return &CreateOrganizationNotAcceptable{}
}

/*CreateOrganizationNotAcceptable handles this case with default header values.

Not Acceptable
*/
type CreateOrganizationNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *CreateOrganizationNotAcceptable) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationNotAcceptable  %+v", 406, o.Payload)
}

func (o *CreateOrganizationNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationConflict creates a CreateOrganizationConflict with default headers values
func NewCreateOrganizationConflict() *CreateOrganizationConflict {
	return &CreateOrganizationConflict{}
}

/*CreateOrganizationConflict handles this case with default header values.

Conflict
*/
type CreateOrganizationConflict struct {
	Payload *api_models.APIError
}

func (o *CreateOrganizationConflict) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationConflict  %+v", 409, o.Payload)
}

func (o *CreateOrganizationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationUnsupportedMediaType creates a CreateOrganizationUnsupportedMediaType with default headers values
func NewCreateOrganizationUnsupportedMediaType() *CreateOrganizationUnsupportedMediaType {
	return &CreateOrganizationUnsupportedMediaType{}
}

/*CreateOrganizationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type CreateOrganizationUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *CreateOrganizationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *CreateOrganizationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationInternalServerError creates a CreateOrganizationInternalServerError with default headers values
func NewCreateOrganizationInternalServerError() *CreateOrganizationInternalServerError {
	return &CreateOrganizationInternalServerError{}
}

/*CreateOrganizationInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateOrganizationInternalServerError struct {
	Payload *api_models.APIError
}

func (o *CreateOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1/organizations][%d] createOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
