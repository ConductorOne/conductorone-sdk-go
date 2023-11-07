// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/utils"
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

func (a ApprovedAction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *ApprovedAction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ApprovedAction) GetApprovedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ApprovedAt
}

func (o *ApprovedAction) GetEntitlements() []AppEntitlementReference {
	if o == nil {
		return nil
	}
	return o.Entitlements
}

func (o *ApprovedAction) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
