// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppsListRequest struct {
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppsListRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppsListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppsListResponse struct {
	ContentType string
	// The ListAppsResponse message contains a list of results and a nextPageToken if applicable.
	ListAppsResponse *shared.ListAppsResponse
	StatusCode       int
	RawResponse      *http.Response
}

func (o *C1APIAppV1AppsListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppsListResponse) GetListAppsResponse() *shared.ListAppsResponse {
	if o == nil {
		return nil
	}
	return o.ListAppsResponse
}

func (o *C1APIAppV1AppsListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppsListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
