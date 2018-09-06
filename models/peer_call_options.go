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

// PeerCallOptions peer call options
// swagger:model PeerCallOptions
type PeerCallOptions struct {

	// constraints
	// Required: true
	Constraints *PeerCallConstraints `json:"constraints"`

	// 自身のPeerのidを指定します
	// Required: true
	PeerID *string `json:"peer_id"`

	// redirect params
	RedirectParams *MediaRedirectOptions `json:"redirect_params,omitempty"`

	// 接続対象のpeer_idを指定します
	// Required: true
	TargetID *string `json:"target_id"`

	// Peerオブジェクトを利用するために必要なトークンです。他のユーザのリソースに対する誤操作防止のために指定します
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this peer call options
func (m *PeerCallOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirectParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PeerCallOptions) validateConstraints(formats strfmt.Registry) error {

	if err := validate.Required("constraints", "body", m.Constraints); err != nil {
		return err
	}

	if m.Constraints != nil {
		if err := m.Constraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("constraints")
			}
			return err
		}
	}

	return nil
}

func (m *PeerCallOptions) validatePeerID(formats strfmt.Registry) error {

	if err := validate.Required("peer_id", "body", m.PeerID); err != nil {
		return err
	}

	return nil
}

func (m *PeerCallOptions) validateRedirectParams(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirectParams) { // not required
		return nil
	}

	if m.RedirectParams != nil {
		if err := m.RedirectParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redirect_params")
			}
			return err
		}
	}

	return nil
}

func (m *PeerCallOptions) validateTargetID(formats strfmt.Registry) error {

	if err := validate.Required("target_id", "body", m.TargetID); err != nil {
		return err
	}

	return nil
}

func (m *PeerCallOptions) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PeerCallOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PeerCallOptions) UnmarshalBinary(b []byte) error {
	var res PeerCallOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
