// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIPolicyV1PoliciesCreateResponse struct {
	ContentType string
	// Successful response
	Policy      *shared.Policy
	StatusCode  int
	RawResponse *http.Response
}
