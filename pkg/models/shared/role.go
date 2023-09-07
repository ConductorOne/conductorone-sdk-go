// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// RoleInput - Role is a role that can be assigned to a user in ConductorOne.
type RoleInput struct {
	// The display name of the role.
	DisplayName *string `json:"displayName,omitempty"`
	// The list of permissions this role has.
	Permissions []string `json:"permissions,omitempty"`
	// The list of serviceRoles that this role has.
	ServiceRoles []string `json:"serviceRoles,omitempty"`
}

func (o *RoleInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *RoleInput) GetPermissions() []string {
	if o == nil {
		return nil
	}
	return o.Permissions
}

func (o *RoleInput) GetServiceRoles() []string {
	if o == nil {
		return nil
	}
	return o.ServiceRoles
}

// Role is a role that can be assigned to a user in ConductorOne.
type Role struct {
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The display name of the role.
	DisplayName *string `json:"displayName,omitempty"`
	// The id of the role.
	ID *string `json:"id,omitempty"`
	// The internal name of the role.
	Name *string `json:"name,omitempty"`
	// The list of permissions this role has.
	Permissions []string `json:"permissions,omitempty"`
	// The list of serviceRoles that this role has.
	ServiceRoles []string `json:"serviceRoles,omitempty"`
	// The system builtin field. If this field is set, the role is not editable.
	SystemBuiltin *bool      `json:"systemBuiltin,omitempty"`
	UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
}

func (o *Role) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Role) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *Role) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Role) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Role) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Role) GetPermissions() []string {
	if o == nil {
		return nil
	}
	return o.Permissions
}

func (o *Role) GetServiceRoles() []string {
	if o == nil {
		return nil
	}
	return o.ServiceRoles
}

func (o *Role) GetSystemBuiltin() *bool {
	if o == nil {
		return nil
	}
	return o.SystemBuiltin
}

func (o *Role) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
