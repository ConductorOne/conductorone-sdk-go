// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// ProvisionInstanceState - The state field.
type ProvisionInstanceState string

const (
	ProvisionInstanceStateProvisionInstanceStateUnspecified                     ProvisionInstanceState = "PROVISION_INSTANCE_STATE_UNSPECIFIED"
	ProvisionInstanceStateProvisionInstanceStateInit                            ProvisionInstanceState = "PROVISION_INSTANCE_STATE_INIT"
	ProvisionInstanceStateProvisionInstanceStateCreateConnectorActionsForTarget ProvisionInstanceState = "PROVISION_INSTANCE_STATE_CREATE_CONNECTOR_ACTIONS_FOR_TARGET"
	ProvisionInstanceStateProvisionInstanceStateSendingNotifications            ProvisionInstanceState = "PROVISION_INSTANCE_STATE_SENDING_NOTIFICATIONS"
	ProvisionInstanceStateProvisionInstanceStateWaiting                         ProvisionInstanceState = "PROVISION_INSTANCE_STATE_WAITING"
	ProvisionInstanceStateProvisionInstanceStateDone                            ProvisionInstanceState = "PROVISION_INSTANCE_STATE_DONE"
)

func (e ProvisionInstanceState) ToPointer() *ProvisionInstanceState {
	return &e
}

func (e *ProvisionInstanceState) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PROVISION_INSTANCE_STATE_UNSPECIFIED":
		fallthrough
	case "PROVISION_INSTANCE_STATE_INIT":
		fallthrough
	case "PROVISION_INSTANCE_STATE_CREATE_CONNECTOR_ACTIONS_FOR_TARGET":
		fallthrough
	case "PROVISION_INSTANCE_STATE_SENDING_NOTIFICATIONS":
		fallthrough
	case "PROVISION_INSTANCE_STATE_WAITING":
		fallthrough
	case "PROVISION_INSTANCE_STATE_DONE":
		*e = ProvisionInstanceState(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProvisionInstanceState: %v", v)
	}
}

// ProvisionInstance - The ProvisionInstance message.
//
// This message contains a oneof named outcome. Only a single field of the following list may be set at a time:
//   - completed
//   - cancelled
//   - errored
//   - reassignedByError
type ProvisionInstance struct {
	// The CancelledAction message.
	CancelledAction *CancelledAction `json:"cancelled,omitempty"`
	// The CompletedAction message.
	CompletedAction *CompletedAction `json:"completed,omitempty"`
	// The ErroredAction message.
	ErroredAction *ErroredAction `json:"errored,omitempty"`
	// The Provision message.
	Provision *Provision `json:"provision,omitempty"`
	// The ReassignedByErrorAction message.
	ReassignedByErrorAction *ReassignedByErrorAction `json:"reassignedByError,omitempty"`
	// The notificationId field.
	NotificationID *string `json:"notificationId,omitempty"`
	// The state field.
	State *ProvisionInstanceState `json:"state,omitempty"`
}
