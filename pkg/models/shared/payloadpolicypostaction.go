// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// PayloadPolicyPostActionExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type PayloadPolicyPostActionExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (p PayloadPolicyPostActionExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PayloadPolicyPostActionExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PayloadPolicyPostActionExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *PayloadPolicyPostActionExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The PayloadPolicyPostAction message.
type PayloadPolicyPostAction struct {
	// Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	TaskView *TaskView `json:"taskView,omitempty"`
	// List of serialized related objects.
	Expanded []PayloadPolicyPostActionExpanded `json:"expanded,omitempty"`
}

func (o *PayloadPolicyPostAction) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *PayloadPolicyPostAction) GetExpanded() []PayloadPolicyPostActionExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
