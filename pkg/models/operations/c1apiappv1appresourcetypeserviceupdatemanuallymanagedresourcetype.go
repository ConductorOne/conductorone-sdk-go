// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeRequest struct {
	UpdateManuallyManagedResourceTypeRequest *shared.UpdateManuallyManagedResourceTypeRequest `request:"mediaType=application/json"`
	AppID                                    string                                           `pathParam:"style=simple,explode=false,name=app_id"`
	ID                                       string                                           `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeRequest) GetUpdateManuallyManagedResourceTypeRequest() *shared.UpdateManuallyManagedResourceTypeRequest {
	if o == nil {
		return nil
	}
	return o.UpdateManuallyManagedResourceTypeRequest
}

func (o *C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	UpdateManuallyManagedResourceTypeResponse *shared.UpdateManuallyManagedResourceTypeResponse
}

func (o *C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppResourceTypeServiceUpdateManuallyManagedResourceTypeResponse) GetUpdateManuallyManagedResourceTypeResponse() *shared.UpdateManuallyManagedResourceTypeResponse {
	if o == nil {
		return nil
	}
	return o.UpdateManuallyManagedResourceTypeResponse
}
