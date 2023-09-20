// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// ErroredAction - The outcome of a provision instance that has errored.
type ErroredAction struct {
	// The description of a provision instance that has errored.
	Description *string `json:"description,omitempty"`
	// The error code of a provision instance that has errored. This is only PEC-1 for now, but more will be added in the future.
	ErrorCode *string    `json:"errorCode,omitempty"`
	ErroredAt *time.Time `json:"erroredAt,omitempty"`
}
