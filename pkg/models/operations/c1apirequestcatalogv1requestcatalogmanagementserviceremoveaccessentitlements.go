// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest struct {
	RequestCatalogManagementServiceRemoveAccessEntitlementsRequest *shared.RequestCatalogManagementServiceRemoveAccessEntitlementsRequest `request:"mediaType=application/json"`
	CatalogID                                                      string                                                                 `pathParam:"style=simple,explode=false,name=catalog_id"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest) GetRequestCatalogManagementServiceRemoveAccessEntitlementsRequest() *shared.RequestCatalogManagementServiceRemoveAccessEntitlementsRequest {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceRemoveAccessEntitlementsRequest
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Empty response with a status code indicating success.
	RequestCatalogManagementServiceRemoveAccessEntitlementsResponse *shared.RequestCatalogManagementServiceRemoveAccessEntitlementsResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsResponse) GetRequestCatalogManagementServiceRemoveAccessEntitlementsResponse() *shared.RequestCatalogManagementServiceRemoveAccessEntitlementsResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceRemoveAccessEntitlementsResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAccessEntitlementsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
