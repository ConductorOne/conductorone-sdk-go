// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The PolicySteps message.
type PolicySteps struct {
	// An array of policy steps indicating the processing flow of a policy. These steps are oneOfs, and only one property may be set for each array index at a time.
	Steps []PolicyStep `json:"steps,omitempty"`
}

func (o *PolicySteps) GetSteps() []PolicyStep {
	if o == nil {
		return nil
	}
	return o.Steps
}
