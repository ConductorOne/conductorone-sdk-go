// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppEntitlementRef - The AppEntitlementRef message.
type AppEntitlementRef struct {
	// The appId field.
	AppID *string `json:"appId,omitempty"`
	// The id field.
	ID *string `json:"id,omitempty"`
}

func (o *AppEntitlementRef) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppEntitlementRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
