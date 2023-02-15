// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LinkTokenCreateResp LinkTokenCreateResp
//
// swagger:model LinkTokenCreateResp
type LinkTokenCreateResp struct {

	// link token
	// Required: true
	LinkToken *string `json:"link_token"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *LinkTokenCreateResp) UnmarshalJSON(data []byte) error {
	var props struct {

		// link token
		// Required: true
		LinkToken *string `json:"link_token"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.LinkToken = props.LinkToken
	return nil
}

// Validate validates this link token create resp
func (m *LinkTokenCreateResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinkToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LinkTokenCreateResp) validateLinkToken(formats strfmt.Registry) error {

	if err := validate.Required("link_token", "body", m.LinkToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this link token create resp based on context it is used
func (m *LinkTokenCreateResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinkTokenCreateResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinkTokenCreateResp) UnmarshalBinary(b []byte) error {
	var res LinkTokenCreateResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
