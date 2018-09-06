// Code generated by go-swagger; DO NOT EDIT.

package peers

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

// NewPeerEventParams creates a new PeerEventParams object
// with the default values initialized.
func NewPeerEventParams() *PeerEventParams {
	var ()
	return &PeerEventParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPeerEventParamsWithTimeout creates a new PeerEventParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPeerEventParamsWithTimeout(timeout time.Duration) *PeerEventParams {
	var ()
	return &PeerEventParams{

		timeout: timeout,
	}
}

// NewPeerEventParamsWithContext creates a new PeerEventParams object
// with the default values initialized, and the ability to set a context for a request
func NewPeerEventParamsWithContext(ctx context.Context) *PeerEventParams {
	var ()
	return &PeerEventParams{

		Context: ctx,
	}
}

// NewPeerEventParamsWithHTTPClient creates a new PeerEventParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPeerEventParamsWithHTTPClient(client *http.Client) *PeerEventParams {
	var ()
	return &PeerEventParams{
		HTTPClient: client,
	}
}

/*PeerEventParams contains all the parameters to send to the API endpoint
for the peer event operation typically these are written to a http.Request
*/
type PeerEventParams struct {

	/*PeerID
	  接続対象のPeerのidを指定します

	*/
	PeerID string
	/*Token
	  Peerオブジェクトを利用するために必要なトークンです。他のユーザのリソースに対する誤操作防止のために指定します

	*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the peer event params
func (o *PeerEventParams) WithTimeout(timeout time.Duration) *PeerEventParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peer event params
func (o *PeerEventParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peer event params
func (o *PeerEventParams) WithContext(ctx context.Context) *PeerEventParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peer event params
func (o *PeerEventParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peer event params
func (o *PeerEventParams) WithHTTPClient(client *http.Client) *PeerEventParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peer event params
func (o *PeerEventParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPeerID adds the peerID to the peer event params
func (o *PeerEventParams) WithPeerID(peerID string) *PeerEventParams {
	o.SetPeerID(peerID)
	return o
}

// SetPeerID adds the peerId to the peer event params
func (o *PeerEventParams) SetPeerID(peerID string) {
	o.PeerID = peerID
}

// WithToken adds the token to the peer event params
func (o *PeerEventParams) WithToken(token string) *PeerEventParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the peer event params
func (o *PeerEventParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PeerEventParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param peer_id
	if err := r.SetPathParam("peer_id", o.PeerID); err != nil {
		return err
	}

	// query param token
	qrToken := o.Token
	qToken := qrToken
	if qToken != "" {
		if err := r.SetQueryParam("token", qToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}