// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestCatalogManagementServiceAddAccessEntitlementsRequest - The RequestCatalogManagementServiceAddAccessEntitlementsRequest message.
type RequestCatalogManagementServiceAddAccessEntitlementsRequest struct {
	// The accessEntitlements field.
	AccessEntitlements []AppEntitlementRef `json:"accessEntitlements,omitempty"`
}

func (o *RequestCatalogManagementServiceAddAccessEntitlementsRequest) GetAccessEntitlements() []AppEntitlementRef {
	if o == nil {
		return nil
	}
	return o.AccessEntitlements
}
