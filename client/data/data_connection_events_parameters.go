// Code generated by go-swagger; DO NOT EDIT.

package data

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDataConnectionEventsParams creates a new DataConnectionEventsParams object
// with the default values initialized.
func NewDataConnectionEventsParams() *DataConnectionEventsParams {
	var ()
	return &DataConnectionEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDataConnectionEventsParamsWithTimeout creates a new DataConnectionEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDataConnectionEventsParamsWithTimeout(timeout time.Duration) *DataConnectionEventsParams {
	var ()
	return &DataConnectionEventsParams{

		timeout: timeout,
	}
}

// NewDataConnectionEventsParamsWithContext creates a new DataConnectionEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDataConnectionEventsParamsWithContext(ctx context.Context) *DataConnectionEventsParams {
	var ()
	return &DataConnectionEventsParams{

		Context: ctx,
	}
}

// NewDataConnectionEventsParamsWithHTTPClient creates a new DataConnectionEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDataConnectionEventsParamsWithHTTPClient(client *http.Client) *DataConnectionEventsParams {
	var ()
	return &DataConnectionEventsParams{
		HTTPClient: client,
	}
}

/*DataConnectionEventsParams contains all the parameters to send to the API endpoint
for the data connection events operation typically these are written to a http.Request
*/
type DataConnectionEventsParams struct {

	/*DataConnectionID
	  DataConnectionを特定するためのidを指定します

	*/
	DataConnectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the data connection events params
func (o *DataConnectionEventsParams) WithTimeout(timeout time.Duration) *DataConnectionEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the data connection events params
func (o *DataConnectionEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the data connection events params
func (o *DataConnectionEventsParams) WithContext(ctx context.Context) *DataConnectionEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the data connection events params
func (o *DataConnectionEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the data connection events params
func (o *DataConnectionEventsParams) WithHTTPClient(client *http.Client) *DataConnectionEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the data connection events params
func (o *DataConnectionEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDataConnectionID adds the dataConnectionID to the data connection events params
func (o *DataConnectionEventsParams) WithDataConnectionID(dataConnectionID string) *DataConnectionEventsParams {
	o.SetDataConnectionID(dataConnectionID)
	return o
}

// SetDataConnectionID adds the dataConnectionId to the data connection events params
func (o *DataConnectionEventsParams) SetDataConnectionID(dataConnectionID string) {
	o.DataConnectionID = dataConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *DataConnectionEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param data_connection_id
	if err := r.SetPathParam("data_connection_id", o.DataConnectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
