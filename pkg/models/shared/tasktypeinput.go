// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskTypeInput - Task Type provides configuration for the type of task: certify, grant, or revoke
//
// This message contains a oneof named task_type. Only a single field of the following list may be set at a time:
//   - grant
//   - revoke
//   - certify
type TaskTypeInput struct {
	// The TaskTypeCertify message indicates that a task is a certify task and all related details.
	TaskTypeCertify *TaskTypeCertifyInput `json:"certify,omitempty"`
	// The TaskTypeGrant message indicates that a task is a grant task and all related details.
	TaskTypeGrant *TaskTypeGrantInput `json:"grant,omitempty"`
	// The TaskTypeRevoke message indicates that a task is a revoke task and all related details.
	TaskTypeRevoke *TaskTypeRevokeInput `json:"revoke,omitempty"`
}

func (o *TaskTypeInput) GetTaskTypeCertify() *TaskTypeCertifyInput {
	if o == nil {
		return nil
	}
	return o.TaskTypeCertify
}

func (o *TaskTypeInput) GetTaskTypeGrant() *TaskTypeGrantInput {
	if o == nil {
		return nil
	}
	return o.TaskTypeGrant
}

func (o *TaskTypeInput) GetTaskTypeRevoke() *TaskTypeRevokeInput {
	if o == nil {
		return nil
	}
	return o.TaskTypeRevoke
}
