// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SetRoleAssignmentsCommand set role assignments command
//
// swagger:model SetRoleAssignmentsCommand
type SetRoleAssignmentsCommand struct {

	// service accounts
	ServiceAccounts []int64 `json:"service_accounts"`

	// teams
	Teams []int64 `json:"teams"`

	// users
	Users []int64 `json:"users"`
}

// Validate validates this set role assignments command
func (m *SetRoleAssignmentsCommand) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this set role assignments command based on context it is used
func (m *SetRoleAssignmentsCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SetRoleAssignmentsCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetRoleAssignmentsCommand) UnmarshalBinary(b []byte) error {
	var res SetRoleAssignmentsCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}