// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetAppUsageControlsResponse - The GetAppUsageControlsResponse message.
type GetAppUsageControlsResponse struct {
	// The AppUsageControls message.
	AppUsageControls *AppUsageControls `json:"appUsageControls,omitempty"`
	// The hasUsageData field.
	HasUsageData *bool `json:"hasUsageData,omitempty"`
}

func (o *GetAppUsageControlsResponse) GetAppUsageControls() *AppUsageControls {
	if o == nil {
		return nil
	}
	return o.AppUsageControls
}

func (o *GetAppUsageControlsResponse) GetHasUsageData() *bool {
	if o == nil {
		return nil
	}
	return o.HasUsageData
}
