// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ExportServiceListEventsResponse message.
type ExportServiceListEventsResponse struct {
	// List contains an array of JSON OCSF events.
	List []map[string]any `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *ExportServiceListEventsResponse) GetList() []map[string]any {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *ExportServiceListEventsResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
