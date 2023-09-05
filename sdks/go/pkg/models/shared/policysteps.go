// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PolicySteps - The PolicySteps message.
type PolicySteps struct {
	// An array of policy steps indicating the processing flow of a policy. These steps are oneOfs, and only one property may be set for each array index at a time.
	Steps []PolicyStep `json:"steps,omitempty"`
}

// PolicyStepsInput - The PolicySteps message.
type PolicyStepsInput struct {
	// An array of policy steps indicating the processing flow of a policy. These steps are oneOfs, and only one property may be set for each array index at a time.
	Steps []PolicyStepInput `json:"steps,omitempty"`
}