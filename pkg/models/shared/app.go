// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// App - The App message.
type App struct {
	// The appAccountId field.
	AppAccountID *string `json:"appAccountId,omitempty"`
	// The appAccountName field.
	AppAccountName *string `json:"appAccountName,omitempty"`
	// The certifyPolicyId field.
	CertifyPolicyID *string    `json:"certifyPolicyId,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	FieldMask   *string `json:"fieldMask,omitempty"`
	// The grantPolicyId field.
	GrantPolicyID *string `json:"grantPolicyId,omitempty"`
	// The iconUrl field.
	IconURL *string `json:"iconUrl,omitempty"`
	// The id field.
	ID *string `json:"id,omitempty"`
	// The logoUri field.
	LogoURI *string `json:"logoUri,omitempty"`
	// The monthlyCostUsd field.
	MonthlyCostUsd *float64 `json:"monthlyCostUsd,omitempty"`
	// The parentAppId field.
	ParentAppID *string `json:"parentAppId,omitempty"`
	// The revokePolicyId field.
	RevokePolicyID *string    `json:"revokePolicyId,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	// The userCount field.
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

// AppInput - The App message.
type AppInput struct {
	// The appAccountId field.
	AppAccountID *string `json:"appAccountId,omitempty"`
	// The appAccountName field.
	AppAccountName *string `json:"appAccountName,omitempty"`
	// The certifyPolicyId field.
	CertifyPolicyID *string `json:"certifyPolicyId,omitempty"`
	// The description field.
	Description *string `json:"description,omitempty"`
	// The displayName field.
	DisplayName *string `json:"displayName,omitempty"`
	FieldMask   *string `json:"fieldMask,omitempty"`
	// The grantPolicyId field.
	GrantPolicyID *string `json:"grantPolicyId,omitempty"`
	// The iconUrl field.
	IconURL *string `json:"iconUrl,omitempty"`
	// The id field.
	ID *string `json:"id,omitempty"`
	// The logoUri field.
	LogoURI *string `json:"logoUri,omitempty"`
	// The monthlyCostUsd field.
	MonthlyCostUsd *float64 `json:"monthlyCostUsd,omitempty"`
	// The parentAppId field.
	ParentAppID *string `json:"parentAppId,omitempty"`
	// The revokePolicyId field.
	RevokePolicyID *string `json:"revokePolicyId,omitempty"`
	// The userCount field.
	UserCount *string `json:"userCount,omitempty"`
}

func (o *AppInput) GetAppAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AppAccountID
}

func (o *AppInput) GetAppAccountName() *string {
	if o == nil {
		return nil
	}
	return o.AppAccountName
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

func (o *AppInput) GetFieldMask() *string {
	if o == nil {
		return nil
	}
	return o.FieldMask
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

func (o *AppInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AppInput) GetLogoURI() *string {
	if o == nil {
		return nil
	}
	return o.LogoURI
}

func (o *AppInput) GetMonthlyCostUsd() *float64 {
	if o == nil {
		return nil
	}
	return o.MonthlyCostUsd
}

func (o *AppInput) GetParentAppID() *string {
	if o == nil {
		return nil
	}
	return o.ParentAppID
}

func (o *AppInput) GetRevokePolicyID() *string {
	if o == nil {
		return nil
	}
	return o.RevokePolicyID
}

func (o *AppInput) GetUserCount() *string {
	if o == nil {
		return nil
	}
	return o.UserCount
}
