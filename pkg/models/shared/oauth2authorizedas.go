// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// OAuth2AuthorizedAs - The OAuth2AuthorizedAs message.
type OAuth2AuthorizedAs struct {
	// The authEmail field.
	AuthEmail    *string    `json:"authEmail,omitempty"`
	AuthorizedAt *time.Time `json:"authorizedAt,omitempty"`
}

func (o *OAuth2AuthorizedAs) GetAuthEmail() *string {
	if o == nil {
		return nil
	}
	return o.AuthEmail
}

func (o *OAuth2AuthorizedAs) GetAuthorizedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AuthorizedAt
}