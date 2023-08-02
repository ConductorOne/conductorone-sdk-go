// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppsCreateResponse struct {
	ContentType string
	// Returns the new app's values.
	CreateAppResponse *shared.CreateAppResponse
	StatusCode        int
	RawResponse       *http.Response
}

func (o *C1APIAppV1AppsCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppsCreateResponse) GetCreateAppResponse() *shared.CreateAppResponse {
	if o == nil {
		return nil
	}
	return o.CreateAppResponse
}

func (o *C1APIAppV1AppsCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppsCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
