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

// ListAllPublicDocumentsReader is a Reader for the ListAllPublicDocuments structure.
type ListAllPublicDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAllPublicDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAllPublicDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewListAllPublicDocumentsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewListAllPublicDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListAllPublicDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListAllPublicDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewListAllPublicDocumentsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewListAllPublicDocumentsNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewListAllPublicDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListAllPublicDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAllPublicDocumentsOK creates a ListAllPublicDocumentsOK with default headers values
func NewListAllPublicDocumentsOK() *ListAllPublicDocumentsOK {
	return &ListAllPublicDocumentsOK{}
}

/*ListAllPublicDocumentsOK handles this case with default header values.

OK
*/
type ListAllPublicDocumentsOK struct {
	Payload *api_models.PaginatedDocumentResponse
}

func (o *ListAllPublicDocumentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/documents/{id4n}][%d] listAllPublicDocumentsOK  %+v", 200, o.Payload)
}

func (o *ListAllPublicDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.PaginatedDocumentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllPublicDocumentsAccepted creates a ListAllPublicDocumentsAccepted with default headers values
func NewListAllPublicDocumentsAccepted() *ListAllPublicDocumentsAccepted {
	return &ListAllPublicDocumentsAccepted{}
}

/*ListAllPublicDocumentsAccepted handles this case with default header values.

Accepted
*/
type ListAllPublicDocumentsAccepted struct {
}

func (o *ListAllPublicDocumentsAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/documents/{id4n}][%d] listAllPublicDocumentsAccepted ", 202)
}

func (o *ListAllPublicDocumentsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListAllPublicDocumentsUnauthorized creates a ListAllPublicDocumentsUnauthorized with default headers values
func NewListAllPublicDocumentsUnauthorized() *ListAllPublicDocumentsUnauthorized {
	return &ListAllPublicDocumentsUnauthorized{}
}

/*ListAllPublicDocumentsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAllPublicDocumentsUnauthorized struct {
	Payload *api_models.APIError
}

func (o *ListAllPublicDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/documents/{id4n}][%d] listAllPublicDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAllPublicDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllPublicDocumentsForbidden creates a ListAllPublicDocumentsForbidden with default headers values
func NewListAllPublicDocumentsForbidden() *ListAllPublicDocumentsForbidden {
	return &ListAllPublicDocumentsForbidden{}
}

/*ListAllPublicDocumentsForbidden handles this case with default header values.

Forbidden
*/
type ListAllPublicDocumentsForbidden struct {
	Payload *api_models.APIError
}

func (o *ListAllPublicDocumentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/documents/{id4n}][%d] listAllPublicDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *ListAllPublicDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllPublicDocumentsNotFound creates a ListAllPublicDocumentsNotFound with default headers values
func NewListAllPublicDocumentsNotFound() *ListAllPublicDocumentsNotFound {
	return &ListAllPublicDocumentsNotFound{}
}

/*ListAllPublicDocumentsNotFound handles this case with default header values.

Not Found
*/
type ListAllPublicDocumentsNotFound struct {
	Payload *api_models.APIError
}

func (o *ListAllPublicDocumentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/documents/{id4n}][%d] listAllPublicDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *ListAllPublicDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllPublicDocumentsMethodNotAllowed creates a ListAllPublicDocumentsMethodNotAllowed with default headers values
func NewListAllPublicDocumentsMethodNotAllowed() *ListAllPublicDocumentsMethodNotAllowed {
	return &ListAllPublicDocumentsMethodNotAllowed{}
}

/*ListAllPublicDocumentsMethodNotAllowed handles this case with default header values.

Method not allowed
*/
type ListAllPublicDocumentsMethodNotAllowed struct {
	Payload *api_models.APIError
}

func (o *ListAllPublicDocumentsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/documents/{id4n}][%d] listAllPublicDocumentsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *ListAllPublicDocumentsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllPublicDocumentsNotAcceptable creates a ListAllPublicDocumentsNotAcceptable with default headers values
func NewListAllPublicDocumentsNotAcceptable() *ListAllPublicDocumentsNotAcceptable {
	return &ListAllPublicDocumentsNotAcceptable{}
}

/*ListAllPublicDocumentsNotAcceptable handles this case with default header values.

Not Acceptable
*/
type ListAllPublicDocumentsNotAcceptable struct {
	Payload *api_models.APIError
}

func (o *ListAllPublicDocumentsNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/documents/{id4n}][%d] listAllPublicDocumentsNotAcceptable  %+v", 406, o.Payload)
}

func (o *ListAllPublicDocumentsNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllPublicDocumentsUnsupportedMediaType creates a ListAllPublicDocumentsUnsupportedMediaType with default headers values
func NewListAllPublicDocumentsUnsupportedMediaType() *ListAllPublicDocumentsUnsupportedMediaType {
	return &ListAllPublicDocumentsUnsupportedMediaType{}
}

/*ListAllPublicDocumentsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type
*/
type ListAllPublicDocumentsUnsupportedMediaType struct {
	Payload *api_models.APIError
}

func (o *ListAllPublicDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/documents/{id4n}][%d] listAllPublicDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *ListAllPublicDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAllPublicDocumentsInternalServerError creates a ListAllPublicDocumentsInternalServerError with default headers values
func NewListAllPublicDocumentsInternalServerError() *ListAllPublicDocumentsInternalServerError {
	return &ListAllPublicDocumentsInternalServerError{}
}

/*ListAllPublicDocumentsInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAllPublicDocumentsInternalServerError struct {
	Payload *api_models.APIError
}

func (o *ListAllPublicDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/public/documents/{id4n}][%d] listAllPublicDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAllPublicDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
