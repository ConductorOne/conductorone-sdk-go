// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIAttributeV1AttributesCreateAttributeValueResponse struct {
	ContentType string
	// CreateAttributeValueResponse is the response for creating an attribute value.
	CreateAttributeValueResponse *shared.CreateAttributeValueResponse
	StatusCode                   int
	RawResponse                  *http.Response
}