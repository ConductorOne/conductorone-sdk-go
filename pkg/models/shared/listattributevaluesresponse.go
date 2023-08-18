// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListAttributeValuesResponse - ListAttributeValuesResponse is the response for listing attribute values for a given AttributeType.
type ListAttributeValuesResponse struct {
	// The list of AttributeValues.
	List []AttributeValue `json:"list,omitempty"`
	// The nextPageToken is shown for the next page if the number of results is larger than the max page size.
	//  The server returns one page of results and the nextPageToken until all results are retreived.
	//  To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The notificationToken field.
	NotificationToken *string `json:"notificationToken,omitempty"`
}

func (o *ListAttributeValuesResponse) GetList() []AttributeValue {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *ListAttributeValuesResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}

func (o *ListAttributeValuesResponse) GetNotificationToken() *string {
	if o == nil {
		return nil
	}
	return o.NotificationToken
}
