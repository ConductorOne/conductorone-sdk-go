// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppResourceView - The app resource view returns an app resource with paths for items in the expand mask filled in when this response is returned and a request expand mask has "*" or "app_id" or "resource_type_id".
type AppResourceView struct {
	// The app resource message is a single resource that can have entitlements.
	AppResource *AppResource `json:"appResource,omitempty"`
	// JSONPATH expression indicating the location of the App object in the array
	AppPath *string `json:"appPath,omitempty"`
	// JSONPATH expression indicating the location of the Resource Type object in the array
	ResourceTypePath *string `json:"resourceTypePath,omitempty"`
}

func (o *AppResourceView) GetAppResource() *AppResource {
	if o == nil {
		return nil
	}
	return o.AppResource
}

func (o *AppResourceView) GetAppPath() *string {
	if o == nil {
		return nil
	}
	return o.AppPath
}

func (o *AppResourceView) GetResourceTypePath() *string {
	if o == nil {
		return nil
	}
	return o.ResourceTypePath
}
