// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APIIamV1RolesUpdateRequest struct {
	UpdateRoleRequest *shared.UpdateRoleRequest `request:"mediaType=application/json"`
	RoleID            string                    `pathParam:"style=simple,explode=false,name=role_id"`
}

func (o *C1APIIamV1RolesUpdateRequest) GetUpdateRoleRequest() *shared.UpdateRoleRequest {
	if o == nil {
		return nil
	}
	return o.UpdateRoleRequest
}

func (o *C1APIIamV1RolesUpdateRequest) GetRoleID() string {
	if o == nil {
		return ""
	}
	return o.RoleID
}

type C1APIIamV1RolesUpdateResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	UpdateRolesResponse *shared.UpdateRolesResponse
}

func (o *C1APIIamV1RolesUpdateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *C1APIIamV1RolesUpdateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *C1APIIamV1RolesUpdateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *C1APIIamV1RolesUpdateResponse) GetUpdateRolesResponse() *shared.UpdateRolesResponse {
	if o == nil {
		return nil
	}
	return o.UpdateRolesResponse
}