// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest struct {
	CatalogID string   `pathParam:"style=simple,explode=false,name=catalog_id"`
	PageSize  *float64 `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string  `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest) GetCatalogID() string {
	if o == nil {
		return ""
	}
	return o.CatalogID
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest) GetPageSize() *float64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The RequestCatalogManagementServiceListEntitlementsForAccessResponse message contains a list of results and a nextPageToken if applicable.
	RequestCatalogManagementServiceListEntitlementsForAccessResponse *shared.RequestCatalogManagementServiceListEntitlementsForAccessResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetRequestCatalogManagementServiceListEntitlementsForAccessResponse() *shared.RequestCatalogManagementServiceListEntitlementsForAccessResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceListEntitlementsForAccessResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsForAccessResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
