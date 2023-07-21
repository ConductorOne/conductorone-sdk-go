// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskActionsServiceCommentRequest -  The TaskActionsServiceCommentRequest object lets you create a new comment on a task.
type TaskActionsServiceCommentRequest struct {
	//  The task expand mask is an array of strings that specifes the related objects the requester wishes to have returned when making a request where the expand mask is part of the input. Use '*' to view all possible responses.
	//
	TaskExpandMask *TaskExpandMask `json:"expandMask,omitempty"`
	//  The comment to be posted to the ticket
	//
	Comment *string `json:"comment,omitempty"`
}

func (o *TaskActionsServiceCommentRequest) GetTaskExpandMask() *TaskExpandMask {
	if o == nil {
		return nil
	}
	return o.TaskExpandMask
}

func (o *TaskActionsServiceCommentRequest) GetComment() *string {
	if o == nil {
		return nil
	}
	return o.Comment
}
