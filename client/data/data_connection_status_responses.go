// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/atotto/skyway-webrtc-gateway-go-client/models"
)

// DataConnectionStatusReader is a Reader for the DataConnectionStatus structure.
type DataConnectionStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DataConnectionStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDataConnectionStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDataConnectionStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDataConnectionStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDataConnectionStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 405:
		result := NewDataConnectionStatusMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewDataConnectionStatusNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 408:
		result := NewDataConnectionStatusRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDataConnectionStatusOK creates a DataConnectionStatusOK with default headers values
func NewDataConnectionStatusOK() *DataConnectionStatusOK {
	return &DataConnectionStatusOK{}
}

/*DataConnectionStatusOK handles this case with default header values.

successful operation
*/
type DataConnectionStatusOK struct {
	Payload *models.DataConnectionStatusMessage
}

func (o *DataConnectionStatusOK) Error() string {
	return fmt.Sprintf("[GET /data/connections/{data_connection_id}/status][%d] dataConnectionStatusOK  %+v", 200, o.Payload)
}

func (o *DataConnectionStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataConnectionStatusMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataConnectionStatusBadRequest creates a DataConnectionStatusBadRequest with default headers values
func NewDataConnectionStatusBadRequest() *DataConnectionStatusBadRequest {
	return &DataConnectionStatusBadRequest{}
}

/*DataConnectionStatusBadRequest handles this case with default header values.

Invalid input
*/
type DataConnectionStatusBadRequest struct {
	Payload *DataConnectionStatusBadRequestBody
}

func (o *DataConnectionStatusBadRequest) Error() string {
	return fmt.Sprintf("[GET /data/connections/{data_connection_id}/status][%d] dataConnectionStatusBadRequest  %+v", 400, o.Payload)
}

func (o *DataConnectionStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DataConnectionStatusBadRequestBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDataConnectionStatusForbidden creates a DataConnectionStatusForbidden with default headers values
func NewDataConnectionStatusForbidden() *DataConnectionStatusForbidden {
	return &DataConnectionStatusForbidden{}
}

/*DataConnectionStatusForbidden handles this case with default header values.

Forbidden(不正な操作を行った場合)
*/
type DataConnectionStatusForbidden struct {
}

func (o *DataConnectionStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /data/connections/{data_connection_id}/status][%d] dataConnectionStatusForbidden ", 403)
}

func (o *DataConnectionStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataConnectionStatusNotFound creates a DataConnectionStatusNotFound with default headers values
func NewDataConnectionStatusNotFound() *DataConnectionStatusNotFound {
	return &DataConnectionStatusNotFound{}
}

/*DataConnectionStatusNotFound handles this case with default header values.

Not Found(data_connection_idが間違っている場合)
*/
type DataConnectionStatusNotFound struct {
}

func (o *DataConnectionStatusNotFound) Error() string {
	return fmt.Sprintf("[GET /data/connections/{data_connection_id}/status][%d] dataConnectionStatusNotFound ", 404)
}

func (o *DataConnectionStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataConnectionStatusMethodNotAllowed creates a DataConnectionStatusMethodNotAllowed with default headers values
func NewDataConnectionStatusMethodNotAllowed() *DataConnectionStatusMethodNotAllowed {
	return &DataConnectionStatusMethodNotAllowed{}
}

/*DataConnectionStatusMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type DataConnectionStatusMethodNotAllowed struct {
}

func (o *DataConnectionStatusMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /data/connections/{data_connection_id}/status][%d] dataConnectionStatusMethodNotAllowed ", 405)
}

func (o *DataConnectionStatusMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataConnectionStatusNotAcceptable creates a DataConnectionStatusNotAcceptable with default headers values
func NewDataConnectionStatusNotAcceptable() *DataConnectionStatusNotAcceptable {
	return &DataConnectionStatusNotAcceptable{}
}

/*DataConnectionStatusNotAcceptable handles this case with default header values.

Not Acceptable
*/
type DataConnectionStatusNotAcceptable struct {
}

func (o *DataConnectionStatusNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /data/connections/{data_connection_id}/status][%d] dataConnectionStatusNotAcceptable ", 406)
}

func (o *DataConnectionStatusNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDataConnectionStatusRequestTimeout creates a DataConnectionStatusRequestTimeout with default headers values
func NewDataConnectionStatusRequestTimeout() *DataConnectionStatusRequestTimeout {
	return &DataConnectionStatusRequestTimeout{}
}

/*DataConnectionStatusRequestTimeout handles this case with default header values.

Request Timeout
*/
type DataConnectionStatusRequestTimeout struct {
}

func (o *DataConnectionStatusRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /data/connections/{data_connection_id}/status][%d] dataConnectionStatusRequestTimeout ", 408)
}

func (o *DataConnectionStatusRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*DataConnectionStatusBadRequestBody data connection status bad request body
swagger:model DataConnectionStatusBadRequestBody
*/
type DataConnectionStatusBadRequestBody struct {

	// command type
	// Required: true
	CommandType *string `json:"command_type"`

	// params
	// Required: true
	Params *DataConnectionStatusBadRequestBodyParams `json:"params"`
}

// Validate validates this data connection status bad request body
func (o *DataConnectionStatusBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommandType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DataConnectionStatusBadRequestBody) validateCommandType(formats strfmt.Registry) error {

	if err := validate.Required("dataConnectionStatusBadRequest"+"."+"command_type", "body", o.CommandType); err != nil {
		return err
	}

	return nil
}

func (o *DataConnectionStatusBadRequestBody) validateParams(formats strfmt.Registry) error {

	if err := validate.Required("dataConnectionStatusBadRequest"+"."+"params", "body", o.Params); err != nil {
		return err
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dataConnectionStatusBadRequest" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DataConnectionStatusBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DataConnectionStatusBadRequestBody) UnmarshalBinary(b []byte) error {
	var res DataConnectionStatusBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DataConnectionStatusBadRequestBodyParams data connection status bad request body params
swagger:model DataConnectionStatusBadRequestBodyParams
*/
type DataConnectionStatusBadRequestBodyParams struct {

	// errors
	// Required: true
	Errors []*DataConnectionStatusBadRequestBodyParamsErrorsItems0 `json:"errors"`
}

// Validate validates this data connection status bad request body params
func (o *DataConnectionStatusBadRequestBodyParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DataConnectionStatusBadRequestBodyParams) validateErrors(formats strfmt.Registry) error {

	if err := validate.Required("dataConnectionStatusBadRequest"+"."+"params"+"."+"errors", "body", o.Errors); err != nil {
		return err
	}

	for i := 0; i < len(o.Errors); i++ {
		if swag.IsZero(o.Errors[i]) { // not required
			continue
		}

		if o.Errors[i] != nil {
			if err := o.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataConnectionStatusBadRequest" + "." + "params" + "." + "errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DataConnectionStatusBadRequestBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DataConnectionStatusBadRequestBodyParams) UnmarshalBinary(b []byte) error {
	var res DataConnectionStatusBadRequestBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DataConnectionStatusBadRequestBodyParamsErrorsItems0 data connection status bad request body params errors items0
swagger:model DataConnectionStatusBadRequestBodyParamsErrorsItems0
*/
type DataConnectionStatusBadRequestBodyParamsErrorsItems0 struct {

	// field
	Field string `json:"field,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this data connection status bad request body params errors items0
func (o *DataConnectionStatusBadRequestBodyParamsErrorsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DataConnectionStatusBadRequestBodyParamsErrorsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DataConnectionStatusBadRequestBodyParamsErrorsItems0) UnmarshalBinary(b []byte) error {
	var res DataConnectionStatusBadRequestBodyParamsErrorsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
