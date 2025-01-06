// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The TaskActionsServiceProcessNowRequest object lets you trigger processing of a task immediately.
type TaskActionsServiceProcessNowRequest struct {
	// The task expand mask is an array of strings that specifes the related objects the requester wishes to have returned when making a request where the expand mask is part of the input. Use '*' to view all possible responses.
	TaskExpandMask *TaskExpandMask `json:"expandMask,omitempty"`
}

func (o *TaskActionsServiceProcessNowRequest) GetTaskExpandMask() *TaskExpandMask {
	if o == nil {
		return nil
	}
	return o.TaskExpandMask
}
