// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceRevokeCredentialRequest struct {
	ConnectorServiceRevokeCredentialRequest *shared.ConnectorServiceRevokeCredentialRequest `request:"mediaType=application/json"`
	AppID                                   string                                          `pathParam:"style=simple,explode=false,name=app_id"`
	ConnectorID                             string                                          `pathParam:"style=simple,explode=false,name=connector_id"`
	ID                                      string                                          `pathParam:"style=simple,explode=false,name=id"`
}

type C1APIAppV1ConnectorServiceRevokeCredentialResponse struct {
	// Empty response body. Status code indicates success.
	ConnectorServiceRevokeCredentialResponse *shared.ConnectorServiceRevokeCredentialResponse
	ContentType                              string
	StatusCode                               int
	RawResponse                              *http.Response
}