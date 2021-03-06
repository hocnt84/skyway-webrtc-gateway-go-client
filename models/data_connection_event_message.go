// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataConnectionEventMessage data connection event message
// swagger:model DataConnectionEventMessage
type DataConnectionEventMessage struct {

	// エラーの内容を示します
	ErrorMessage string `json:"error_message,omitempty"`

	// イベントの種別を示します
	// Required: true
	// Enum: [OPEN CLOSE ERROR]
	Event *string `json:"event"`
}

// Validate validates this data connection event message
func (m *DataConnectionEventMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dataConnectionEventMessageTypeEventPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OPEN","CLOSE","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataConnectionEventMessageTypeEventPropEnum = append(dataConnectionEventMessageTypeEventPropEnum, v)
	}
}

const (

	// DataConnectionEventMessageEventOPEN captures enum value "OPEN"
	DataConnectionEventMessageEventOPEN string = "OPEN"

	// DataConnectionEventMessageEventCLOSE captures enum value "CLOSE"
	DataConnectionEventMessageEventCLOSE string = "CLOSE"

	// DataConnectionEventMessageEventERROR captures enum value "ERROR"
	DataConnectionEventMessageEventERROR string = "ERROR"
)

// prop value enum
func (m *DataConnectionEventMessage) validateEventEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, dataConnectionEventMessageTypeEventPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DataConnectionEventMessage) validateEvent(formats strfmt.Registry) error {

	if err := validate.Required("event", "body", m.Event); err != nil {
		return err
	}

	// value enum
	if err := m.validateEventEnum("event", "body", *m.Event); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataConnectionEventMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataConnectionEventMessage) UnmarshalBinary(b []byte) error {
	var res DataConnectionEventMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
