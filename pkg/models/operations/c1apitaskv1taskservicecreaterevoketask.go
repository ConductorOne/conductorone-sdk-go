// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskServiceCreateRevokeTaskResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	TaskServiceCreateRevokeResponse *shared.TaskServiceCreateRevokeResponse
}

func (o *C1APITaskV1TaskServiceCreateRevokeTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskServiceCreateRevokeTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskServiceCreateRevokeTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskServiceCreateRevokeTaskResponse) GetTaskServiceCreateRevokeResponse() *shared.TaskServiceCreateRevokeResponse {
	if o == nil {
		return nil
	}
	return o.TaskServiceCreateRevokeResponse
}
