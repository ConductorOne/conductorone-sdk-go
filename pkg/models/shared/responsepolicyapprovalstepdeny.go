// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ResponsePolicyApprovalStepDeny message.
type ResponsePolicyApprovalStepDeny struct {
	// optional comment
	Comment *string `json:"comment,omitempty"`
}

func (o *ResponsePolicyApprovalStepDeny) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}