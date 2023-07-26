// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// TaskRevokeSourceNonUsage -  The TaskRevokeSourceNonUsage message indicates that the source of the revoke task is due to the grant not being used.
type TaskRevokeSourceNonUsage struct {
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	LastLogin *time.Time `json:"lastLogin,omitempty"`
}

func (o *TaskRevokeSourceNonUsage) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *TaskRevokeSourceNonUsage) GetLastLogin() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastLogin
}
