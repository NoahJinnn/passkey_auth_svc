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

// GetSandboxAccessTokenReader is a Reader for the GetSandboxAccessToken structure.
type GetSandboxAccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSandboxAccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSandboxAccessTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSandboxAccessTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSandboxAccessTokenOK creates a GetSandboxAccessTokenOK with default headers values
func NewGetSandboxAccessTokenOK() *GetSandboxAccessTokenOK {
	return &GetSandboxAccessTokenOK{}
}

/*
GetSandboxAccessTokenOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetSandboxAccessTokenOK struct {
	Payload *model.GetAccessTokenResp
}

// IsSuccess returns true when this get sandbox access token o k response has a 2xx status code
func (o *GetSandboxAccessTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get sandbox access token o k response has a 3xx status code
func (o *GetSandboxAccessTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get sandbox access token o k response has a 4xx status code
func (o *GetSandboxAccessTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get sandbox access token o k response has a 5xx status code
func (o *GetSandboxAccessTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get sandbox access token o k response a status code equal to that given
func (o *GetSandboxAccessTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get sandbox access token o k response
func (o *GetSandboxAccessTokenOK) Code() int {
	return 200
}

func (o *GetSandboxAccessTokenOK) Error() string {
	return fmt.Sprintf("[GET /sandbox_access_token/{institution_id}][%d] getSandboxAccessTokenOK  %+v", 200, o.Payload)
}

func (o *GetSandboxAccessTokenOK) String() string {
	return fmt.Sprintf("[GET /sandbox_access_token/{institution_id}][%d] getSandboxAccessTokenOK  %+v", 200, o.Payload)
}

func (o *GetSandboxAccessTokenOK) GetPayload() *model.GetAccessTokenResp {
	return o.Payload
}

func (o *GetSandboxAccessTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.GetAccessTokenResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSandboxAccessTokenDefault creates a GetSandboxAccessTokenDefault with default headers values
func NewGetSandboxAccessTokenDefault(code int) *GetSandboxAccessTokenDefault {
	return &GetSandboxAccessTokenDefault{
		_statusCode: code,
	}
}

/*
GetSandboxAccessTokenDefault describes a response with status code -1, with default header values.

General errors using same model as used by go-swagger for validation errors.
*/
type GetSandboxAccessTokenDefault struct {
	_statusCode int

	Payload *model.Error
}

// IsSuccess returns true when this get sandbox access token default response has a 2xx status code
func (o *GetSandboxAccessTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get sandbox access token default response has a 3xx status code
func (o *GetSandboxAccessTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get sandbox access token default response has a 4xx status code
func (o *GetSandboxAccessTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get sandbox access token default response has a 5xx status code
func (o *GetSandboxAccessTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get sandbox access token default response a status code equal to that given
func (o *GetSandboxAccessTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get sandbox access token default response
func (o *GetSandboxAccessTokenDefault) Code() int {
	return o._statusCode
}

func (o *GetSandboxAccessTokenDefault) Error() string {
	return fmt.Sprintf("[GET /sandbox_access_token/{institution_id}][%d] GetSandboxAccessToken default  %+v", o._statusCode, o.Payload)
}

func (o *GetSandboxAccessTokenDefault) String() string {
	return fmt.Sprintf("[GET /sandbox_access_token/{institution_id}][%d] GetSandboxAccessToken default  %+v", o._statusCode, o.Payload)
}

func (o *GetSandboxAccessTokenDefault) GetPayload() *model.Error {
	return o.Payload
}

func (o *GetSandboxAccessTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
