// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListRolesResponse - The ListRolesResponse message contains a list of results and a nextPageToken if applicable.
type ListRolesResponse struct {
	// The list of results containing up to X results, where X is the page size defined in the request.
	List []Role `json:"list,omitempty"`
	// The nextPageToken is shown for the next page if the number of results is larger than the max page size.
	//  The server returns one page of results and the nextPageToken until all results are retreived.
	//  To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *ListRolesResponse) GetList() []Role {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *ListRolesResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
