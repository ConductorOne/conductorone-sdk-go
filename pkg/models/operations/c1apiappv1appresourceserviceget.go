// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceServiceGetRequest struct {
	AppID             string `pathParam:"style=simple,explode=false,name=app_id"`
	AppResourceTypeID string `pathParam:"style=simple,explode=false,name=app_resource_type_id"`
	ID                string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1AppResourceServiceGetRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppResourceServiceGetRequest) GetAppResourceTypeID() string {
	if o == nil {
		return ""
	}
	return o.AppResourceTypeID
}

func (o *C1APIAppV1AppResourceServiceGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1AppResourceServiceGetResponse struct {
	// The app resource service get response contains the app resource view and array of expanded items indicated by the request's expand mask.
	AppResourceServiceGetResponse *shared.AppResourceServiceGetResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppResourceServiceGetResponse) GetAppResourceServiceGetResponse() *shared.AppResourceServiceGetResponse {
	if o == nil {
		return nil
	}
	return o.AppResourceServiceGetResponse
}

func (o *C1APIAppV1AppResourceServiceGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppResourceServiceGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppResourceServiceGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
