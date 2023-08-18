// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskSearchServiceSearchResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The TaskSearchResponse message contains a list of results and a nextPageToken if applicable.
	TaskSearchResponse *shared.TaskSearchResponse
}

func (o *C1APITaskV1TaskSearchServiceSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APITaskV1TaskSearchServiceSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APITaskV1TaskSearchServiceSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APITaskV1TaskSearchServiceSearchResponse) GetTaskSearchResponse() *shared.TaskSearchResponse {
	if o == nil {
		return nil
	}
	return o.TaskSearchResponse
}
