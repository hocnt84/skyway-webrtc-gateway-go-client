// Code generated by go-swagger; DO NOT EDIT.

package media

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

// NewMediaConnectionEventParams creates a new MediaConnectionEventParams object
// with the default values initialized.
func NewMediaConnectionEventParams() *MediaConnectionEventParams {
	var ()
	return &MediaConnectionEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMediaConnectionEventParamsWithTimeout creates a new MediaConnectionEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMediaConnectionEventParamsWithTimeout(timeout time.Duration) *MediaConnectionEventParams {
	var ()
	return &MediaConnectionEventParams{

		timeout: timeout,
	}
}

// NewMediaConnectionEventParamsWithContext creates a new MediaConnectionEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewMediaConnectionEventParamsWithContext(ctx context.Context) *MediaConnectionEventParams {
	var ()
	return &MediaConnectionEventParams{

		Context: ctx,
	}
}

// NewMediaConnectionEventParamsWithHTTPClient creates a new MediaConnectionEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMediaConnectionEventParamsWithHTTPClient(client *http.Client) *MediaConnectionEventParams {
	var ()
	return &MediaConnectionEventParams{
		HTTPClient: client,
	}
}

/*MediaConnectionEventParams contains all the parameters to send to the API endpoint
for the media connection event operation typically these are written to a http.Request
*/
type MediaConnectionEventParams struct {

	/*MediaConnectionID
	  MediaConnectionを特定するためのidを指定します

	*/
	MediaConnectionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the media connection event params
func (o *MediaConnectionEventParams) WithTimeout(timeout time.Duration) *MediaConnectionEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the media connection event params
func (o *MediaConnectionEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the media connection event params
func (o *MediaConnectionEventParams) WithContext(ctx context.Context) *MediaConnectionEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the media connection event params
func (o *MediaConnectionEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the media connection event params
func (o *MediaConnectionEventParams) WithHTTPClient(client *http.Client) *MediaConnectionEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the media connection event params
func (o *MediaConnectionEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMediaConnectionID adds the mediaConnectionID to the media connection event params
func (o *MediaConnectionEventParams) WithMediaConnectionID(mediaConnectionID string) *MediaConnectionEventParams {
	o.SetMediaConnectionID(mediaConnectionID)
	return o
}

// SetMediaConnectionID adds the mediaConnectionId to the media connection event params
func (o *MediaConnectionEventParams) SetMediaConnectionID(mediaConnectionID string) {
	o.MediaConnectionID = mediaConnectionID
}

// WriteToRequest writes these params to a swagger request
func (o *MediaConnectionEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param media_connection_id
	if err := r.SetPathParam("media_connection_id", o.MediaConnectionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
