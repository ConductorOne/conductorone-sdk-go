// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogRequest struct {
	CatalogID string `pathParam:"style=simple,explode=false,name=catalog_id"`
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceListEntitlementsPerCatalogResponse struct {
	ContentType string
	// Successful response
	RequestCatalogManagementServiceListEntitlementsPerCatalogResponse *shared.RequestCatalogManagementServiceListEntitlementsPerCatalogResponse
	StatusCode                                                        int
	RawResponse                                                       *http.Response
}
