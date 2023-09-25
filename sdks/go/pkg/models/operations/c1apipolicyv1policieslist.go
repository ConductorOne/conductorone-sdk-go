// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIPolicyV1PoliciesListRequest struct {
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIPolicyV1PoliciesListRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIPolicyV1PoliciesListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIPolicyV1PoliciesListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	ListPolicyResponse *shared.ListPolicyResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIPolicyV1PoliciesListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIPolicyV1PoliciesListResponse) GetListPolicyResponse() *shared.ListPolicyResponse {
	if o == nil {
		return nil
	}
	return o.ListPolicyResponse
}

func (o *C1APIPolicyV1PoliciesListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIPolicyV1PoliciesListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
