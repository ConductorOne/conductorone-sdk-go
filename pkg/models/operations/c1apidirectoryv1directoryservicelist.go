// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIDirectoryV1DirectoryServiceListResponse struct {
	ContentType string
	// Successful response
	DirectoryServiceListResponse *shared.DirectoryServiceListResponse
	StatusCode                   int
	RawResponse                  *http.Response
}