// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The UpdatePolicyRequest message contains the policy object to update and a field mask to indicate which fields to update. It uses URL value for input.
type UpdatePolicyRequest struct {
	// A policy describes the behavior of the ConductorOne system when processing a task. You can describe the type, approvers, fallback behavior, and escalation processes.
	Policy     *PolicyInput `json:"policy,omitempty"`
	UpdateMask *string      `json:"updateMask,omitempty"`
}

func (o *UpdatePolicyRequest) GetPolicy() *PolicyInput {
	if o == nil {
		return nil
	}
	return o.Policy
}

func (o *UpdatePolicyRequest) GetUpdateMask() *string {
	if o == nil {
		return nil
	}
	return o.UpdateMask
}
