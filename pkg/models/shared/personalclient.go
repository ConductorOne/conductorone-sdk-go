// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// PersonalClient - The PersonalClient message.
type PersonalClient struct {
	// The allowSourceCidr field.
	AllowSourceCidr []string `json:"allowSourceCidr,omitempty"`
	// The clientId field.
	ClientID  *string    `json:"clientId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	// The displayName field.
	DisplayName *string    `json:"displayName,omitempty"`
	ExpiresTime *time.Time `json:"expiresTime,omitempty"`
	// The id field.
	ID         *string    `json:"id,omitempty"`
	LastUsedAt *time.Time `json:"lastUsedAt,omitempty"`
	//  scoped_roles provides a list of IAM Roles
	//  that this OAuth2 Client's API permissions
	//  are reduced to. The permissions granted to OAuth2 Client
	//  are AND'ed against the owning User's own permissions.
	//
	ScopedRoles []string   `json:"scopedRoles,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	// The userId field.
	UserID *string `json:"userId,omitempty"`
}