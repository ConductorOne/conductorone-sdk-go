// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesDeleteAttributeValueRequest struct {
	ID                          string                              `pathParam:"style=simple,explode=false,name=id"`
	DeleteAttributeValueRequest *shared.DeleteAttributeValueRequest `request:"mediaType=application/json"`
}

func (o *C1APIAttributeV1AttributesDeleteAttributeValueRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *C1APIAttributeV1AttributesDeleteAttributeValueRequest) GetDeleteAttributeValueRequest() *shared.DeleteAttributeValueRequest {
	if o == nil {
		return nil
	}
	return o.DeleteAttributeValueRequest
}

type C1APIAttributeV1AttributesDeleteAttributeValueResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// DeleteAttributeValueResponse is the empty response for deleting an attribute value.
	DeleteAttributeValueResponse *shared.DeleteAttributeValueResponse
}

func (o *C1APIAttributeV1AttributesDeleteAttributeValueResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAttributeV1AttributesDeleteAttributeValueResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAttributeV1AttributesDeleteAttributeValueResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIAttributeV1AttributesDeleteAttributeValueResponse) GetDeleteAttributeValueResponse() *shared.DeleteAttributeValueResponse {
	if o == nil {
		return nil
	}
	return o.DeleteAttributeValueResponse
}
