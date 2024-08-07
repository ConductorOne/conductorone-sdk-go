// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// SearchAppsRequest - Search Apps by a few properties.
type SearchAppsRequest struct {
	// A list of app IDs to restrict the search to.
	AppIds []string `json:"appIds,omitempty"`
	// Search for apps with a case insensitive match on the display name.
	DisplayName *string `json:"displayName,omitempty"`
	// A list of app IDs to remove from the results.
	ExcludeAppIds []string `json:"excludeAppIds,omitempty"`
	// Only return apps which are directories
	OnlyDirectories *bool `json:"onlyDirectories,omitempty"`
	// The pageSize where 0 <= pageSize <= 100. Values < 10 will be set to 10. A value of 0 returns the default page size (currently 25)
	PageSize *int `json:"pageSize,omitempty"`
	// The pageToken field.
	PageToken *string `json:"pageToken,omitempty"`
	// Query the apps with a fuzzy search on display name and description.
	Query *string `json:"query,omitempty"`
}

func (o *SearchAppsRequest) GetAppIds() []string {
	if o == nil {
		return nil
	}
	return o.AppIds
}

func (o *SearchAppsRequest) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *SearchAppsRequest) GetExcludeAppIds() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeAppIds
}

func (o *SearchAppsRequest) GetOnlyDirectories() *bool {
	if o == nil {
		return nil
	}
	return o.OnlyDirectories
}

func (o *SearchAppsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *SearchAppsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *SearchAppsRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}
