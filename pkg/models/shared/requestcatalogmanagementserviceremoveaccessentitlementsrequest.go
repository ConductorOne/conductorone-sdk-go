// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestCatalogManagementServiceRemoveAccessEntitlementsRequest -  The RequestCatalogManagementServiceRemoveAccessEntitlementsRequest message is used to remove access entitlements from a request catalog.
//
//	The access entitlements are used to determine which users can view the request catalog.
type RequestCatalogManagementServiceRemoveAccessEntitlementsRequest struct {
	//  The list of access entitlements to remove from the catalog.
	//
	AccessEntitlements []AppEntitlementRef `json:"accessEntitlements,omitempty"`
}

func (o *RequestCatalogManagementServiceRemoveAccessEntitlementsRequest) GetAccessEntitlements() []AppEntitlementRef {
	if o == nil {
		return nil
	}
	return o.AccessEntitlements
}
