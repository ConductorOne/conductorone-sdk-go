// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The BundleAutomationRuleEntitlement message.
type BundleAutomationRuleEntitlement struct {
	// The entitlementRefs field.
	EntitlementRefs []AppEntitlementRef `json:"entitlementRefs,omitempty"`
}

func (o *BundleAutomationRuleEntitlement) GetEntitlementRefs() []AppEntitlementRef {
	if o == nil {
		return nil
	}
	return o.EntitlementRefs
}