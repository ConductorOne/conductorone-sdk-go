// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppResourceServiceGetResponse -  The app resource service get response contains the app resource view and array of expanded items indicated by the request's expand mask.
type AppResourceServiceGetResponse struct {
	//  The app resource view returns an app resource with paths for items in the expand mask filled in when this response is returned and a request expand mask has "*" or "app_id" or "resource_type_id".
	//
	AppResourceView *AppResourceView `json:"appResourceView,omitempty"`
	//  List of serialized related objects.
	//
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
}

func (o *AppResourceServiceGetResponse) GetAppResourceView() *AppResourceView {
	if o == nil {
		return nil
	}
	return o.AppResourceView
}

func (o *AppResourceServiceGetResponse) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}
