// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskActionsServiceDenyRequest - The TaskActionsServiceDenyRequest object lets you deny a task.
type TaskActionsServiceDenyRequest struct {
	// The task expand mask is an array of strings that specifes the related objects the requester wishes to have returned when making a request where the expand mask is part of the input. Use '*' to view all possible responses.
	TaskExpandMask *TaskExpandMask `json:"expandMask,omitempty"`
	// The comment attached to the request.
	Comment *string `json:"comment,omitempty"`
	// The ID of the currently policy step. This is the step you want to deny.
	PolicyStepID *string `json:"policyStepId,omitempty"`
}

func (o *TaskActionsServiceDenyRequest) GetTaskExpandMask() *TaskExpandMask {
	if o == nil {
		return nil
	}
	return o.TaskExpandMask
}

func (o *TaskActionsServiceDenyRequest) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}

func (o *TaskActionsServiceDenyRequest) GetPolicyStepID() *string {
	if o == nil {
		return nil
	}
	return o.PolicyStepID
}
