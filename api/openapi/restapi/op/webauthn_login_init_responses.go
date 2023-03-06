// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/hellohq/hqservice/api/openapi/model"
)

// WebauthnLoginInitOKCode is the HTTP code returned for type WebauthnLoginInitOK
const WebauthnLoginInitOKCode int = 200

/*
WebauthnLoginInitOK Successful initialization

swagger:response webauthnLoginInitOK
*/
type WebauthnLoginInitOK struct {

	/*
	  In: Body
	*/
	Payload *model.CredentialRequestOptions `json:"body,omitempty"`
}

// NewWebauthnLoginInitOK creates WebauthnLoginInitOK with default headers values
func NewWebauthnLoginInitOK() *WebauthnLoginInitOK {

	return &WebauthnLoginInitOK{}
}

// WithPayload adds the payload to the webauthn login init o k response
func (o *WebauthnLoginInitOK) WithPayload(payload *model.CredentialRequestOptions) *WebauthnLoginInitOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the webauthn login init o k response
func (o *WebauthnLoginInitOK) SetPayload(payload *model.CredentialRequestOptions) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WebauthnLoginInitOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
