// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateVolume update volume
//
// swagger:model UpdateVolume
type UpdateVolume struct {

	// Indicates if the volume is boot capable
	Bootable *bool `json:"bootable,omitempty"`

	// Name
	Name *string `json:"name,omitempty"`

	// Indicates if the volume is shareable between VMs
	Shareable *bool `json:"shareable,omitempty"`

	// New Volume size
	Size float64 `json:"size,omitempty"`
}

// Validate validates this update volume
func (m *UpdateVolume) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update volume based on context it is used
func (m *UpdateVolume) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateVolume) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateVolume) UnmarshalBinary(b []byte) error {
	var res UpdateVolume
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
