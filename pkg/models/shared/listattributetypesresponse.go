// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListAttributeTypesResponse - The ListAttributeTypesResponse message.
type ListAttributeTypesResponse struct {
	// The list field.
	List []AttributeType `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The notificationToken field.
	NotificationToken *string `json:"notificationToken,omitempty"`
}

func (o *ListAttributeTypesResponse) GetList() []AttributeType {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *ListAttributeTypesResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}

func (o *ListAttributeTypesResponse) GetNotificationToken() *string {
	if o == nil {
		return nil
	}
	return o.NotificationToken
}