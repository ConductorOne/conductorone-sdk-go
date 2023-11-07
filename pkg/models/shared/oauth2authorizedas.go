// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/utils"
	"time"
)

// OAuth2AuthorizedAsInput - OAuth2AuthorizedAs tracks the user that OAuthed with the connector.
type OAuth2AuthorizedAsInput struct {
}

// OAuth2AuthorizedAs tracks the user that OAuthed with the connector.
type OAuth2AuthorizedAs struct {
	// authEmail is the email of the user that authorized the connector using OAuth.
	AuthEmail    *string    `json:"authEmail,omitempty"`
	AuthorizedAt *time.Time `json:"authorizedAt,omitempty"`
}

func (o OAuth2AuthorizedAs) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OAuth2AuthorizedAs) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
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
