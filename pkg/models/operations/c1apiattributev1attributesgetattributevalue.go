// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesGetAttributeValueRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *C1APIAttributeV1AttributesGetAttributeValueRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type C1APIAttributeV1AttributesGetAttributeValueResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// GetAttributeValueResponse is the response for getting an attribute value by id.
	GetAttributeValueResponse *shared.GetAttributeValueResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAttributeV1AttributesGetAttributeValueResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAttributeV1AttributesGetAttributeValueResponse) GetGetAttributeValueResponse() *shared.GetAttributeValueResponse {
	if o == nil {
		return nil
	}
	return o.GetAttributeValueResponse
}

func (o *C1APIAttributeV1AttributesGetAttributeValueResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAttributeV1AttributesGetAttributeValueResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
