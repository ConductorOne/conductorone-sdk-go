// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ExportsSearchServiceSearchResponse message.
type ExportsSearchServiceSearchResponse struct {
	// The list field.
	List []Exporter `json:"list,omitempty"`
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *ExportsSearchServiceSearchResponse) GetList() []Exporter {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *ExportsSearchServiceSearchResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
