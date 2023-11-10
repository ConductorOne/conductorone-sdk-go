// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsListForAppResourceRequest struct {
	AppID             string   `pathParam:"style=simple,explode=false,name=app_id"`
	AppResourceID     string   `pathParam:"style=simple,explode=false,name=app_resource_id"`
	AppResourceTypeID string   `pathParam:"style=simple,explode=false,name=app_resource_type_id"`
	PageSize          *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken         *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppEntitlementsListForAppResourceRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementsListForAppResourceRequest) GetAppResourceID() string {
	if o == nil {
		return ""
	}
	return o.AppResourceID
}

func (o *C1APIAppV1AppEntitlementsListForAppResourceRequest) GetAppResourceTypeID() string {
	if o == nil {
		return ""
	}
	return o.AppResourceTypeID
}

func (o *C1APIAppV1AppEntitlementsListForAppResourceRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppEntitlementsListForAppResourceRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppEntitlementsListForAppResourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The ListAppEntitlementsResponse message contains a list of results and a nextPageToken if applicable.
	ListAppEntitlementsResponse *shared.ListAppEntitlementsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppEntitlementsListForAppResourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsListForAppResourceResponse) GetListAppEntitlementsResponse() *shared.ListAppEntitlementsResponse {
	if o == nil {
		return nil
	}
	return o.ListAppEntitlementsResponse
}

func (o *C1APIAppV1AppEntitlementsListForAppResourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsListForAppResourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
