// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsDeleteRequest struct {
	DeleteAppEntitlementRequest *shared.DeleteAppEntitlementRequest `request:"mediaType=application/json"`
	AppID                       string                              `pathParam:"style=simple,explode=false,name=app_id"`
	ID                          string                              `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1AppEntitlementsDeleteRequest) GetDeleteAppEntitlementRequest() *shared.DeleteAppEntitlementRequest {
	if o == nil {
		return nil
	}
	return o.DeleteAppEntitlementRequest
}

func (o *C1APIAppV1AppEntitlementsDeleteRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppEntitlementsDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1AppEntitlementsDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	DeleteAppEntitlementResponse *shared.DeleteAppEntitlementResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppEntitlementsDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppEntitlementsDeleteResponse) GetDeleteAppEntitlementResponse() *shared.DeleteAppEntitlementResponse {
	if o == nil {
		return nil
	}
	return o.DeleteAppEntitlementResponse
}

func (o *C1APIAppV1AppEntitlementsDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppEntitlementsDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}