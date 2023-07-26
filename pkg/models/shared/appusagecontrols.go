// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AppUsageControls -  The AppUsageControls object describes some peripheral configuration for an app.
type AppUsageControls struct {
	//  The app that this object belongs to.
	//
	AppID *string `json:"appId,omitempty"`
	//  Whether or not to notify some if they have access to the app, but has not used it within a configurable amount of time.
	//
	Notify *bool `json:"notify,omitempty"`
	//  The duration in days after which we notify users of nonusage.
	//
	NotifyAfterDays *float64 `json:"notifyAfterDays,omitempty"`
	//  Whether or not to revoke a grant if they have access to the app, but has not used it within a configurable amount of time.
	//
	Revoke *bool `json:"revoke,omitempty"`
	//  The duration in days after which we revoke users that have not used that grant.
	//
	RevokeAfterDays *float64 `json:"revokeAfterDays,omitempty"`
}

func (o *AppUsageControls) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppUsageControls) GetNotify() *bool {
	if o == nil {
		return nil
	}
	return o.Notify
}

func (o *AppUsageControls) GetNotifyAfterDays() *float64 {
	if o == nil {
		return nil
	}
	return o.NotifyAfterDays
}

func (o *AppUsageControls) GetRevoke() *bool {
	if o == nil {
		return nil
	}
	return o.Revoke
}

func (o *AppUsageControls) GetRevokeAfterDays() *float64 {
	if o == nil {
		return nil
	}
	return o.RevokeAfterDays
}
