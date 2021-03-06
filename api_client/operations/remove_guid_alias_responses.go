// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// RemoveGUIDAliasReader is a Reader for the RemoveGUIDAlias structure.
type RemoveGUIDAliasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveGUIDAliasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveGUIDAliasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveGUIDAliasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveGUIDAliasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveGUIDAliasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveGUIDAliasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRemoveGUIDAliasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveGUIDAliasOK creates a RemoveGUIDAliasOK with default headers values
func NewRemoveGUIDAliasOK() *RemoveGUIDAliasOK {
	return &RemoveGUIDAliasOK{}
}

/*RemoveGUIDAliasOK handles this case with default header values.

OK
*/
type RemoveGUIDAliasOK struct {
}

func (o *RemoveGUIDAliasOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] removeGuidAliasOK ", 200)
}

func (o *RemoveGUIDAliasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveGUIDAliasBadRequest creates a RemoveGUIDAliasBadRequest with default headers values
func NewRemoveGUIDAliasBadRequest() *RemoveGUIDAliasBadRequest {
	return &RemoveGUIDAliasBadRequest{}
}

/*RemoveGUIDAliasBadRequest handles this case with default header values.

Bad Request
*/
type RemoveGUIDAliasBadRequest struct {
	Payload *api_models.APIError
}

func (o *RemoveGUIDAliasBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] removeGuidAliasBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveGUIDAliasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveGUIDAliasUnauthorized creates a RemoveGUIDAliasUnauthorized with default headers values
func NewRemoveGUIDAliasUnauthorized() *RemoveGUIDAliasUnauthorized {
	return &RemoveGUIDAliasUnauthorized{}
}

/*RemoveGUIDAliasUnauthorized handles this case with default header values.

Unauthorized
*/
type RemoveGUIDAliasUnauthorized struct {
	Payload *api_models.APIError
}

func (o *RemoveGUIDAliasUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] removeGuidAliasUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveGUIDAliasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveGUIDAliasForbidden creates a RemoveGUIDAliasForbidden with default headers values
func NewRemoveGUIDAliasForbidden() *RemoveGUIDAliasForbidden {
	return &RemoveGUIDAliasForbidden{}
}

/*RemoveGUIDAliasForbidden handles this case with default header values.

Forbidden
*/
type RemoveGUIDAliasForbidden struct {
	Payload *api_models.APIError
}

func (o *RemoveGUIDAliasForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] removeGuidAliasForbidden  %+v", 403, o.Payload)
}

func (o *RemoveGUIDAliasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveGUIDAliasNotFound creates a RemoveGUIDAliasNotFound with default headers values
func NewRemoveGUIDAliasNotFound() *RemoveGUIDAliasNotFound {
	return &RemoveGUIDAliasNotFound{}
}

/*RemoveGUIDAliasNotFound handles this case with default header values.

Not Found
*/
type RemoveGUIDAliasNotFound struct {
	Payload *api_models.APIError
}

func (o *RemoveGUIDAliasNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] removeGuidAliasNotFound  %+v", 404, o.Payload)
}

func (o *RemoveGUIDAliasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveGUIDAliasInternalServerError creates a RemoveGUIDAliasInternalServerError with default headers values
func NewRemoveGUIDAliasInternalServerError() *RemoveGUIDAliasInternalServerError {
	return &RemoveGUIDAliasInternalServerError{}
}

/*RemoveGUIDAliasInternalServerError handles this case with default header values.

Internal Server Error
*/
type RemoveGUIDAliasInternalServerError struct {
	Payload *api_models.APIError
}

func (o *RemoveGUIDAliasInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/id4ns/{id4n}/alias/{aliasType}][%d] removeGuidAliasInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveGUIDAliasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
