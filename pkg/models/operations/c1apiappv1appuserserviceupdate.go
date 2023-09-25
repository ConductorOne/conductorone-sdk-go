// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppUserServiceUpdateRequest struct {
	AppUserServiceUpdateRequestInput *shared.AppUserServiceUpdateRequestInput `request:"mediaType=application/json"`
	AppUserAppID                     string                                   `pathParam:"style=simple,explode=false,name=app_user_app_id"`
	AppUserID                        string                                   `pathParam:"style=simple,explode=false,name=app_user_id"`
}

func (o *C1APIAppV1AppUserServiceUpdateRequest) GetAppUserServiceUpdateRequestInput() *shared.AppUserServiceUpdateRequestInput {
	if o == nil {
		return nil
	}
	return o.AppUserServiceUpdateRequestInput
}

func (o *C1APIAppV1AppUserServiceUpdateRequest) GetAppUserAppID() string {
	if o == nil {
		return ""
	}
	return o.AppUserAppID
}

func (o *C1APIAppV1AppUserServiceUpdateRequest) GetAppUserID() string {
	if o == nil {
		return ""
	}
	return o.AppUserID
}

type C1APIAppV1AppUserServiceUpdateResponse struct {
	// Successful response
	AppUserServiceUpdateResponse *shared.AppUserServiceUpdateResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppUserServiceUpdateResponse) GetAppUserServiceUpdateResponse() *shared.AppUserServiceUpdateResponse {
	if o == nil {
		return nil
	}
	return o.AppUserServiceUpdateResponse
}

func (o *C1APIAppV1AppUserServiceUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppUserServiceUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppUserServiceUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
