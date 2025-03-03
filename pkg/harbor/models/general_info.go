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

// GeneralInfo general info
//
// swagger:model GeneralInfo
type GeneralInfo struct {

	// The auth mode of current Harbor instance.
	AuthMode *string `json:"auth_mode,omitempty" js:"authMode"`

	// The setting of auth proxy this is only available when Harbor relies on authproxy for authentication.
	AuthproxySettings *AuthproxySetting `json:"authproxy_settings,omitempty" js:"authproxySettings"`

	// The external URL of Harbor, with protocol.
	ExternalURL *string `json:"external_url,omitempty" js:"externalURL"`

	// The build version of Harbor.
	HarborVersion *string `json:"harbor_version,omitempty" js:"harborVersion"`

	// Indicate whether there is a ca root cert file ready for download in the file system.
	HasCaRoot *bool `json:"has_ca_root,omitempty" js:"hasCaRoot"`

	// The flag to indicate whether notification mechanism is enabled on Harbor instance.
	NotificationEnable *bool `json:"notification_enable,omitempty" js:"notificationEnable"`

	// Indicate who can create projects, it could be 'adminonly' or 'everyone'.
	ProjectCreationRestriction *string `json:"project_creation_restriction,omitempty" js:"projectCreationRestriction"`

	// The flag to indicate whether Harbor is in readonly mode.
	ReadOnly *bool `json:"read_only,omitempty" js:"readOnly"`

	// The storage provider's name of Harbor registry
	RegistryStorageProviderName *string `json:"registry_storage_provider_name,omitempty" js:"registryStorageProviderName"`

	// The url of registry against which the docker command should be issued.
	RegistryURL *string `json:"registry_url,omitempty" js:"registryURL"`

	// Indicate whether the Harbor instance enable user to register himself.
	SelfRegistration *bool `json:"self_registration,omitempty" js:"selfRegistration"`

	// If the Harbor instance is deployed with nested chartmuseum.
	WithChartmuseum *bool `json:"with_chartmuseum,omitempty" js:"withChartmuseum"`

	// If the Harbor instance is deployed with nested notary.
	WithNotary *bool `json:"with_notary,omitempty" js:"withNotary"`
}

// Validate validates this general info
func (m *GeneralInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthproxySettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneralInfo) validateAuthproxySettings(formats strfmt.Registry) error {
	if swag.IsZero(m.AuthproxySettings) { // not required
		return nil
	}

	if m.AuthproxySettings != nil {
		if err := m.AuthproxySettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authproxy_settings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this general info based on the context it is used
func (m *GeneralInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthproxySettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneralInfo) contextValidateAuthproxySettings(ctx context.Context, formats strfmt.Registry) error {

	if m.AuthproxySettings != nil {
		if err := m.AuthproxySettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authproxy_settings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeneralInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneralInfo) UnmarshalBinary(b []byte) error {
	var res GeneralInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
