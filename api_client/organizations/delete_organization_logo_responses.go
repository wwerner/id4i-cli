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

// DeleteOrganizationLogoReader is a Reader for the DeleteOrganizationLogo structure.
type DeleteOrganizationLogoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationLogoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteOrganizationLogoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteOrganizationLogoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteOrganizationLogoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteOrganizationLogoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteOrganizationLogoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteOrganizationLogoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOrganizationLogoOK creates a DeleteOrganizationLogoOK with default headers values
func NewDeleteOrganizationLogoOK() *DeleteOrganizationLogoOK {
	return &DeleteOrganizationLogoOK{}
}

/*DeleteOrganizationLogoOK handles this case with default header values.

OK
*/
type DeleteOrganizationLogoOK struct {
}

func (o *DeleteOrganizationLogoOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organizations/{organizationId}/logo][%d] deleteOrganizationLogoOK ", 200)
}

func (o *DeleteOrganizationLogoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationLogoBadRequest creates a DeleteOrganizationLogoBadRequest with default headers values
func NewDeleteOrganizationLogoBadRequest() *DeleteOrganizationLogoBadRequest {
	return &DeleteOrganizationLogoBadRequest{}
}

/*DeleteOrganizationLogoBadRequest handles this case with default header values.

Bad Request
*/
type DeleteOrganizationLogoBadRequest struct {
	Payload *api_models.APIError
}

func (o *DeleteOrganizationLogoBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organizations/{organizationId}/logo][%d] deleteOrganizationLogoBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOrganizationLogoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationLogoUnauthorized creates a DeleteOrganizationLogoUnauthorized with default headers values
func NewDeleteOrganizationLogoUnauthorized() *DeleteOrganizationLogoUnauthorized {
	return &DeleteOrganizationLogoUnauthorized{}
}

/*DeleteOrganizationLogoUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteOrganizationLogoUnauthorized struct {
	Payload *api_models.APIError
}

func (o *DeleteOrganizationLogoUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organizations/{organizationId}/logo][%d] deleteOrganizationLogoUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOrganizationLogoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationLogoForbidden creates a DeleteOrganizationLogoForbidden with default headers values
func NewDeleteOrganizationLogoForbidden() *DeleteOrganizationLogoForbidden {
	return &DeleteOrganizationLogoForbidden{}
}

/*DeleteOrganizationLogoForbidden handles this case with default header values.

Forbidden
*/
type DeleteOrganizationLogoForbidden struct {
	Payload *api_models.APIError
}

func (o *DeleteOrganizationLogoForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organizations/{organizationId}/logo][%d] deleteOrganizationLogoForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrganizationLogoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationLogoNotFound creates a DeleteOrganizationLogoNotFound with default headers values
func NewDeleteOrganizationLogoNotFound() *DeleteOrganizationLogoNotFound {
	return &DeleteOrganizationLogoNotFound{}
}

/*DeleteOrganizationLogoNotFound handles this case with default header values.

Not Found
*/
type DeleteOrganizationLogoNotFound struct {
	Payload *api_models.APIError
}

func (o *DeleteOrganizationLogoNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organizations/{organizationId}/logo][%d] deleteOrganizationLogoNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrganizationLogoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationLogoInternalServerError creates a DeleteOrganizationLogoInternalServerError with default headers values
func NewDeleteOrganizationLogoInternalServerError() *DeleteOrganizationLogoInternalServerError {
	return &DeleteOrganizationLogoInternalServerError{}
}

/*DeleteOrganizationLogoInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteOrganizationLogoInternalServerError struct {
	Payload *api_models.APIError
}

func (o *DeleteOrganizationLogoInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/organizations/{organizationId}/logo][%d] deleteOrganizationLogoInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOrganizationLogoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
