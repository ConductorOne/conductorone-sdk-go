// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppEntitlementWithUserBindings - The AppEntitlementWithUserBindings message.
type AppEntitlementWithUserBindings struct {
	// The AppEntitlementView message.
	AppEntitlementView *AppEntitlementView `json:"entitlement,omitempty"`
	// The appEntitlementUserBindings field.
	AppEntitlementUserBindings []AppEntitlementUserBinding `json:"appEntitlementUserBindings,omitempty"`
}

func (o *AppEntitlementWithUserBindings) GetAppEntitlementView() *AppEntitlementView {
	if o == nil {
		return nil
	}
	return o.AppEntitlementView
}

func (o *AppEntitlementWithUserBindings) GetAppEntitlementUserBindings() []AppEntitlementUserBinding {
	if o == nil {
		return nil
	}
	return o.AppEntitlementUserBindings
}
