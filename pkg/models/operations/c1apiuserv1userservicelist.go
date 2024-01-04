// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type C1APIUserV1UserServiceListRequest struct {
	PageSize  *int    `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIUserV1UserServiceListRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIUserV1UserServiceListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIUserV1UserServiceListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The UserServiceListResponse message contains a list of results and a nextPageToken if applicable.
	UserServiceListResponse *shared.UserServiceListResponse
}

func (o *C1APIUserV1UserServiceListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIUserV1UserServiceListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIUserV1UserServiceListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIUserV1UserServiceListResponse) GetUserServiceListResponse() *shared.UserServiceListResponse {
	if o == nil {
		return nil
	}
	return o.UserServiceListResponse
}
