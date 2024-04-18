// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// Expanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type Expanded struct {
	// The type of the serialized message.
	AtType               *string                `json:"@type,omitempty"`
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
}

func (e Expanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *Expanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Expanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *Expanded) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The ConnectorServiceListResponse message contains a list of results and a nextPageToken if applicable
type ConnectorServiceListResponse struct {
	// List of serialized related objects
	Expanded []Expanded `json:"expanded,omitempty"`
	// The list of results containing up to X results, where X is the page size defined in the request
	List []ConnectorView `json:"list,omitempty"`
	// The nextPageToken is shown for the next page if the number of results is larger than the max page size. The server returns one page of results and the nextPageToken until all results are retreived. To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *ConnectorServiceListResponse) GetExpanded() []Expanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *ConnectorServiceListResponse) GetList() []ConnectorView {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *ConnectorServiceListResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
