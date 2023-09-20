// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// The AppEntitlementSearchServiceSearchResponse message.
type AppEntitlementSearchServiceSearchResponse struct {
	// Indicates one value of a facet.
	Facets *Facets `json:"facets,omitempty"`
	// List of related objects.
	Expanded []map[string]interface{} `json:"expanded,omitempty"`
	// List of app entitlement view objects.
	List []AppEntitlementView `json:"list,omitempty"`
	// The nextPageToken is shown for the next page if the number of results is larger than the max page size. The server returns one page of results and the nextPageToken until all results are retreived. To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *AppEntitlementSearchServiceSearchResponse) GetFacets() *Facets {
	if o == nil {
		return nil
	}
	return o.Facets
}

func (o *AppEntitlementSearchServiceSearchResponse) GetExpanded() []map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *AppEntitlementSearchServiceSearchResponse) GetList() []AppEntitlementView {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *AppEntitlementSearchServiceSearchResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
