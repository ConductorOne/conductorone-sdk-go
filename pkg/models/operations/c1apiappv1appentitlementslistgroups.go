// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsListGroupsRequest struct {
	AppEntitlementID string   `pathParam:"style=simple,explode=false,name=app_entitlement_id"`
	AppID            string   `pathParam:"style=simple,explode=false,name=app_id"`
	PageSize         *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken        *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppEntitlementsListGroupsRequest) GetAppEntitlementID() string {
	if o == nil {
		return ""
	}
	return o.AppEntitlementID
}

func (o *C1APIAppV1AppEntitlementsListGroupsRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementsListGroupsRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppEntitlementsListGroupsRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppEntitlementsListGroupsResponse struct {
	ContentType string
	// The ListAppEntitlementGroupsResponse message contains a list of results and a nextPageToken if applicable.
	ListAppEntitlementGroupsResponse *shared.ListAppEntitlementGroupsResponse
	StatusCode                       int
	RawResponse                      *http.Response
}

func (o *C1APIAppV1AppEntitlementsListGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsListGroupsResponse) GetListAppEntitlementGroupsResponse() *shared.ListAppEntitlementGroupsResponse {
	if o == nil {
		return nil
	}
	return o.ListAppEntitlementGroupsResponse
}

func (o *C1APIAppV1AppEntitlementsListGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsListGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
