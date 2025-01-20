// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The WorkflowStep message.
//
// This message contains a oneof named kind. Only a single field of the following list may be set at a time:
//   - createAccessReview
type WorkflowStep struct {
	// The CreateAccessReview message.
	CreateAccessReview *CreateAccessReview `json:"createAccessReview,omitempty"`
}

func (o *WorkflowStep) GetCreateAccessReview() *CreateAccessReview {
	if o == nil {
		return nil
	}
	return o.CreateAccessReview
}
