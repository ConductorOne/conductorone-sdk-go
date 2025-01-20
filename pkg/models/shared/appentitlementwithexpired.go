// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// The AppEntitlementWithExpired message.
type AppEntitlementWithExpired struct {
	// Application User that represents an account in the application.
	AppUser *AppUser `json:"appUser,omitempty"`
	// The User object provides all of the details for an user, as well as some configuration.
	User *User `json:"user,omitempty"`
	// The appEntitlementId field.
	AppEntitlementID *string `json:"appEntitlementId,omitempty"`
	// The appId field.
	AppID *string `json:"appId,omitempty"`
	// The appUserId field.
	AppUserID  *string    `json:"appUserId,omitempty"`
	Discovered *time.Time `json:"discovered,omitempty"`
	Expired    *time.Time `json:"expired,omitempty"`
	// The grantReasons field.
	GrantReasons []GrantReason `json:"grantReasons,omitempty"`
	// The grantSources field.
	GrantSources []AppEntitlementRef `json:"grantSources,omitempty"`
}

func (a AppEntitlementWithExpired) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppEntitlementWithExpired) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AppEntitlementWithExpired) GetAppUser() *AppUser {
	if o == nil {
		return nil
	}
	return o.AppUser
}

func (o *AppEntitlementWithExpired) GetUser() *User {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *AppEntitlementWithExpired) GetAppEntitlementID() *string {
	if o == nil {
		return nil
	}
	return o.AppEntitlementID
}

func (o *AppEntitlementWithExpired) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppEntitlementWithExpired) GetAppUserID() *string {
	if o == nil {
		return nil
	}
	return o.AppUserID
}

func (o *AppEntitlementWithExpired) GetDiscovered() *time.Time {
	if o == nil {
		return nil
	}
	return o.Discovered
}

func (o *AppEntitlementWithExpired) GetExpired() *time.Time {
	if o == nil {
		return nil
	}
	return o.Expired
}

func (o *AppEntitlementWithExpired) GetGrantReasons() []GrantReason {
	if o == nil {
		return nil
	}
	return o.GrantReasons
}

func (o *AppEntitlementWithExpired) GetGrantSources() []AppEntitlementRef {
	if o == nil {
		return nil
	}
	return o.GrantSources
}
