// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateAppRequestInput - The UpdateAppRequest message contains the app to update and the fields to update.
type UpdateAppRequestInput struct {
	// The App object provides all of the details for an app, as well as some configuration.
	App        *AppInput `json:"app,omitempty"`
	UpdateMask *string   `json:"updateMask,omitempty"`
}
