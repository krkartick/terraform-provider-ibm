// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TenantCreate tenant create
// swagger:model TenantCreate
type TenantCreate struct {

	// Billing account ID
	BillingAccountID string `json:"billingAccountID,omitempty"`

	// Entitlement ID
	EntitlementID string `json:"entitlementID,omitempty"`

	// Tenant SSH Keys
	SSHKeys []*SSHKey `json:"sshKeys"`

	// Tenant ID
	// Required: true
	TenantID *string `json:"tenantID"`
}

// Validate validates this tenant create
func (m *TenantCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSSHKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TenantCreate) validateSSHKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.SSHKeys) { // not required
		return nil
	}

	for i := 0; i < len(m.SSHKeys); i++ {
		if swag.IsZero(m.SSHKeys[i]) { // not required
			continue
		}

		if m.SSHKeys[i] != nil {
			if err := m.SSHKeys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sshKeys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TenantCreate) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantID", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TenantCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantCreate) UnmarshalBinary(b []byte) error {
	var res TenantCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}