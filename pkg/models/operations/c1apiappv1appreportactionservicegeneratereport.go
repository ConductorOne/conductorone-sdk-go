// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAppV1AppReportActionServiceGenerateReportRequest struct {
	AppID                                  string                                         `pathParam:"style=simple,explode=false,name=app_id"`
	AppActionsServiceGenerateReportRequest *shared.AppActionsServiceGenerateReportRequest `request:"mediaType=application/json"`
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportRequest) GetAppID() string {
	if o == nil {
		return ""
	}
	return o.AppID
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportRequest) GetAppActionsServiceGenerateReportRequest() *shared.AppActionsServiceGenerateReportRequest {
	if o == nil {
		return nil
	}
	return o.AppActionsServiceGenerateReportRequest
}

type C1APIAppV1AppReportActionServiceGenerateReportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Empty response body. Status code indicates success.
	AppActionsServiceGenerateReportResponse *shared.AppActionsServiceGenerateReportResponse
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAppV1AppReportActionServiceGenerateReportResponse) GetAppActionsServiceGenerateReportResponse() *shared.AppActionsServiceGenerateReportResponse {
	if o == nil {
		return nil
	}
	return o.AppActionsServiceGenerateReportResponse
}
