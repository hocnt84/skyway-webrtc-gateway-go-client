// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataConnectionRedirectOptions data connection redirect options
// swagger:model DataConnectionRedirectOptions
type DataConnectionRedirectOptions struct {

	// Dataの転送先IPv4アドレスを指定します。ip_v4またはip_v6アドレスのいずれかの情報は必須です。両方が指定された場合はip_v4が優先されます
	IPV4 string `json:"ip_v4,omitempty"`

	// Dataの転送先IPv6アドレスを指定します。ip_v4またはip_v6アドレスのいずれかの情報は必須です。両方が指定された場合はip_v4が優先されます
	IPV6 string `json:"ip_v6,omitempty"`

	// port
	// Required: true
	Port *uint16 `json:"port"`
}

// Validate validates this data connection redirect options
func (m *DataConnectionRedirectOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataConnectionRedirectOptions) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataConnectionRedirectOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataConnectionRedirectOptions) UnmarshalBinary(b []byte) error {
	var res DataConnectionRedirectOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}