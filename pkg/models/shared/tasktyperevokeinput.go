// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskTypeRevokeInput - The TaskTypeRevoke message indicates that a task is a revoke task and all related details.
type TaskTypeRevokeInput struct {
	// The TaskRevokeSource message indicates the source of the revoke task is one of expired, nonUsage, request, or review.
	//
	// This message contains a oneof named origin. Only a single field of the following list may be set at a time:
	//   - review
	//   - request
	//   - expired
	//   - nonUsage
	//
	TaskRevokeSource *TaskRevokeSource `json:"source,omitempty"`
}

func (o *TaskTypeRevokeInput) GetTaskRevokeSource() *TaskRevokeSource {
	if o == nil {
		return nil
	}
	return o.TaskRevokeSource
}
