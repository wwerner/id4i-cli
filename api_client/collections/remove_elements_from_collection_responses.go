// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	api_models "github.com/BlueRainSoftware/id4i-cli/api_models"
)

// RemoveElementsFromCollectionReader is a Reader for the RemoveElementsFromCollection structure.
type RemoveElementsFromCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveElementsFromCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveElementsFromCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRemoveElementsFromCollectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewRemoveElementsFromCollectionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveElementsFromCollectionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRemoveElementsFromCollectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRemoveElementsFromCollectionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveElementsFromCollectionOK creates a RemoveElementsFromCollectionOK with default headers values
func NewRemoveElementsFromCollectionOK() *RemoveElementsFromCollectionOK {
	return &RemoveElementsFromCollectionOK{}
}

/*RemoveElementsFromCollectionOK handles this case with default header values.

OK
*/
type RemoveElementsFromCollectionOK struct {
}

func (o *RemoveElementsFromCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/collections/{id4n}/elements][%d] removeElementsFromCollectionOK ", 200)
}

func (o *RemoveElementsFromCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveElementsFromCollectionBadRequest creates a RemoveElementsFromCollectionBadRequest with default headers values
func NewRemoveElementsFromCollectionBadRequest() *RemoveElementsFromCollectionBadRequest {
	return &RemoveElementsFromCollectionBadRequest{}
}

/*RemoveElementsFromCollectionBadRequest handles this case with default header values.

Bad Request
*/
type RemoveElementsFromCollectionBadRequest struct {
	Payload *api_models.APIError
}

func (o *RemoveElementsFromCollectionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/collections/{id4n}/elements][%d] removeElementsFromCollectionBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveElementsFromCollectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveElementsFromCollectionUnauthorized creates a RemoveElementsFromCollectionUnauthorized with default headers values
func NewRemoveElementsFromCollectionUnauthorized() *RemoveElementsFromCollectionUnauthorized {
	return &RemoveElementsFromCollectionUnauthorized{}
}

/*RemoveElementsFromCollectionUnauthorized handles this case with default header values.

Unauthorized
*/
type RemoveElementsFromCollectionUnauthorized struct {
	Payload *api_models.APIError
}

func (o *RemoveElementsFromCollectionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/collections/{id4n}/elements][%d] removeElementsFromCollectionUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveElementsFromCollectionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveElementsFromCollectionForbidden creates a RemoveElementsFromCollectionForbidden with default headers values
func NewRemoveElementsFromCollectionForbidden() *RemoveElementsFromCollectionForbidden {
	return &RemoveElementsFromCollectionForbidden{}
}

/*RemoveElementsFromCollectionForbidden handles this case with default header values.

Forbidden
*/
type RemoveElementsFromCollectionForbidden struct {
	Payload *api_models.APIError
}

func (o *RemoveElementsFromCollectionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/collections/{id4n}/elements][%d] removeElementsFromCollectionForbidden  %+v", 403, o.Payload)
}

func (o *RemoveElementsFromCollectionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveElementsFromCollectionNotFound creates a RemoveElementsFromCollectionNotFound with default headers values
func NewRemoveElementsFromCollectionNotFound() *RemoveElementsFromCollectionNotFound {
	return &RemoveElementsFromCollectionNotFound{}
}

/*RemoveElementsFromCollectionNotFound handles this case with default header values.

Not Found
*/
type RemoveElementsFromCollectionNotFound struct {
	Payload *api_models.APIError
}

func (o *RemoveElementsFromCollectionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/collections/{id4n}/elements][%d] removeElementsFromCollectionNotFound  %+v", 404, o.Payload)
}

func (o *RemoveElementsFromCollectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveElementsFromCollectionInternalServerError creates a RemoveElementsFromCollectionInternalServerError with default headers values
func NewRemoveElementsFromCollectionInternalServerError() *RemoveElementsFromCollectionInternalServerError {
	return &RemoveElementsFromCollectionInternalServerError{}
}

/*RemoveElementsFromCollectionInternalServerError handles this case with default header values.

Internal Server Error
*/
type RemoveElementsFromCollectionInternalServerError struct {
	Payload *api_models.APIError
}

func (o *RemoveElementsFromCollectionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/collections/{id4n}/elements][%d] removeElementsFromCollectionInternalServerError  %+v", 500, o.Payload)
}

func (o *RemoveElementsFromCollectionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
