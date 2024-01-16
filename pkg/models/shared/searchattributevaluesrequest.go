// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SearchAttributeValuesRequest - Search Attributes by a few properties.
type SearchAttributeValuesRequest struct {
	// The attribute type ids for what type of attributes to search for.
	AttributeTypeIds []string `json:"attributeTypeIds,omitempty"`
	// Exclude attributes with these ids from the search results.
	ExcludeIds []string `json:"excludeIds,omitempty"`
	// Include attributes with these ids in the search results.
	Ids []string `json:"ids,omitempty"`
	// The pageSize where 0 <= pageSize <= 100. Values < 10 will be set to 10. A value of 0 returns the default page size (currently 25)
	PageSize *int `json:"pageSize,omitempty"`
	// The pageToken field.
	PageToken *string `json:"pageToken,omitempty"`
	// Query the attributes with a fuzzy search on display name and description.
	Query *string `json:"query,omitempty"`
	// Search for attributes with a case insensitive match on the attribute value which is the attribute name.
	Value *string `json:"value,omitempty"`
}

func (o *SearchAttributeValuesRequest) GetAttributeTypeIds() []string {
	if o == nil {
		return nil
	}
	return o.AttributeTypeIds
}

func (o *SearchAttributeValuesRequest) GetExcludeIds() []string {
	if o == nil {
		return nil
	}
	return o.ExcludeIds
}

func (o *SearchAttributeValuesRequest) GetIds() []string {
	if o == nil {
		return nil
	}
	return o.Ids
}

func (o *SearchAttributeValuesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *SearchAttributeValuesRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

func (o *SearchAttributeValuesRequest) GetQuery() *string {
	if o == nil {
		return nil
	}
	return o.Query
}

func (o *SearchAttributeValuesRequest) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
