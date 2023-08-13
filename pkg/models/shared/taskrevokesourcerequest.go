// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskRevokeSourceRequest - The TaskRevokeSourceRequest message indicates that the source of the revoke task was a request.
type TaskRevokeSourceRequest struct {
	// The ID of the user who initiated the revoke request.
	RequestUserID *string `json:"requestUserId,omitempty"`
}

func (o *TaskRevokeSourceRequest) GetRequestUserID() *string {
	if o == nil {
		return nil
	}
	return o.RequestUserID
}
