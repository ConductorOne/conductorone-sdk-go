// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskServiceCreateRevokeResponse - The TaskServiceCreateRevokeResponse message.
type TaskServiceCreateRevokeResponse struct {
	//  Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	//
	TaskView *TaskView `json:"taskView,omitempty"`
	// The expanded field.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
}

func (o *TaskServiceCreateRevokeResponse) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *TaskServiceCreateRevokeResponse) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}
