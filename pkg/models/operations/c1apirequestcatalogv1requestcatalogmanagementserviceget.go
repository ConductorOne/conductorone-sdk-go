// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceGetRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIRequestcatalogV1RequestCatalogManagementServiceGetResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// The request catalog management service get response returns a request catalog view with the expanded items in the expanded array indicated by the expand mask in the request.
	RequestCatalogManagementServiceGetResponse *shared.RequestCatalogManagementServiceGetResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceGetResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceGetResponse) GetRequestCatalogManagementServiceGetResponse() *shared.RequestCatalogManagementServiceGetResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceGetResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceGetResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceGetResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
