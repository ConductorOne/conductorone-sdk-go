// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ManagerApproval - The manager approval object provides configuration options for approval when the target of the approval is the manager of the user in the task.
type ManagerApproval struct {
	// Configuration to allow self approval if the target user is their own manager. This may occur if a service account has an identity user and manager specified as the same person.
	AllowSelfApproval *bool `json:"allowSelfApproval,omitempty"`
	// The array of users determined to be the manager during processing time.
	AssignedUserIds []string `json:"assignedUserIds,omitempty"`
	// Configuration to allow a fallback if no manager is found.
	Fallback *bool `json:"fallback,omitempty"`
	// Configuration to specific which users to fallback to if fallback is enabled and no manager is found.
	FallbackUserIds []string `json:"fallbackUserIds,omitempty"`
}

func (o *ManagerApproval) GetAllowSelfApproval() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSelfApproval
}

func (o *ManagerApproval) GetAssignedUserIds() []string {
	if o == nil {
		return nil
	}
	return o.AssignedUserIds
}

func (o *ManagerApproval) GetFallback() *bool {
	if o == nil {
		return nil
	}
	return o.Fallback
}

func (o *ManagerApproval) GetFallbackUserIds() []string {
	if o == nil {
		return nil
	}
	return o.FallbackUserIds
}

// ManagerApprovalInput - The manager approval object provides configuration options for approval when the target of the approval is the manager of the user in the task.
type ManagerApprovalInput struct {
}
