// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest struct {
	ID                                           string                                               `pathParam:"style=simple,explode=false,name=id"`
	RequestCatalogManagementServiceDeleteRequest *shared.RequestCatalogManagementServiceDeleteRequest `request:"mediaType=application/json"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteRequest) GetRequestCatalogManagementServiceDeleteRequest() *shared.RequestCatalogManagementServiceDeleteRequest {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceDeleteRequest
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Empty response with a status code indicating success.
	RequestCatalogManagementServiceDeleteResponse *shared.RequestCatalogManagementServiceDeleteResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceDeleteResponse) GetRequestCatalogManagementServiceDeleteResponse() *shared.RequestCatalogManagementServiceDeleteResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceDeleteResponse
}
