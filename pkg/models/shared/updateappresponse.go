// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateAppResponse - Returns the updated app's new values.
type UpdateAppResponse struct {
	// The App object provides all of the details for an app, as well as some configuration.
	App *App `json:"app,omitempty"`
}

func (o *UpdateAppResponse) GetApp() *App {
	if o == nil {
		return nil
	}
	return o.App
}
