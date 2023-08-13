// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppsGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1AppsGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1AppsGetResponse struct {
	ContentType string
	// The GetAppResponse message contains the details of the requested app in the app field.
	GetAppResponse *shared.GetAppResponse
	StatusCode     int
	RawResponse    *http.Response
}

func (o *C1APIAppV1AppsGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppsGetResponse) GetGetAppResponse() *shared.GetAppResponse {
	if o == nil {
		return nil
	}
	return o.GetAppResponse
}

func (o *C1APIAppV1AppsGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppsGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
