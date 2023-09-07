// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// The App object provides all of the details for an app, as well as some configuration.
type App struct {
	// The ID of the Account named by AccountName.
	AppAccountID *string `json:"appAccountId,omitempty"`
	// The AccountName of the app. For example, AWS is AccountID, Github is Org Name, and Okta is Okta Subdomain.
	AppAccountName *string `json:"appAccountName,omitempty"`
	// The ID of the Certify Policy associated with this App.
	CertifyPolicyID *string    `json:"certifyPolicyId,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	// The app's description.
	Description *string `json:"description,omitempty"`
	// The app's display name.
	DisplayName *string `json:"displayName,omitempty"`
	FieldMask   *string `json:"fieldMask,omitempty"`
	// The ID of the Grant Policy associated with this App.
	GrantPolicyID *string `json:"grantPolicyId,omitempty"`
	// The URL of an icon to display for the app.
	IconURL *string `json:"iconUrl,omitempty"`
	// The ID of the app.
	ID *string `json:"id,omitempty"`
	// The URL of a logo to display for the app.
	LogoURI *string `json:"logoUri,omitempty"`
	// The cost of an app per-seat, so that total cost can be calculated by the grant count.
	MonthlyCostUsd *float64 `json:"monthlyCostUsd,omitempty"`
	// The ID of the app that created this app, if any.
	ParentAppID *string `json:"parentAppId,omitempty"`
	// The ID of the Revoke Policy associated with this App.
	RevokePolicyID *string    `json:"revokePolicyId,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	// The number of users with grants to this app.
	UserCount *string `json:"userCount,omitempty"`
}

func (o *App) GetAppAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AppAccountID
}

func (o *App) GetAppAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AppAccountName
}

func (o *App) GetCertifyPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.CertifyPolicyID
}

func (o *App) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *App) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *App) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *App) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *App) GetFieldMask() *string {
	if o == nil {
		return nil
	}
	return o.FieldMask
}

func (o *App) GetGrantPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.GrantPolicyID
}

func (o *App) GetIconURL() *string {
	if o == nil {
		return nil
	}
	return o.IconURL
}

func (o *App) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *App) GetLogoURI() *string {
	if o == nil {
		return nil
	}
	return o.LogoURI
}

func (o *App) GetMonthlyCostUsd() *float64 {
	if o == nil {
		return nil
	}
	return o.MonthlyCostUsd
}

func (o *App) GetParentAppID() *string {
	if o == nil {
		return nil
	}
	return o.ParentAppID
}

func (o *App) GetRevokePolicyID() *string {
	if o == nil {
		return nil
	}
	return o.RevokePolicyID
}

func (o *App) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *App) GetUserCount() *string {
	if o == nil {
		return nil
	}
	return o.UserCount
}

// AppInput - The App object provides all of the details for an app, as well as some configuration.
type AppInput struct {
	// The ID of the Certify Policy associated with this App.
	CertifyPolicyID *string `json:"certifyPolicyId,omitempty"`
	// The app's description.
	Description *string `json:"description,omitempty"`
	// The app's display name.
	DisplayName *string `json:"displayName,omitempty"`
	// The ID of the Grant Policy associated with this App.
	GrantPolicyID *string `json:"grantPolicyId,omitempty"`
	// The URL of an icon to display for the app.
	IconURL *string `json:"iconUrl,omitempty"`
	// The cost of an app per-seat, so that total cost can be calculated by the grant count.
	MonthlyCostUsd *float64 `json:"monthlyCostUsd,omitempty"`
	// The ID of the Revoke Policy associated with this App.
	RevokePolicyID *string `json:"revokePolicyId,omitempty"`
}

func (o *AppInput) GetCertifyPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.CertifyPolicyID
}

func (o *AppInput) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AppInput) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *AppInput) GetGrantPolicyID() *string {
	if o == nil {
		return nil
	}
	return o.GrantPolicyID
}

func (o *AppInput) GetIconURL() *string {
	if o == nil {
		return nil
	}
	return o.IconURL
}

func (o *AppInput) GetMonthlyCostUsd() *float64 {
	if o == nil {
		return nil
	}
	return o.MonthlyCostUsd
}

func (o *AppInput) GetRevokePolicyID() *string {
	if o == nil {
		return nil
	}
	return o.RevokePolicyID
}
