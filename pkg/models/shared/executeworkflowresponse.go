// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ExecuteWorkflowResponse message.
type ExecuteWorkflowResponse struct {
	// The executionId field.
	ExecutionID *int64 `integer:"string" json:"executionId,omitempty"`
}

func (o *ExecuteWorkflowResponse) GetExecutionID() *int64 {
	if o == nil {
		return nil
	}
	return o.ExecutionID
}
