// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListAppUsersForIdentityWithGrantResponse - The ListAppUsersForIdentityWithGrantResponse message.
type ListAppUsersForIdentityWithGrantResponse struct {
	// The list of app users that may also have grant information.
	Bindings []AppEntitlementUserBinding `json:"bindings,omitempty"`
}
