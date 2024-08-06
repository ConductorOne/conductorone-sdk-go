// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The AppReportServiceListResponse message contains a list of results and a nextPageToken if applicable.
type AppReportServiceListResponse struct {
	// The list of results containing up to X results, where X is the page size defined in the request.
	List []AppPopulationReport `json:"list,omitempty"`
	// The nextPageToken is shown for the next page if the number of results is larger than the max page size.
	//  The server returns one page of results and the nextPageToken until all results are retreived.
	//  To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *AppReportServiceListResponse) GetList() []AppPopulationReport {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *AppReportServiceListResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
