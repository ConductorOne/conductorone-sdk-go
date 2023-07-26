// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceTypeServiceListRequest struct {
	AppID     string   `pathParam:"style=simple,explode=false,name=app_id"`
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppResourceTypeServiceListRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppResourceTypeServiceListRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppResourceTypeServiceListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppResourceTypeServiceListResponse struct {
	//  The AppResourceTypeServiceListResponse message contains a list of results and a nextPageToken if applicable.
	//
	AppResourceTypeServiceListResponse *shared.AppResourceTypeServiceListResponse
	ContentType                        string
	StatusCode                         int
	RawResponse                        *http.Response
}

func (o *C1APIAppV1AppResourceTypeServiceListResponse) GetAppResourceTypeServiceListResponse() *shared.AppResourceTypeServiceListResponse {
	if o == nil {
		return nil
	}
	return o.AppResourceTypeServiceListResponse
}

func (o *C1APIAppV1AppResourceTypeServiceListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppResourceTypeServiceListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppResourceTypeServiceListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
