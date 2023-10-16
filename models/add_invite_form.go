// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddInviteForm add invite form
//
// swagger:model AddInviteForm
type AddInviteForm struct {

	// login or email
	LoginOrEmail string `json:"loginOrEmail,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// role
	// Enum: [Viewer Editor Admin]
	Role string `json:"role,omitempty"`

	// send email
	SendEmail bool `json:"sendEmail,omitempty"`
}

// Validate validates this add invite form
func (m *AddInviteForm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addInviteFormTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Viewer","Editor","Admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addInviteFormTypeRolePropEnum = append(addInviteFormTypeRolePropEnum, v)
	}
}

const (

	// AddInviteFormRoleViewer captures enum value "Viewer"
	AddInviteFormRoleViewer string = "Viewer"

	// AddInviteFormRoleEditor captures enum value "Editor"
	AddInviteFormRoleEditor string = "Editor"

	// AddInviteFormRoleAdmin captures enum value "Admin"
	AddInviteFormRoleAdmin string = "Admin"
)

// prop value enum
func (m *AddInviteForm) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, addInviteFormTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AddInviteForm) validateRole(formats strfmt.Registry) error {
	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this add invite form based on context it is used
func (m *AddInviteForm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddInviteForm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddInviteForm) UnmarshalBinary(b []byte) error {
	var res AddInviteForm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
