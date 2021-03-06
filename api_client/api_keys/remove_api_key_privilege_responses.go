// Code generated by go-swagger; DO NOT EDIT.

package api_keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// RemoveAPIKeyPrivilegeReader is a Reader for the RemoveAPIKeyPrivilege structure.
type RemoveAPIKeyPrivilegeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAPIKeyPrivilegeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveAPIKeyPrivilegeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveAPIKeyPrivilegeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveAPIKeyPrivilegeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveAPIKeyPrivilegeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveAPIKeyPrivilegeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRemoveAPIKeyPrivilegeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveAPIKeyPrivilegeOK creates a RemoveAPIKeyPrivilegeOK with default headers values
func NewRemoveAPIKeyPrivilegeOK() *RemoveAPIKeyPrivilegeOK {
	return &RemoveAPIKeyPrivilegeOK{}
}

/*RemoveAPIKeyPrivilegeOK handles this case with default header values.

OK
*/
type RemoveAPIKeyPrivilegeOK struct {
}

func (o *RemoveAPIKeyPrivilegeOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges][%d] removeApiKeyPrivilegeOK ", 200)
}

func (o *RemoveAPIKeyPrivilegeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAPIKeyPrivilegeBadRequest creates a RemoveAPIKeyPrivilegeBadRequest with default headers values
func NewRemoveAPIKeyPrivilegeBadRequest() *RemoveAPIKeyPrivilegeBadRequest {
	return &RemoveAPIKeyPrivilegeBadRequest{}
}

/*RemoveAPIKeyPrivilegeBadRequest handles this case with default header values.

Bad Request
*/
type RemoveAPIKeyPrivilegeBadRequest struct {
	Payload *api_models.APIError
}

func (o *RemoveAPIKeyPrivilegeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges][%d] removeApiKeyPrivilegeBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveAPIKeyPrivilegeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAPIKeyPrivilegeUnauthorized creates a RemoveAPIKeyPrivilegeUnauthorized with default headers values
func NewRemoveAPIKeyPrivilegeUnauthorized() *RemoveAPIKeyPrivilegeUnauthorized {
	return &RemoveAPIKeyPrivilegeUnauthorized{}
}

/*RemoveAPIKeyPrivilegeUnauthorized handles this case with default header values.

Unauthorized
*/
type RemoveAPIKeyPrivilegeUnauthorized struct {
	Payload *api_models.APIError
}

func (o *RemoveAPIKeyPrivilegeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges][%d] removeApiKeyPrivilegeUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveAPIKeyPrivilegeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAPIKeyPrivilegeForbidden creates a RemoveAPIKeyPrivilegeForbidden with default headers values
func NewRemoveAPIKeyPrivilegeForbidden() *RemoveAPIKeyPrivilegeForbidden {
	return &RemoveAPIKeyPrivilegeForbidden{}
}

/*RemoveAPIKeyPrivilegeForbidden handles this case with default header values.

Forbidden
*/
type RemoveAPIKeyPrivilegeForbidden struct {
	Payload *api_models.APIError
}

func (o *RemoveAPIKeyPrivilegeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges][%d] removeApiKeyPrivilegeForbidden  %+v", 403, o.Payload)
}

func (o *RemoveAPIKeyPrivilegeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAPIKeyPrivilegeNotFound creates a RemoveAPIKeyPrivilegeNotFound with default headers values
func NewRemoveAPIKeyPrivilegeNotFound() *RemoveAPIKeyPrivilegeNotFound {
	return &RemoveAPIKeyPrivilegeNotFound{}
}

/*RemoveAPIKeyPrivilegeNotFound handles this case with default header values.

Not Found
*/
type RemoveAPIKeyPrivilegeNotFound struct {
	Payload *api_models.APIError
}

func (o *RemoveAPIKeyPrivilegeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges][%d] removeApiKeyPrivilegeNotFound  %+v", 404, o.Payload)
}

func (o *RemoveAPIKeyPrivilegeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAPIKeyPrivilegeInternalServerError creates a RemoveAPIKeyPrivilegeInternalServerError with default headers values
func NewRemoveAPIKeyPrivilegeInternalServerError() *RemoveAPIKeyPrivilegeInternalServerError {
	return &RemoveAPIKeyPrivilegeInternalServerError{}
}

/*RemoveAPIKeyPrivilegeInternalServerError handles this case with default header values.

Internal Server Error
*/
type RemoveAPIKeyPrivilegeInternalServerError struct {
	Payload *api_models.APIError
}

func (o *RemoveAPIKeyPrivilegeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/apikeys/{key}/privileges][%d] removeApiKeyPrivilegeInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveAPIKeyPrivilegeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
