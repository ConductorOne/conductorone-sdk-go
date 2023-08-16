// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// AttributeValue - AttributeValue is the value of an attribute of a defined type.
type AttributeValue struct {
	// The ID of the AttributeType that this AttributeValue belongs to.
	AttributeTypeID *string    `json:"attributeTypeId,omitempty"`
	CreatedAt       *time.Time `json:"createdAt,omitempty"`
	DeletedAt       *time.Time `json:"deletedAt,omitempty"`
	// The ID of the AttributeValue.
	ID        *string    `json:"id,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The value of the AttributeValue. This is the string that will be displayed to the user.
	Value *string `json:"value,omitempty"`
}

func (o *AttributeValue) GetAttributeTypeID() *string {
	if o == nil {
		return nil
	}
	return o.AttributeTypeID
}

func (o *AttributeValue) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AttributeValue) GetDeletedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *AttributeValue) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AttributeValue) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AttributeValue) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
