// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The UpdateAppRequest message contains the app to update and the fields to update.
type UpdateAppRequest struct {
	// The App object provides all of the details for an app, as well as some configuration.
	App        *AppInput `json:"app,omitempty"`
	UpdateMask *string   `json:"updateMask,omitempty"`
}

func (o *UpdateAppRequest) GetApp() *AppInput {
	if o == nil {
		return nil
	}
	return o.App
}

func (o *UpdateAppRequest) GetUpdateMask() *string {
	if o == nil {
		return nil
	}
	return o.UpdateMask
}
