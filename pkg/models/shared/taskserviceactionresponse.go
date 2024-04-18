// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// TaskServiceActionResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type TaskServiceActionResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string                `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
}

func (t TaskServiceActionResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskServiceActionResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskServiceActionResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *TaskServiceActionResponseExpanded) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The TaskServiceActionResponse message.
type TaskServiceActionResponse struct {
	// The expanded field.
	Expanded []TaskServiceActionResponseExpanded `json:"expanded,omitempty"`
	// Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	TaskView *TaskView `json:"taskView,omitempty"`
	// The ticketActionId field.
	TicketActionID *string `json:"ticketActionId,omitempty"`
}

func (o *TaskServiceActionResponse) GetExpanded() []TaskServiceActionResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *TaskServiceActionResponse) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *TaskServiceActionResponse) GetTicketActionID() *string {
	if o == nil {
		return nil
	}
	return o.TicketActionID
}
