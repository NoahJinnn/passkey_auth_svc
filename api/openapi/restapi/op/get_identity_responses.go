// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hellohq/hqservice/api/openapi/model"
)

// GetIdentityOKCode is the HTTP code returned for type GetIdentityOK
const GetIdentityOKCode int = 200

/*
GetIdentityOK A successful response.

swagger:response getIdentityOK
*/
type GetIdentityOK struct {

	/*
	  In: Body
	*/
	Payload *model.GetIdentityResp `json:"body,omitempty"`
}

// NewGetIdentityOK creates GetIdentityOK with default headers values
func NewGetIdentityOK() *GetIdentityOK {

	return &GetIdentityOK{}
}

// WithPayload adds the payload to the get identity o k response
func (o *GetIdentityOK) WithPayload(payload *model.GetIdentityResp) *GetIdentityOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity o k response
func (o *GetIdentityOK) SetPayload(payload *model.GetIdentityResp) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetIdentityDefault General errors using same model as used by go-swagger for validation errors.

swagger:response getIdentityDefault
*/
type GetIdentityDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *model.Error `json:"body,omitempty"`
}

// NewGetIdentityDefault creates GetIdentityDefault with default headers values
func NewGetIdentityDefault(code int) *GetIdentityDefault {
	if code <= 0 {
		code = 500
	}

	return &GetIdentityDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get identity default response
func (o *GetIdentityDefault) WithStatusCode(code int) *GetIdentityDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get identity default response
func (o *GetIdentityDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get identity default response
func (o *GetIdentityDefault) WithPayload(payload *model.Error) *GetIdentityDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get identity default response
func (o *GetIdentityDefault) SetPayload(payload *model.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetIdentityDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
