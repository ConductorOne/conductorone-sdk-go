// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppReportServiceListRequest struct {
	AppID     string  `pathParam:"style=simple,explode=false,name=app_id"`
	PageSize  *int    `queryParam:"style=form,explode=true,name=page_size"`
	PageToken *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAppV1AppReportServiceListRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppReportServiceListRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAppV1AppReportServiceListRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAppV1AppReportServiceListResponse struct {
	// The AppReportServiceListResponse message contains a list of results and a nextPageToken if applicable.
	AppReportServiceListResponse *shared.AppReportServiceListResponse
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAppV1AppReportServiceListResponse) GetAppReportServiceListResponse() *shared.AppReportServiceListResponse {
	if o == nil {
		return nil
	}
	return o.AppReportServiceListResponse
}

func (o *C1APIAppV1AppReportServiceListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppReportServiceListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppReportServiceListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
