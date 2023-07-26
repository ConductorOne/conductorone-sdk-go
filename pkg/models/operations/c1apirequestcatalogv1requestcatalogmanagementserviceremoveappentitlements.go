// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsRequest struct {
	RequestCatalogManagementServiceRemoveAppEntitlementsRequest *shared.RequestCatalogManagementServiceRemoveAppEntitlementsRequest `request:"mediaType=application/json"`
	CatalogID                                                   string                                                              `pathParam:"style=simple,explode=false,name=catalog_id"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsRequest) GetRequestCatalogManagementServiceRemoveAppEntitlementsRequest() *shared.RequestCatalogManagementServiceRemoveAppEntitlementsRequest {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceRemoveAppEntitlementsRequest
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsResponse struct {
	ContentType string
	//  Empty response with a status code indicating success
	//
	RequestCatalogManagementServiceRemoveAppEntitlementsResponse *shared.RequestCatalogManagementServiceRemoveAppEntitlementsResponse
	StatusCode                                                   int
	RawResponse                                                  *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsResponse) GetRequestCatalogManagementServiceRemoveAppEntitlementsResponse() *shared.RequestCatalogManagementServiceRemoveAppEntitlementsResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceRemoveAppEntitlementsResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceRemoveAppEntitlementsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
