// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The UpdateSessionSettingsResponse message.
type UpdateSessionSettingsResponse struct {
	// The SessionSettings message.
	SessionSettings *SessionSettings `json:"sessionSettings,omitempty"`
}

func (o *UpdateSessionSettingsResponse) GetSessionSettings() *SessionSettings {
	if o == nil {
		return nil
	}
	return o.SessionSettings
}
