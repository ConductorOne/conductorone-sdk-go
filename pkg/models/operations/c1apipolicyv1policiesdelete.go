// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIPolicyV1PoliciesDeleteRequest struct {
	DeletePolicyRequest *shared.DeletePolicyRequest `request:"mediaType=application/json"`
	ID                  string                      `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIPolicyV1PoliciesDeleteRequest) GetDeletePolicyRequest() *shared.DeletePolicyRequest {
	if o == nil {
		return nil
	}
	return o.DeletePolicyRequest
}

func (o *C1APIPolicyV1PoliciesDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIPolicyV1PoliciesDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Empty response with a status code indicating success.
	DeletePolicyResponse *shared.DeletePolicyResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIPolicyV1PoliciesDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIPolicyV1PoliciesDeleteResponse) GetDeletePolicyResponse() *shared.DeletePolicyResponse {
	if o == nil {
		return nil
	}
	return o.DeletePolicyResponse
}

func (o *C1APIPolicyV1PoliciesDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIPolicyV1PoliciesDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
