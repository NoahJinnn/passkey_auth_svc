// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hellohq/hqservice/api/openapi/model"
)

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateUserOK creates a CreateUserOK with default headers values
func NewCreateUserOK() *CreateUserOK {
	return &CreateUserOK{}
}

/*
CreateUserOK describes a response with status code 200, with default header values.

Create a user successfully.
*/
type CreateUserOK struct {
	Payload *model.LinkTokenCreateResp
}

// IsSuccess returns true when this create user o k response has a 2xx status code
func (o *CreateUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user o k response has a 3xx status code
func (o *CreateUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user o k response has a 4xx status code
func (o *CreateUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user o k response has a 5xx status code
func (o *CreateUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create user o k response a status code equal to that given
func (o *CreateUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create user o k response
func (o *CreateUserOK) Code() int {
	return 200
}

func (o *CreateUserOK) Error() string {
	return fmt.Sprintf("[POST /user][%d] createUserOK  %+v", 200, o.Payload)
}

func (o *CreateUserOK) String() string {
	return fmt.Sprintf("[POST /user][%d] createUserOK  %+v", 200, o.Payload)
}

func (o *CreateUserOK) GetPayload() *model.LinkTokenCreateResp {
	return o.Payload
}

func (o *CreateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.LinkTokenCreateResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserDefault creates a CreateUserDefault with default headers values
func NewCreateUserDefault(code int) *CreateUserDefault {
	return &CreateUserDefault{
		_statusCode: code,
	}
}

/*
CreateUserDefault describes a response with status code -1, with default header values.

General errors using same model as used by go-swagger for validation errors.
*/
type CreateUserDefault struct {
	_statusCode int

	Payload *model.Error
}

// IsSuccess returns true when this create user default response has a 2xx status code
func (o *CreateUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create user default response has a 3xx status code
func (o *CreateUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create user default response has a 4xx status code
func (o *CreateUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create user default response has a 5xx status code
func (o *CreateUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create user default response a status code equal to that given
func (o *CreateUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create user default response
func (o *CreateUserDefault) Code() int {
	return o._statusCode
}

func (o *CreateUserDefault) Error() string {
	return fmt.Sprintf("[POST /user][%d] CreateUser default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserDefault) String() string {
	return fmt.Sprintf("[POST /user][%d] CreateUser default  %+v", o._statusCode, o.Payload)
}

func (o *CreateUserDefault) GetPayload() *model.Error {
	return o.Payload
}

func (o *CreateUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
