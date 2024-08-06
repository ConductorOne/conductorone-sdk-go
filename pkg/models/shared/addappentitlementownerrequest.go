// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// AddAppEntitlementOwnerRequest - The request message for adding an app entitlement owner.
type AddAppEntitlementOwnerRequest struct {
	// The user_id field for the user to add as an owner of the app entitlement.
	UserID *string `json:"userId,omitempty"`
}

func (o *AddAppEntitlementOwnerRequest) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
