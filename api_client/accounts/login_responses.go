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

// LoginReader is a Reader for the Login structure.
type LoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoginOK creates a LoginOK with default headers values
func NewLoginOK() *LoginOK {
	return &LoginOK{}
}

/*LoginOK handles this case with default header values.

OK
*/
type LoginOK struct {
}

func (o *LoginOK) Error() string {
	return fmt.Sprintf("[POST /login][%d] loginOK ", 200)
}

func (o *LoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewLoginUnauthorized creates a LoginUnauthorized with default headers values
func NewLoginUnauthorized() *LoginUnauthorized {
	return &LoginUnauthorized{}
}

/*LoginUnauthorized handles this case with default header values.

Unauthorized
*/
type LoginUnauthorized struct {
	Payload *api_models.APIError
}

func (o *LoginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /login][%d] loginUnauthorized  %+v", 401, o.Payload)
}

func (o *LoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(api_models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}