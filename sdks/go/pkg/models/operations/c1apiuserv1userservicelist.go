// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIUserV1UserServiceListRequest struct {
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

type C1APIUserV1UserServiceListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The UserServiceListResponse message contains a list of results and a nextPageToken if applicable.
	UserServiceListResponse *shared.UserServiceListResponse
}