// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/v2/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesListAttributeValuesRequest struct {
	AttributeTypeID string  `pathParam:"style=simple,explode=false,name=attribute_type_id"`
	PageSize        *int    `queryParam:"style=form,explode=true,name=page_size"`
	PageToken       *string `queryParam:"style=form,explode=true,name=page_token"`
}

func (o *C1APIAttributeV1AttributesListAttributeValuesRequest) GetAttributeTypeID() string {
	if o == nil {
		return ""
	}
	return o.AttributeTypeID
}

func (o *C1APIAttributeV1AttributesListAttributeValuesRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *C1APIAttributeV1AttributesListAttributeValuesRequest) GetPageToken() *string {
	if o == nil {
		return nil
	}
	return o.PageToken
}

type C1APIAttributeV1AttributesListAttributeValuesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// ListAttributeValuesResponse is the response for listing attribute values for a given AttributeType.
	ListAttributeValuesResponse *shared.ListAttributeValuesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *C1APIAttributeV1AttributesListAttributeValuesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIAttributeV1AttributesListAttributeValuesResponse) GetListAttributeValuesResponse() *shared.ListAttributeValuesResponse {
	if o == nil {
		return nil
	}
	return o.ListAttributeValuesResponse
}

func (o *C1APIAttributeV1AttributesListAttributeValuesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIAttributeV1AttributesListAttributeValuesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
