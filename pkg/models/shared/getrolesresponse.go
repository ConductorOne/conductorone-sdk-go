// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The GetRolesResponse message contains the retrieved role.
type GetRolesResponse struct {
	// Role is a role that can be assigned to a user in ConductorOne.
	Role *Role `json:"role,omitempty"`
}

func (o *GetRolesResponse) GetRole() *Role {
	if o == nil {
		return nil
	}
	return o.Role
}
