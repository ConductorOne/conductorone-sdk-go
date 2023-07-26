// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceOwnersListRequest struct {
	AppID          string `pathParam:"style=simple,explode=false,name=app_id"`
	ResourceID     string `pathParam:"style=simple,explode=false,name=resource_id"`
	ResourceTypeID string `pathParam:"style=simple,explode=false,name=resource_type_id"`
}

func (o *C1APIAppV1AppResourceOwnersListRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppResourceOwnersListRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

func (o *C1APIAppV1AppResourceOwnersListRequest) GetResourceTypeID() string {
	if o == nil {
		return ""
	}
	return o.ResourceTypeID
}

type C1APIAppV1AppResourceOwnersListResponse struct {
	ContentType string
	//  The ListAppResourceOwnersResponse message contains a list of results and a nextPageToken if applicable
	//
	ListAppResourceOwnersResponse *shared.ListAppResourceOwnersResponse
	StatusCode                    int
	RawResponse                   *http.Response
}

func (o *C1APIAppV1AppResourceOwnersListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppResourceOwnersListResponse) GetListAppResourceOwnersResponse() *shared.ListAppResourceOwnersResponse {
	if o == nil {
		return nil
	}
	return o.ListAppResourceOwnersResponse
}

func (o *C1APIAppV1AppResourceOwnersListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppResourceOwnersListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
