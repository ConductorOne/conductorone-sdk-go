// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// Payload - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type Payload struct {
	// The type of the serialized message.
	AtType               *string                `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
}

func (p Payload) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *Payload) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Payload) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *Payload) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The WebhookSpec message.
type WebhookSpec struct {
	// The destination field.
	Destination *string `json:"destination,omitempty"`
	// Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
	Payload *Payload `json:"payload,omitempty"`
}

func (o *WebhookSpec) GetDestination() *string {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *WebhookSpec) GetPayload() *Payload {
	if o == nil {
		return nil
	}
	return o.Payload
}