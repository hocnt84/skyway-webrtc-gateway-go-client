// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DataConnectionFeedOptions data connection feed options
// swagger:model DataConnectionFeedOptions
type DataConnectionFeedOptions struct {

	// params
	Params *DataConnectionFeedOptionsParams `json:"params,omitempty"`
}

// Validate validates this data connection feed options
func (m *DataConnectionFeedOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataConnectionFeedOptions) validateParams(formats strfmt.Registry) error {

	if swag.IsZero(m.Params) { // not required
		return nil
	}

	if m.Params != nil {
		if err := m.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataConnectionFeedOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataConnectionFeedOptions) UnmarshalBinary(b []byte) error {
	var res DataConnectionFeedOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DataConnectionFeedOptionsParams data connection feed options params
// swagger:model DataConnectionFeedOptionsParams
type DataConnectionFeedOptionsParams struct {

	// Dataを特定するためのIDです
	DataID string `json:"data_id,omitempty"`
}

// Validate validates this data connection feed options params
func (m *DataConnectionFeedOptionsParams) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataConnectionFeedOptionsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataConnectionFeedOptionsParams) UnmarshalBinary(b []byte) error {
	var res DataConnectionFeedOptionsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}