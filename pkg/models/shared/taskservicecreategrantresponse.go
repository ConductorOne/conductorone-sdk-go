// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/utils"
)

// TaskServiceCreateGrantResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type TaskServiceCreateGrantResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string                `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
}

func (t TaskServiceCreateGrantResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskServiceCreateGrantResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskServiceCreateGrantResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *TaskServiceCreateGrantResponseExpanded) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The TaskServiceCreateGrantResponse returns a task view which has a task including JSONPATHs to the expanded items in the expanded array.
type TaskServiceCreateGrantResponse struct {
	// Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	TaskView *TaskView `json:"taskView,omitempty"`
	// List of serialized related objects.
	Expanded []TaskServiceCreateGrantResponseExpanded `json:"expanded,omitempty"`
}

func (o *TaskServiceCreateGrantResponse) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *TaskServiceCreateGrantResponse) GetExpanded() []TaskServiceCreateGrantResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
