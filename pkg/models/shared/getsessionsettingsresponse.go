// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The GetSessionSettingsResponse message.
type GetSessionSettingsResponse struct {
	// The SessionSettings message.
	SessionSettings *SessionSettings `json:"sessionSettings,omitempty"`
}

func (o *GetSessionSettingsResponse) GetSessionSettings() *SessionSettings {
	if o == nil {
		return nil
	}
	return o.SessionSettings
}
