// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// RestartAction - The RestartAction message.
type RestartAction struct {
	// The oldPolicyStepId field.
	OldPolicyStepID *string    `json:"oldPolicyStepId,omitempty"`
	RestartedAt     *time.Time `json:"restartedAt,omitempty"`
	// The userId field.
	UserID *string `json:"userId,omitempty"`
}
