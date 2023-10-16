// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateOrgOKBody create org o k body
//
// swagger:model createOrgOKBody
type CreateOrgOKBody struct {

	// Message Message of the created org.
	// Example: Data source added
	// Required: true
	Message *string `json:"message"`

	// ID Identifier of the created org.
	// Example: 65
	// Required: true
	OrgID *int64 `json:"orgId"`
}

// Validate validates this create org o k body
func (m *CreateOrgOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrgID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateOrgOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *CreateOrgOKBody) validateOrgID(formats strfmt.Registry) error {

	if err := validate.Required("orgId", "body", m.OrgID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create org o k body based on context it is used
func (m *CreateOrgOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateOrgOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateOrgOKBody) UnmarshalBinary(b []byte) error {
	var res CreateOrgOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
