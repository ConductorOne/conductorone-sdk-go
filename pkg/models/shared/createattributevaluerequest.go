// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The CreateAttributeValueRequest message.
type CreateAttributeValueRequest struct {
	// The attributeTypeId field.
	AttributeTypeID *string `json:"attributeTypeId,omitempty"`
	// The value field.
	Value *string `json:"value,omitempty"`
}

func (o *CreateAttributeValueRequest) GetAttributeTypeID() *string {
	if o == nil {
		return nil
	}
	return o.AttributeTypeID
}

func (o *CreateAttributeValueRequest) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
