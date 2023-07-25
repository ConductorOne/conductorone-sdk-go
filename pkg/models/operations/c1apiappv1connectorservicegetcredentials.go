// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceGetCredentialsRequest struct {
	AppID       string `pathParam:"style=simple,explode=false,name=app_id"`
	ConnectorID string `pathParam:"style=simple,explode=false,name=connector_id"`
	ID          string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1ConnectorServiceGetCredentialsRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1ConnectorServiceGetCredentialsRequest) GetConnectorID() string {
	if o == nil {
		return ""
	}
	return o.ConnectorID
}

func (o *C1APIAppV1ConnectorServiceGetCredentialsRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1ConnectorServiceGetCredentialsResponse struct {
	// Successful response
	ConnectorServiceGetCredentialsResponse *shared.ConnectorServiceGetCredentialsResponse
	ContentType                            string
	StatusCode                             int
	RawResponse                            *http.Response
}

func (o *C1APIAppV1ConnectorServiceGetCredentialsResponse) GetConnectorServiceGetCredentialsResponse() *shared.ConnectorServiceGetCredentialsResponse {
	if o == nil {
		return nil
	}
	return o.ConnectorServiceGetCredentialsResponse
}

func (o *C1APIAppV1ConnectorServiceGetCredentialsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1ConnectorServiceGetCredentialsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1ConnectorServiceGetCredentialsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
