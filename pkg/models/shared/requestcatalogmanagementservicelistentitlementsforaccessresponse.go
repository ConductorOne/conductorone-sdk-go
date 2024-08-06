// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (r RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The RequestCatalogManagementServiceListEntitlementsForAccessResponse message contains a list of results and a nextPageToken if applicable.
type RequestCatalogManagementServiceListEntitlementsForAccessResponse struct {
	// List of serialized related objects.
	Expanded []RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded `json:"expanded,omitempty"`
	// The list of results containing up to X results, where X is the page size defined in the request.
	List []AppEntitlementView `json:"list,omitempty"`
	// The nextPageToken is shown for the next page if the number of results is larger than the max page size.
	//  The server returns one page of results and the nextPageToken until all results are retreived.
	//  To retrieve the next page, use the same request and append a pageToken field with the value of nextPageToken shown on the previous page.
	NextPageToken *string `json:"nextPageToken,omitempty"`
}

func (o *RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetExpanded() []RequestCatalogManagementServiceListEntitlementsForAccessResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}

func (o *RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetList() []AppEntitlementView {
	if o == nil {
		return nil
	}
	return o.List
}

func (o *RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}
