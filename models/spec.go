// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Spec Spec defines model for Spec.
//
// swagger:model Spec
type Spec struct {

	// UID for the home dashboard
	HomeDashboardUID string `json:"homeDashboardUID,omitempty"`

	// Selected language (beta)
	Language string `json:"language,omitempty"`

	// query history
	QueryHistory *QueryHistoryPreference `json:"queryHistory,omitempty"`

	// Theme light, dark, empty is default
	Theme string `json:"theme,omitempty"`

	// The timezone selection
	// TODO: this should use the timezone defined in common
	Timezone string `json:"timezone,omitempty"`

	// WeekStart day of the week (sunday, monday, etc)
	WeekStart string `json:"weekStart,omitempty"`
}

// Validate validates this spec
func (m *Spec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQueryHistory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Spec) validateQueryHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.QueryHistory) { // not required
		return nil
	}

	if m.QueryHistory != nil {
		if err := m.QueryHistory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryHistory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryHistory")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this spec based on the context it is used
func (m *Spec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateQueryHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Spec) contextValidateQueryHistory(ctx context.Context, formats strfmt.Registry) error {

	if m.QueryHistory != nil {

		if swag.IsZero(m.QueryHistory) { // not required
			return nil
		}

		if err := m.QueryHistory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queryHistory")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queryHistory")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Spec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Spec) UnmarshalBinary(b []byte) error {
	var res Spec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
