// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// CancelledAction - The outcome of a provision instance that is cancelled.
type CancelledAction struct {
	CancelledAt *time.Time `json:"cancelledAt,omitempty"`
	// The userID, usually the system, that cancells a provision instance.
	CancelledByUserID *string `json:"cancelledByUserId,omitempty"`
}

func (o *CancelledAction) GetCancelledAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CancelledAt
}

func (o *CancelledAction) GetCancelledByUserID() *string {
	if o == nil {
		return nil
	}
	return o.CancelledByUserID
}
