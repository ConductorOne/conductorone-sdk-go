// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceRevokeCredentialRequest struct {
	ConnectorServiceRevokeCredentialRequest *shared.ConnectorServiceRevokeCredentialRequest `request:"mediaType=application/json"`
	AppID                                   string                                          `pathParam:"style=simple,explode=false,name=app_id"`
	ConnectorID                             string                                          `pathParam:"style=simple,explode=false,name=connector_id"`
	ID                                      string                                          `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1ConnectorServiceRevokeCredentialRequest) GetConnectorServiceRevokeCredentialRequest() *shared.ConnectorServiceRevokeCredentialRequest {
	if o == nil {
		return nil
	}
	return o.ConnectorServiceRevokeCredentialRequest
}

func (o *C1APIAppV1ConnectorServiceRevokeCredentialRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1ConnectorServiceRevokeCredentialRequest) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *C1APIAppV1ConnectorServiceRevokeCredentialRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1ConnectorServiceRevokeCredentialResponse struct {
	// Empty response body. Status code indicates success.
	ConnectorServiceRevokeCredentialResponse *shared.ConnectorServiceRevokeCredentialResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1ConnectorServiceRevokeCredentialResponse) GetConnectorServiceRevokeCredentialResponse() *shared.ConnectorServiceRevokeCredentialResponse {
	if o == nil {
		return nil
	}
	return o.ConnectorServiceRevokeCredentialResponse
}

func (o *C1APIAppV1ConnectorServiceRevokeCredentialResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1ConnectorServiceRevokeCredentialResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1ConnectorServiceRevokeCredentialResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
