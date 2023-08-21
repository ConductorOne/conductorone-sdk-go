// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

// AppUserAppUserType - The appplication user type. Type can be user, system or service.
type AppUserAppUserType string

const (
	AppUserAppUserTypeAppUserTypeUnspecified    AppUserAppUserType = "APP_USER_TYPE_UNSPECIFIED"
	AppUserAppUserTypeAppUserTypeUser           AppUserAppUserType = "APP_USER_TYPE_USER"
	AppUserAppUserTypeAppUserTypeServiceAccount AppUserAppUserType = "APP_USER_TYPE_SERVICE_ACCOUNT"
	AppUserAppUserTypeAppUserTypeSystemAccount  AppUserAppUserType = "APP_USER_TYPE_SYSTEM_ACCOUNT"
)

func (e AppUserAppUserType) ToPointer() *AppUserAppUserType {
	return &e
}

func (e *AppUserAppUserType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "APP_USER_TYPE_UNSPECIFIED":
		fallthrough
	case "APP_USER_TYPE_USER":
		fallthrough
	case "APP_USER_TYPE_SERVICE_ACCOUNT":
		fallthrough
	case "APP_USER_TYPE_SYSTEM_ACCOUNT":
		*e = AppUserAppUserType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppUserAppUserType: %v", v)
	}
}

// AppUser - Application User that represents an account in the application.
type AppUser struct {
	// The satus of the applicaiton user.
	AppUserStatus *AppUserStatus `json:"status,omitempty"`
	// The ID of the application.
	AppID *string `json:"appId,omitempty"`
	// The appplication user type. Type can be user, system or service.
	AppUserType *AppUserAppUserType `json:"appUserType,omitempty"`
	CreatedAt   *time.Time          `json:"createdAt,omitempty"`
	DeletedAt   *time.Time          `json:"deletedAt,omitempty"`
	// The display name of the application user.
	DisplayName *string `json:"displayName,omitempty"`
	// The email field of the application user.
	Email *string `json:"email,omitempty"`
	// A unique idenditfier of the application user.
	ID *string `json:"id,omitempty"`
	// The conductor one user ID of the account owner.
	IdentityUserID *string                `json:"identityUserId,omitempty"`
	Profile        map[string]interface{} `json:"profile,omitempty"`
	UpdatedAt      *time.Time             `json:"updatedAt,omitempty"`
}

func (o *AppUser) GetAppUserStatus() *AppUserStatus {
	if o == nil {
		return nil
	}
	return o.AppUserStatus
}

func (o *AppUser) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *AppUser) GetAppUserType() *AppUserAppUserType {
	if o == nil {
		return nil
	}
	return o.AppUserType
}

func (o *AppUser) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AppUser) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *AppUser) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppUser) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AppUser) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AppUser) GetIdentityUserID() *string {
	if o == nil {
		return nil
	}
	return o.IdentityUserID
}

func (o *AppUser) GetProfile() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *AppUser) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
