// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APISystemlogV1ExportsSearchServiceSearchResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successful response
	ExportsSearchServiceSearchResponse *shared.ExportsSearchServiceSearchResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APISystemlogV1ExportsSearchServiceSearchResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APISystemlogV1ExportsSearchServiceSearchResponse) GetExportsSearchServiceSearchResponse() *shared.ExportsSearchServiceSearchResponse {
	if o == nil {
		return nil
	}
	return o.ExportsSearchServiceSearchResponse
}

func (o *C1APISystemlogV1ExportsSearchServiceSearchResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APISystemlogV1ExportsSearchServiceSearchResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}