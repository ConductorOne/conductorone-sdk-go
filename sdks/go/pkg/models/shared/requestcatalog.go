// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// RequestCatalog - The RequestCatalog is used for managing which entitlements are requestable, and who can request them.
type RequestCatalog struct {
	// An array of app entitlements that, if the user has, can view the contents of this catalog.
	AccessEntitlements []AppEntitlement `json:"accessEntitlements,omitempty"`
	// The Apps contained in this request catalog.
	AppIds    []string   `json:"appIds,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The id of the user this request catalog was created by.
	CreatedByUserID *string    `json:"createdByUserId,omitempty"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	// The description of the request catalog.
	Description *string `json:"description,omitempty"`
	// The display name of the request catalog.
	DisplayName *string `json:"displayName,omitempty"`
	// The id of the request catalog.
	ID *string `json:"id,omitempty"`
	// Whether or not this catalog is published.
	Published *bool      `json:"published,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// If this is true, the access entitlement requirement is ignored.
	VisibleToEveryone *bool `json:"visibleToEveryone,omitempty"`
}

// RequestCatalogInput - The RequestCatalog is used for managing which entitlements are requestable, and who can request them.
type RequestCatalogInput struct {
	// An array of app entitlements that, if the user has, can view the contents of this catalog.
	AccessEntitlements []AppEntitlementInput `json:"accessEntitlements,omitempty"`
	// The Apps contained in this request catalog.
	AppIds []string `json:"appIds,omitempty"`
	// The id of the user this request catalog was created by.
	CreatedByUserID *string `json:"createdByUserId,omitempty"`
	// The description of the request catalog.
	Description *string `json:"description,omitempty"`
	// The display name of the request catalog.
	DisplayName *string `json:"displayName,omitempty"`
	// The id of the request catalog.
	ID *string `json:"id,omitempty"`
	// Whether or not this catalog is published.
	Published *bool `json:"published,omitempty"`
	// If this is true, the access entitlement requirement is ignored.
	VisibleToEveryone *bool `json:"visibleToEveryone,omitempty"`
}