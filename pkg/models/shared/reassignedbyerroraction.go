// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// ReassignedByErrorAction - The ReassignedByErrorAction message.
type ReassignedByErrorAction struct {
	// The description field.
	Description *string `json:"description,omitempty"`
	// The errorCode field.
	ErrorCode *string `json:"errorCode,omitempty"`
	// The errorUserId field.
	ErrorUserID *string    `json:"errorUserId,omitempty"`
	ErroredAt   *time.Time `json:"erroredAt,omitempty"`
	// The newPolicyStepId field.
	NewPolicyStepID *string    `json:"newPolicyStepId,omitempty"`
	ReassignedAt    *time.Time `json:"reassignedAt,omitempty"`
}

func (o *ReassignedByErrorAction) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ReassignedByErrorAction) GetErrorCode() *string {
	if o == nil {
		return nil
	}
	return o.ErrorCode
}

func (o *ReassignedByErrorAction) GetErrorUserID() *string {
	if o == nil {
		return nil
	}
	return o.ErrorUserID
}

func (o *ReassignedByErrorAction) GetErroredAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ErroredAt
}

func (o *ReassignedByErrorAction) GetNewPolicyStepID() *string {
	if o == nil {
		return nil
	}
	return o.NewPolicyStepID
}

func (o *ReassignedByErrorAction) GetReassignedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ReassignedAt
}
