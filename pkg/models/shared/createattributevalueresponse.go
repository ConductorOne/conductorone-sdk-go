// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateAttributeValueResponse - The CreateAttributeValueResponse message.
type CreateAttributeValueResponse struct {
	// The AttributeValue message.
	AttributeValue *AttributeValue `json:"value,omitempty"`
}

func (o *CreateAttributeValueResponse) GetAttributeValue() *AttributeValue {
	if o == nil {
		return nil
	}
	return o.AttributeValue
}