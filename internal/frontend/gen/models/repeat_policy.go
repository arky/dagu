// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RepeatPolicy Configuration for step retry behavior
//
// swagger:model RepeatPolicy
type RepeatPolicy struct {

	// Time in seconds to wait between retry attempts
	Interval int64 `json:"Interval,omitempty"`

	// Whether the step should be retried on failure
	Repeat bool `json:"Repeat,omitempty"`
}

// Validate validates this repeat policy
func (m *RepeatPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this repeat policy based on context it is used
func (m *RepeatPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RepeatPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepeatPolicy) UnmarshalBinary(b []byte) error {
	var res RepeatPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
