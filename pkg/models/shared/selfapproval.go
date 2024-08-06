// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// SelfApproval - The self approval object describes the configuration of a policy step that needs to be approved by the target of the request.
type SelfApproval struct {
	// The array of users determined to be themselves during approval. This should only ever be one person, but is saved because it may change if the owner of an app user changes while the ticket is open.
	AssignedUserIds []string `json:"assignedUserIds,omitempty"`
	// Configuration to allow a fallback if the identity user of the target app user cannot be determined.
	Fallback *bool `json:"fallback,omitempty"`
	// Configuration to specific which users to fallback to if fallback is enabled and the identity user of the target app user cannot be determined.
	FallbackUserIds []string `json:"fallbackUserIds,omitempty"`
}

func (o *SelfApproval) GetAssignedUserIds() []string {
	if o == nil {
		return nil
	}
	return o.AssignedUserIds
}

func (o *SelfApproval) GetFallback() *bool {
	if o == nil {
		return nil
	}
	return o.Fallback
}

func (o *SelfApproval) GetFallbackUserIds() []string {
	if o == nil {
		return nil
	}
	return o.FallbackUserIds
}
