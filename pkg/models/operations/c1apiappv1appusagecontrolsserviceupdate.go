// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppUsageControlsServiceUpdateRequest struct {
	UpdateAppUsageControlsRequest *shared.UpdateAppUsageControlsRequest `request:"mediaType=application/json"`
	AppID                         string                                `pathParam:"style=simple,explode=false,name=app_id"`
}

func (o *C1APIAppV1AppUsageControlsServiceUpdateRequest) GetUpdateAppUsageControlsRequest() *shared.UpdateAppUsageControlsRequest {
	if o == nil {
		return nil
	}
	return o.UpdateAppUsageControlsRequest
}

func (o *C1APIAppV1AppUsageControlsServiceUpdateRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

type C1APIAppV1AppUsageControlsServiceUpdateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The UpdateAppUsageControlsResponse message contains the updated AppUsageControls object.
	UpdateAppUsageControlsResponse *shared.UpdateAppUsageControlsResponse
}

func (o *C1APIAppV1AppUsageControlsServiceUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppUsageControlsServiceUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppUsageControlsServiceUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppUsageControlsServiceUpdateResponse) GetUpdateAppUsageControlsResponse() *shared.UpdateAppUsageControlsResponse {
	if o == nil {
		return nil
	}
	return o.UpdateAppUsageControlsResponse
}
