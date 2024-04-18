// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsRequest struct {
	ID                              string                                  `pathParam:"style=simple,explode=false,name=id"`
	GetAccessRequestDefaultsRequest *shared.GetAccessRequestDefaultsRequest `request:"mediaType=application/json"`
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsRequest) GetGetAccessRequestDefaultsRequest() *shared.GetAccessRequestDefaultsRequest {
	if o == nil {
		return nil
	}
	return o.GetAccessRequestDefaultsRequest
}

type C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	AppAccessRequestDefaults *shared.AppAccessRequestDefaults
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppAccessRequestsDefaultsServiceGetAppAccessRequestsDefaultsResponse) GetAppAccessRequestDefaults() *shared.AppAccessRequestDefaults {
	if o == nil {
		return nil
	}
	return o.AppAccessRequestDefaults
}
