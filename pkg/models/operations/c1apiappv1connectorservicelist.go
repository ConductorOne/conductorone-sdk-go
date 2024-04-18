// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1ConnectorServiceListRequest struct {
	AppID     string  `pathParam:"style=simple,explode=false,name=app_id"`
	PageSize  *int    `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1ConnectorServiceListRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1ConnectorServiceListRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1ConnectorServiceListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1ConnectorServiceListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The ConnectorServiceListResponse message contains a list of results and a nextPageToken if applicable
	ConnectorServiceListResponse *shared.ConnectorServiceListResponse
}

func (o *C1APIAppV1ConnectorServiceListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1ConnectorServiceListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1ConnectorServiceListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1ConnectorServiceListResponse) GetConnectorServiceListResponse() *shared.ConnectorServiceListResponse {
	if o == nil {
		return nil
	}
	return o.ConnectorServiceListResponse
}
