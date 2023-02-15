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

// AddressData AddressData
//
// swagger:model AddressData
type AddressData struct {

	//
	Empty AdditionalProperties `json:",omitempty"`

	// city
	// Required: true
	City *string `json:"city"`

	// country
	// Required: true
	Country *NullableString `json:"country"`

	// postal code
	// Required: true
	PostalCode *NullableString `json:"postal_code"`

	// region
	// Required: true
	Region *NullableString `json:"region"`

	// street
	// Required: true
	Street *string `json:"street"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *AddressData) UnmarshalJSON(data []byte) error {
	var props struct {

		//
		Empty AdditionalProperties `json:",omitempty"`

		// city
		// Required: true
		City *string `json:"city"`

		// country
		// Required: true
		Country *NullableString `json:"country"`

		// postal code
		// Required: true
		PostalCode *NullableString `json:"postal_code"`

		// region
		// Required: true
		Region *NullableString `json:"region"`

		// street
		// Required: true
		Street *string `json:"street"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Empty = props.Empty
	m.City = props.City
	m.Country = props.Country
	m.PostalCode = props.PostalCode
	m.Region = props.Region
	m.Street = props.Street
	return nil
}

// Validate validates this address data
func (m *AddressData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostalCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressData) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *AddressData) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	if m.Country != nil {
		if err := m.Country.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("country")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("country")
			}
			return err
		}
	}

	return nil
}

func (m *AddressData) validatePostalCode(formats strfmt.Registry) error {

	if err := validate.Required("postal_code", "body", m.PostalCode); err != nil {
		return err
	}

	if m.PostalCode != nil {
		if err := m.PostalCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postal_code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postal_code")
			}
			return err
		}
	}

	return nil
}

func (m *AddressData) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	if m.Region != nil {
		if err := m.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (m *AddressData) validateStreet(formats strfmt.Registry) error {

	if err := validate.Required("street", "body", m.Street); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this address data based on the context it is used
func (m *AddressData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEmpty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostalCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddressData) contextValidateEmpty(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Empty.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("")
		}
		return err
	}

	return nil
}

func (m *AddressData) contextValidateCountry(ctx context.Context, formats strfmt.Registry) error {

	if m.Country != nil {
		if err := m.Country.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("country")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("country")
			}
			return err
		}
	}

	return nil
}

func (m *AddressData) contextValidatePostalCode(ctx context.Context, formats strfmt.Registry) error {

	if m.PostalCode != nil {
		if err := m.PostalCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postal_code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postal_code")
			}
			return err
		}
	}

	return nil
}

func (m *AddressData) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if m.Region != nil {
		if err := m.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddressData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddressData) UnmarshalBinary(b []byte) error {
	var res AddressData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
