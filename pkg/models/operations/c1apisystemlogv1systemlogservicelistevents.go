// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APISystemlogV1SystemLogServiceListEventsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful response
	SystemLogServiceListEventsResponse *shared.SystemLogServiceListEventsResponse
}

func (o *C1APISystemlogV1SystemLogServiceListEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APISystemlogV1SystemLogServiceListEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APISystemlogV1SystemLogServiceListEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APISystemlogV1SystemLogServiceListEventsResponse) GetSystemLogServiceListEventsResponse() *shared.SystemLogServiceListEventsResponse {
	if o == nil {
		return nil
	}
	return o.SystemLogServiceListEventsResponse
}