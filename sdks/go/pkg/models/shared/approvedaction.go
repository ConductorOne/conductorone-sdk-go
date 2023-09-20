// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// ApprovedAction - The approved action indicates that the approvalinstance had an outcome of approved.
type ApprovedAction struct {
	ApprovedAt *time.Time `json:"approvedAt,omitempty"`
	// The entitlements that were approved. This will only ever be a list of one entitlement.
	Entitlements []AppEntitlementReference `json:"entitlements,omitempty"`
	// The UserID that approved this step.
	UserID *string `json:"userId,omitempty"`
}
