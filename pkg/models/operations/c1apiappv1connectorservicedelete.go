// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceDeleteRequest struct {
	ConnectorServiceDeleteRequest *shared.ConnectorServiceDeleteRequest `request:"mediaType=application/json"`
	AppID                         string                                `pathParam:"style=simple,explode=false,name=app_id"`
	ID                            string                                `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAppV1ConnectorServiceDeleteRequest) GetConnectorServiceDeleteRequest() *shared.ConnectorServiceDeleteRequest {
	if o == nil {
		return nil
	}
	return o.ConnectorServiceDeleteRequest
}

func (o *C1APIAppV1ConnectorServiceDeleteRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1ConnectorServiceDeleteRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAppV1ConnectorServiceDeleteResponse struct {
	//  Empty response body. Status code indicates success.
	//
	ConnectorServiceDeleteResponse *shared.ConnectorServiceDeleteResponse
	ContentType                    string
	StatusCode                     int
	RawResponse                    *http.Response
}

func (o *C1APIAppV1ConnectorServiceDeleteResponse) GetConnectorServiceDeleteResponse() *shared.ConnectorServiceDeleteResponse {
	if o == nil {
		return nil
	}
	return o.ConnectorServiceDeleteResponse
}

func (o *C1APIAppV1ConnectorServiceDeleteResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1ConnectorServiceDeleteResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1ConnectorServiceDeleteResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
