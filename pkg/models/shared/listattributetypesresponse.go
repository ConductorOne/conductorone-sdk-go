// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListAttributeTypesResponse - ListAttributeTypesResponse is the response for listing attribute types.
type ListAttributeTypesResponse struct {
	// The list of AttributeTypes.
	List []AttributeType `json:"list,omitempty"`
	// The nextPageToken is shown for the next page if the number of results is larger than the max page size.
	//  The server returns one page of results and the nextPageToken until all results are retreived.
	//  To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
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
