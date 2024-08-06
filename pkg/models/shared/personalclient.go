// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
	"time"
)

// The PersonalClient message contains information about a presonal client credential.
type PersonalClient struct {
	// If set, only allows the CIDRs in the array to use the credential.
	AllowSourceCidr []string `json:"allowSourceCidr,omitempty"`
	// The clientID of the credential.
	ClientID  *string    `json:"clientId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The display name of the personal client credential.
	DisplayName *string    `json:"displayName,omitempty"`
	ExpiresTime *time.Time `json:"expiresTime,omitempty"`
	// The unique ID of the personal client credential.
	ID         *string    `json:"id,omitempty"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	// scoped_roles provides a list of IAM Roles
	//  that this OAuth2 Client's API permissions
	//  are reduced to. The permissions granted to OAuth2 Client
	//  are AND'ed against the owning User's own permissions.
	ScopedRoles []string   `json:"scopedRoles,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	// The ID of the user that this credential is created for.
	UserID *string `json:"userId,omitempty"`
}

func (p PersonalClient) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PersonalClient) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PersonalClient) GetAllowSourceCidr() []string {
	if o == nil {
		return nil
	}
	return o.AllowSourceCidr
}

func (o *PersonalClient) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *PersonalClient) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PersonalClient) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *PersonalClient) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *PersonalClient) GetExpiresTime() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresTime
}

func (o *PersonalClient) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PersonalClient) GetLastUsedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastUsedAt
}

func (o *PersonalClient) GetScopedRoles() []string {
	if o == nil {
		return nil
	}
	return o.ScopedRoles
}

func (o *PersonalClient) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *PersonalClient) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}
