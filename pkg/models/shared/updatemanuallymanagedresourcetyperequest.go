// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The UpdateManuallyManagedResourceTypeRequest message.
type UpdateManuallyManagedResourceTypeRequest struct {
	// The AppResourceType is referenced by an app entitlement defining its resource types. Commonly things like Group or Role.
	AppResourceType *AppResourceTypeInput `json:"appResourceType,omitempty"`
	UpdateMask      *string               `json:"updateMask,omitempty"`
}

func (o *UpdateManuallyManagedResourceTypeRequest) GetAppResourceType() *AppResourceTypeInput {
	if o == nil {
		return nil
	}
	return o.AppResourceType
}

func (o *UpdateManuallyManagedResourceTypeRequest) GetUpdateMask() *string {
	if o == nil {
		return nil
	}
	return o.UpdateMask
}