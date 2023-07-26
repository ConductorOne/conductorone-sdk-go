// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateRoleRequestInput -  The UpdateRoleRequest message contains the role to update and the update mask.
type UpdateRoleRequestInput struct {
	//  Role is a role that can be assigned to a user in ConductorOne.
	//
	Role       *RoleInput `json:"role,omitempty"`
	UpdateMask *string    `json:"updateMask,omitempty"`
}

func (o *UpdateRoleRequestInput) GetRole() *RoleInput {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *UpdateRoleRequestInput) GetUpdateMask() *string {
	if o == nil {
		return nil
	}
	return o.UpdateMask
}
