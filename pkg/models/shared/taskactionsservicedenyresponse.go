// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// The TaskActionsServiceDenyResponse returns a task view with paths indicating the location of expanded items in the array.
type TaskActionsServiceDenyResponse struct {
	// Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	TaskView *TaskView `json:"taskView,omitempty"`
	// List of serialized related objects.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
	// The ID of the ticket (task) deny action created by this request.
	TicketActionID *string `json:"ticketActionId,omitempty"`
}

func (o *TaskActionsServiceDenyResponse) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *TaskActionsServiceDenyResponse) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *TaskActionsServiceDenyResponse) GetTicketActionID() *string {
	if o == nil {
		return nil
	}
	return o.TicketActionID
}
