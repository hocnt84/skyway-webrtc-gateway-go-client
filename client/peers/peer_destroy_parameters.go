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

// NewPeerDestroyParams creates a new PeerDestroyParams object
// with the default values initialized.
func NewPeerDestroyParams() *PeerDestroyParams {
	var ()
	return &PeerDestroyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPeerDestroyParamsWithTimeout creates a new PeerDestroyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPeerDestroyParamsWithTimeout(timeout time.Duration) *PeerDestroyParams {
	var ()
	return &PeerDestroyParams{

		timeout: timeout,
	}
}

// NewPeerDestroyParamsWithContext creates a new PeerDestroyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPeerDestroyParamsWithContext(ctx context.Context) *PeerDestroyParams {
	var ()
	return &PeerDestroyParams{

		Context: ctx,
	}
}

// NewPeerDestroyParamsWithHTTPClient creates a new PeerDestroyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPeerDestroyParamsWithHTTPClient(client *http.Client) *PeerDestroyParams {
	var ()
	return &PeerDestroyParams{
		HTTPClient: client,
	}
}

/*PeerDestroyParams contains all the parameters to send to the API endpoint
for the peer destroy operation typically these are written to a http.Request
*/
type PeerDestroyParams struct {

	/*PeerID
	  他のピアがこのピアへ接続するときに使われるIDです。Peerオブジェクトの特定にも利用します

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

// WithTimeout adds the timeout to the peer destroy params
func (o *PeerDestroyParams) WithTimeout(timeout time.Duration) *PeerDestroyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the peer destroy params
func (o *PeerDestroyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the peer destroy params
func (o *PeerDestroyParams) WithContext(ctx context.Context) *PeerDestroyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the peer destroy params
func (o *PeerDestroyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the peer destroy params
func (o *PeerDestroyParams) WithHTTPClient(client *http.Client) *PeerDestroyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the peer destroy params
func (o *PeerDestroyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPeerID adds the peerID to the peer destroy params
func (o *PeerDestroyParams) WithPeerID(peerID string) *PeerDestroyParams {
	o.SetPeerID(peerID)
	return o
}

// SetPeerID adds the peerId to the peer destroy params
func (o *PeerDestroyParams) SetPeerID(peerID string) {
	o.PeerID = peerID
}

// WithToken adds the token to the peer destroy params
func (o *PeerDestroyParams) WithToken(token string) *PeerDestroyParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the peer destroy params
func (o *PeerDestroyParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PeerDestroyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
