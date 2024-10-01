// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationRequest struct {
	ForceRunBundleAutomationRequest *shared.ForceRunBundleAutomationRequest `request:"mediaType=application/json"`
	RequestCatalogID                string                                  `pathParam:"style=simple,explode=false,name=request_catalog_id"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationRequest) GetForceRunBundleAutomationRequest() *shared.ForceRunBundleAutomationRequest {
	if o == nil {
		return nil
	}
	return o.ForceRunBundleAutomationRequest
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationRequest) GetRequestCatalogID() string {
	if o == nil {
		return ""
	}
	return o.RequestCatalogID
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	ForceRunBundleAutomationResponse *shared.ForceRunBundleAutomationResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationResponse) GetForceRunBundleAutomationResponse() *shared.ForceRunBundleAutomationResponse {
	if o == nil {
		return nil
	}
	return o.ForceRunBundleAutomationResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceForceRunBundleAutomationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
