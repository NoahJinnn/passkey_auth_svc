// Code generated by go-swagger; DO NOT EDIT.

package web_authn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/hellohq/hqservice/api/openapi/model"
)

// WebauthnLoginFinalReader is a Reader for the WebauthnLoginFinal structure.
type WebauthnLoginFinalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebauthnLoginFinalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebauthnLoginFinalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWebauthnLoginFinalOK creates a WebauthnLoginFinalOK with default headers values
func NewWebauthnLoginFinalOK() *WebauthnLoginFinalOK {
	return &WebauthnLoginFinalOK{}
}

/*
WebauthnLoginFinalOK describes a response with status code 200, with default header values.

Successful login
*/
type WebauthnLoginFinalOK struct {

	/* Contains the JSON Web Token (JWT) that must be provided to protected endpoints.
	Cookie attributes (e.g. domain) can be set via [configuration](https://github.com/teamhanko/hanko/blob/main/backend/docs/Config.md#hanko-backend-config) option `session.cookie`.
	Value `<JWT>` is a [JSON Web Token](https://www.rfc-editor.org/rfc/rfc7519.html)

	*/
	SetCookie string

	/* Present only when enabled via [configuration](https://github.com/teamhanko/hanko/blob/main/backend/docs/Config.md#hanko-backend-config) option `session.enable_auth_token_header`
	for purposes of cross-domain communication between client and Hanko API.


	     Format: JWT
	*/
	XAuthToken string

	Payload *model.WebauthnLoginResponse
}

// IsSuccess returns true when this webauthn login final o k response has a 2xx status code
func (o *WebauthnLoginFinalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this webauthn login final o k response has a 3xx status code
func (o *WebauthnLoginFinalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this webauthn login final o k response has a 4xx status code
func (o *WebauthnLoginFinalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this webauthn login final o k response has a 5xx status code
func (o *WebauthnLoginFinalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this webauthn login final o k response a status code equal to that given
func (o *WebauthnLoginFinalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the webauthn login final o k response
func (o *WebauthnLoginFinalOK) Code() int {
	return 200
}

func (o *WebauthnLoginFinalOK) Error() string {
	return fmt.Sprintf("[POST /webauthn/login/finalize][%d] webauthnLoginFinalOK  %+v", 200, o.Payload)
}

func (o *WebauthnLoginFinalOK) String() string {
	return fmt.Sprintf("[POST /webauthn/login/finalize][%d] webauthnLoginFinalOK  %+v", 200, o.Payload)
}

func (o *WebauthnLoginFinalOK) GetPayload() *model.WebauthnLoginResponse {
	return o.Payload
}

func (o *WebauthnLoginFinalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Set-Cookie
	hdrSetCookie := response.GetHeader("Set-Cookie")

	if hdrSetCookie != "" {
		o.SetCookie = hdrSetCookie
	}

	// hydrates response header X-Auth-Token
	hdrXAuthToken := response.GetHeader("X-Auth-Token")

	if hdrXAuthToken != "" {
		o.XAuthToken = hdrXAuthToken
	}

	o.Payload = new(model.WebauthnLoginResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
