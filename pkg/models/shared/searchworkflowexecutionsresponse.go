// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The SearchWorkflowExecutionsResponse message.
type SearchWorkflowExecutionsResponse struct {
	// The list field.
	List []WorkflowExecution `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *SearchWorkflowExecutionsResponse) GetList() []WorkflowExecution {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *SearchWorkflowExecutionsResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
