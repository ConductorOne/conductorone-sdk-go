// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceServiceUpdateRequest struct {
	AppResourceServiceUpdateRequest *shared.AppResourceServiceUpdateRequest `request:"mediaType=application/json"`
	AppID                           string                                  `pathParam:"style=simple,explode=false,name=app_id"`
	AppResourceTypeID               string                                  `pathParam:"style=simple,explode=false,name=app_resource_type_id"`
	ID                              string                                  `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1AppResourceServiceUpdateRequest) GetAppResourceServiceUpdateRequest() *shared.AppResourceServiceUpdateRequest {
	if o == nil {
		return nil
	}
	return o.AppResourceServiceUpdateRequest
}

func (o *C1APIAppV1AppResourceServiceUpdateRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppResourceServiceUpdateRequest) GetAppResourceTypeID() string {
	if o == nil {
		return ""
	}
	return o.AppResourceTypeID
}

func (o *C1APIAppV1AppResourceServiceUpdateRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1AppResourceServiceUpdateResponse struct {
	// Successful response
	AppResourceServiceUpdateResponse *shared.AppResourceServiceUpdateResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppResourceServiceUpdateResponse) GetAppResourceServiceUpdateResponse() *shared.AppResourceServiceUpdateResponse {
	if o == nil {
		return nil
	}
	return o.AppResourceServiceUpdateResponse
}

func (o *C1APIAppV1AppResourceServiceUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppResourceServiceUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppResourceServiceUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
