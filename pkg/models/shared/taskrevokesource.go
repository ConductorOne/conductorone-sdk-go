// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskRevokeSource - The TaskRevokeSource message indicates the source of the revoke task is one of expired, nonUsage, request, or review.
//
// This message contains a oneof named origin. Only a single field of the following list may be set at a time:
//   - review
//   - request
//   - expired
//   - nonUsage
type TaskRevokeSource struct {
	// The TaskRevokeSourceExpired message indicates that the source of the revoke task is due to a grant expiring.
	TaskRevokeSourceExpired *TaskRevokeSourceExpired `json:"expired,omitempty"`
	// The TaskRevokeSourceNonUsage message indicates that the source of the revoke task is due to the grant not being used.
	TaskRevokeSourceNonUsage *TaskRevokeSourceNonUsage `json:"nonUsage,omitempty"`
	// The TaskRevokeSourceRequest message indicates that the source of the revoke task was a request.
	TaskRevokeSourceRequest *TaskRevokeSourceRequest `json:"request,omitempty"`
	// The TaskRevokeSourceReview message tracks which access review was the source of the specificed revoke ticket.
	TaskRevokeSourceReview *TaskRevokeSourceReview `json:"review,omitempty"`
}

func (o *TaskRevokeSource) GetTaskRevokeSourceExpired() *TaskRevokeSourceExpired {
	if o == nil {
		return nil
	}
	return o.TaskRevokeSourceExpired
}

func (o *TaskRevokeSource) GetTaskRevokeSourceNonUsage() *TaskRevokeSourceNonUsage {
	if o == nil {
		return nil
	}
	return o.TaskRevokeSourceNonUsage
}

func (o *TaskRevokeSource) GetTaskRevokeSourceRequest() *TaskRevokeSourceRequest {
	if o == nil {
		return nil
	}
	return o.TaskRevokeSourceRequest
}

func (o *TaskRevokeSource) GetTaskRevokeSourceReview() *TaskRevokeSourceReview {
	if o == nil {
		return nil
	}
	return o.TaskRevokeSourceReview
}
