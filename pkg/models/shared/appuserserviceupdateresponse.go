// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppUserServiceUpdateResponse - The AppUserServiceUpdateResponse message.
type AppUserServiceUpdateResponse struct {
	// The AppUserView contains an app user as well as paths for apps, identity users, and last usage in expanded arrays.
	AppUserView *AppUserView `json:"appUserView,omitempty"`
	// The expanded field.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
}

func (o *AppUserServiceUpdateResponse) GetAppUserView() *AppUserView {
	if o == nil {
		return nil
	}
	return o.AppUserView
}

func (o *AppUserServiceUpdateResponse) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}