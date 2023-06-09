// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppEntitlementsGetRequest struct {
	AppID string `pathParam:"style=simple,explode=false,name=app_id"`
	ID    string `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAppV1AppEntitlementsGetResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	C1APIAppV1GetAppEntitlementResponse *shared.C1APIAppV1GetAppEntitlementResponse
}
