// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdatePolicyRequest - The UpdatePolicyRequest message.
type UpdatePolicyRequest struct {
	// The Policy message.
	Policy     *Policy `json:"policy,omitempty"`
	UpdateMask *string `json:"updateMask,omitempty"`
}
