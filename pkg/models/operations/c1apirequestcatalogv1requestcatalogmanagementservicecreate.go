// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse struct {
	ContentType string
	//  The request catalog management service get response returns a request catalog view with the expanded items in the expanded array indicated by the expand mask in the request.
	//
	RequestCatalogManagementServiceGetResponse *shared.RequestCatalogManagementServiceGetResponse
	StatusCode                                 int
	RawResponse                                *http.Response
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse) GetRequestCatalogManagementServiceGetResponse() *shared.RequestCatalogManagementServiceGetResponse {
	if o == nil {
		return nil
	}
	return o.RequestCatalogManagementServiceGetResponse
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIRequestcatalogV1RequestCatalogManagementServiceCreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
