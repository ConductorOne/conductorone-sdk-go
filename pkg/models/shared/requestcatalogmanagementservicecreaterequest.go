// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequestCatalogManagementServiceCreateRequest - Create a request catalog.
type RequestCatalogManagementServiceCreateRequest struct {
	// The RequestCatalogExpandMask includes the paths in the catalog view to expand in the return value of this call.
	RequestCatalogExpandMask *RequestCatalogExpandMask `json:"expandMask,omitempty"`
	// The description of the new request catalog.
	Description *string `json:"description,omitempty"`
	// The display name of the new request catalog.
	DisplayName *string `json:"displayName,omitempty"`
	// Whether or not the new catalog should be created as published.
	Published *bool `json:"published,omitempty"`
	// Whether or not the new catalog is visible to everyone by default.
	VisibleToEveryone *bool `json:"visibleToEveryone,omitempty"`
}

func (o *RequestCatalogManagementServiceCreateRequest) GetRequestCatalogExpandMask() *RequestCatalogExpandMask {
	if o == nil {
		return nil
	}
	return o.RequestCatalogExpandMask
}

func (o *RequestCatalogManagementServiceCreateRequest) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *RequestCatalogManagementServiceCreateRequest) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *RequestCatalogManagementServiceCreateRequest) GetPublished() *bool {
	if o == nil {
		return nil
	}
	return o.Published
}

func (o *RequestCatalogManagementServiceCreateRequest) GetVisibleToEveryone() *bool {
	if o == nil {
		return nil
	}
	return o.VisibleToEveryone
}
