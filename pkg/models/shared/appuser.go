// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// AppUserType - The appplication user type. Type can be user, system or service.
type AppUserType string

const (
	AppUserTypeAppUserTypeUnspecified    AppUserType = "APP_USER_TYPE_UNSPECIFIED"
	AppUserTypeAppUserTypeUser           AppUserType = "APP_USER_TYPE_USER"
	AppUserTypeAppUserTypeServiceAccount AppUserType = "APP_USER_TYPE_SERVICE_ACCOUNT"
	AppUserTypeAppUserTypeSystemAccount  AppUserType = "APP_USER_TYPE_SYSTEM_ACCOUNT"
)

func (e AppUserType) ToPointer() *AppUserType {
	return &e
}

func (e *AppUserType) UnmarshalJSON(data []byte) error {
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
		*e = AppUserType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AppUserType: %v", v)
	}
}

type Three struct {
}

type ProfileType string

const (
	ProfileTypeStr        ProfileType = "str"
	ProfileTypeNumber     ProfileType = "number"
	ProfileTypeThree      ProfileType = "3"
	ProfileTypeArrayOfany ProfileType = "arrayOfany"
	ProfileTypeBoolean    ProfileType = "boolean"
)

type Profile struct {
	Str        *string
	Number     *float64
	Three      *Three
	ArrayOfany []interface{}
	Boolean    *bool

	Type ProfileType
}

func CreateProfileStr(str string) Profile {
	typ := ProfileTypeStr

	return Profile{
		Str:  &str,
		Type: typ,
	}
}

func CreateProfileNumber(number float64) Profile {
	typ := ProfileTypeNumber

	return Profile{
		Number: &number,
		Type:   typ,
	}
}

func CreateProfileThree(three Three) Profile {
	typ := ProfileTypeThree

	return Profile{
		Three: &three,
		Type:  typ,
	}
}

func CreateProfileArrayOfany(arrayOfany []interface{}) Profile {
	typ := ProfileTypeArrayOfany

	return Profile{
		ArrayOfany: arrayOfany,
		Type:       typ,
	}
}

func CreateProfileBoolean(boolean bool) Profile {
	typ := ProfileTypeBoolean

	return Profile{
		Boolean: &boolean,
		Type:    typ,
	}
}

func (u *Profile) UnmarshalJSON(data []byte) error {

	three := Three{}
	if err := utils.UnmarshalJSON(data, &three, "", true, true); err == nil {
		u.Three = &three
		u.Type = ProfileTypeThree
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ProfileTypeStr
		return nil
	}

	number := float64(0)
	if err := utils.UnmarshalJSON(data, &number, "", true, true); err == nil {
		u.Number = &number
		u.Type = ProfileTypeNumber
		return nil
	}

	arrayOfany := []interface{}{}
	if err := utils.UnmarshalJSON(data, &arrayOfany, "", true, true); err == nil {
		u.ArrayOfany = arrayOfany
		u.Type = ProfileTypeArrayOfany
		return nil
	}

	boolean := false
	if err := utils.UnmarshalJSON(data, &boolean, "", true, true); err == nil {
		u.Boolean = &boolean
		u.Type = ProfileTypeBoolean
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Profile) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.Number != nil {
		return utils.MarshalJSON(u.Number, "", true)
	}

	if u.Three != nil {
		return utils.MarshalJSON(u.Three, "", true)
	}

	if u.ArrayOfany != nil {
		return utils.MarshalJSON(u.ArrayOfany, "", true)
	}

	if u.Boolean != nil {
		return utils.MarshalJSON(u.Boolean, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// AppUser - Application User that represents an account in the application.
type AppUser struct {
	// The satus of the applicaiton user.
	AppUserStatus *AppUserStatus `json:"status,omitempty"`
	// The ID of the application.
	AppID *string `json:"appId,omitempty"`
	// The appplication user type. Type can be user, system or service.
	AppUserType *AppUserType `json:"appUserType,omitempty"`
	CreatedAt   *time.Time   `json:"createdAt,omitempty"`
	DeletedAt   *time.Time   `json:"deletedAt,omitempty"`
	// The display name of the application user.
	DisplayName *string `json:"displayName,omitempty"`
	// The email field of the application user.
	Email *string `json:"email,omitempty"`
	// The emails field of the application user.
	Emails []string `json:"emails,omitempty"`
	// A unique idenditfier of the application user.
	ID *string `json:"id,omitempty"`
	// The conductor one user ID of the account owner.
	IdentityUserID *string            `json:"identityUserId,omitempty"`
	Profile        map[string]Profile `json:"profile,omitempty"`
	UpdatedAt      *time.Time         `json:"updatedAt,omitempty"`
	// The username field of the application user.
	Username *string `json:"username,omitempty"`
	// The usernames field of the application user.
	Usernames []string `json:"usernames,omitempty"`
}

func (a AppUser) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AppUser) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
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

func (o *AppUser) GetAppUserType() *AppUserType {
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

func (o *AppUser) GetEmails() []string {
	if o == nil {
		return nil
	}
	return o.Emails
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

func (o *AppUser) GetProfile() map[string]Profile {
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

func (o *AppUser) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *AppUser) GetUsernames() []string {
	if o == nil {
		return nil
	}
	return o.Usernames
}

// AppUserInput - Application User that represents an account in the application.
type AppUserInput struct {
	// The satus of the applicaiton user.
	AppUserStatus *AppUserStatusInput `json:"status,omitempty"`
	// The appplication user type. Type can be user, system or service.
	AppUserType *AppUserType `json:"appUserType,omitempty"`
}

func (o *AppUserInput) GetAppUserStatus() *AppUserStatusInput {
	if o == nil {
		return nil
	}
	return o.AppUserStatus
}

func (o *AppUserInput) GetAppUserType() *AppUserType {
	if o == nil {
		return nil
	}
	return o.AppUserType
}
