// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsListForAppUserRequest struct {
	AppID     string `pathParam:"style=simple,explode=false,name=app_id"`
	AppUserID string `pathParam:"style=simple,explode=false,name=app_user_id"`
}

func (o *C1APIAppV1AppEntitlementsListForAppUserRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementsListForAppUserRequest) GetAppUserID() string {
	if o == nil {
		return ""
	}
	return o.AppUserID
}

type C1APIAppV1AppEntitlementsListForAppUserResponse struct {
	ContentType string
	//  The ListAppEntitlementsResponse message contains a list of results and a nextPageToken if applicable.
	//
	ListAppEntitlementsResponse *shared.ListAppEntitlementsResponse
	StatusCode                  int
	RawResponse                 *http.Response
}

func (o *C1APIAppV1AppEntitlementsListForAppUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsListForAppUserResponse) GetListAppEntitlementsResponse() *shared.ListAppEntitlementsResponse {
	if o == nil {
		return nil
	}
	return o.ListAppEntitlementsResponse
}

func (o *C1APIAppV1AppEntitlementsListForAppUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsListForAppUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
