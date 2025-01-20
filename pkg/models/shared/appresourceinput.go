// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// AppResourceInput - The app resource message is a single resource that can have entitlements.
type AppResourceInput struct {
	// The app that this resource belongs to.
	AppID *string `json:"appId,omitempty"`
	// The resource type that this resource is.
	AppResourceTypeID *string `json:"appResourceTypeId,omitempty"`
	// A custom description that can be set for a resource.
	CustomDescription *string `json:"customDescription,omitempty"`
	// The display name for this resource.
	DisplayName *string `json:"displayName,omitempty"`
	// The number of grants to this resource.
	GrantCount *int64 `integer:"string" json:"grantCount,omitempty"`
	// The id of the resource.
	ID *string `json:"id,omitempty"`
	// The parent resource id, if this resource is a child of another resource.
	ParentAppResourceID *string `json:"parentAppResourceId,omitempty"`
	// The parent resource type id, if this resource is a child of another resource.
	ParentAppResourceTypeID *string `json:"parentAppResourceTypeId,omitempty"`
}

func (o *AppResourceInput) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppResourceInput) GetAppResourceTypeID() *string {
	if o == nil {
		return nil
	}
	return o.AppResourceTypeID
}

func (o *AppResourceInput) GetCustomDescription() *string {
	if o == nil {
		return nil
	}
	return o.CustomDescription
}

func (o *AppResourceInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppResourceInput) GetGrantCount() *int64 {
	if o == nil {
		return nil
	}
	return o.GrantCount
}

func (o *AppResourceInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AppResourceInput) GetParentAppResourceID() *string {
	if o == nil {
		return nil
	}
	return o.ParentAppResourceID
}

func (o *AppResourceInput) GetParentAppResourceTypeID() *string {
	if o == nil {
		return nil
	}
	return o.ParentAppResourceTypeID
}