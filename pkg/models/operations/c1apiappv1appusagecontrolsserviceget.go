// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppUsageControlsServiceGetRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
}

func (o *C1APIAppV1AppUsageControlsServiceGetRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

type C1APIAppV1AppUsageControlsServiceGetResponse struct {
	ContentType string
	//  The GetAppUsageControlsResponse message contains the retrieved AppUsageControls object.
	//
	GetAppUsageControlsResponse *shared.GetAppUsageControlsResponse
	StatusCode                  int
	RawResponse                 *http.Response
}

func (o *C1APIAppV1AppUsageControlsServiceGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppUsageControlsServiceGetResponse) GetGetAppUsageControlsResponse() *shared.GetAppUsageControlsResponse {
	if o == nil {
		return nil
	}
	return o.GetAppUsageControlsResponse
}

func (o *C1APIAppV1AppUsageControlsServiceGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppUsageControlsServiceGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
