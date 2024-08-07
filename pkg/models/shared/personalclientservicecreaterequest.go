// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The PersonalClientServiceCreateRequest message contains the fields for creating a new personal client.
type PersonalClientServiceCreateRequest struct {
	// A list of CIDRs to restrict this credential to.
	AllowSourceCidr []string `json:"allowSourceCidr,omitempty"`
	// The display name for the new personal client.
	DisplayName *string `json:"displayName,omitempty"`
	Expires     *string `json:"expires,omitempty"`
	// The list of roles to restrict the credential to.
	ScopedRoles []string `json:"scopedRoles,omitempty"`
}

func (o *PersonalClientServiceCreateRequest) GetAllowSourceCidr() []string {
	if o == nil {
		return nil
	}
	return o.AllowSourceCidr
}

func (o *PersonalClientServiceCreateRequest) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *PersonalClientServiceCreateRequest) GetExpires() *string {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *PersonalClientServiceCreateRequest) GetScopedRoles() []string {
	if o == nil {
		return nil
	}
	return o.ScopedRoles
}
