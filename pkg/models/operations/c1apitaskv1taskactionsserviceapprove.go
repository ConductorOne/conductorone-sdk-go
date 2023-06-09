// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ConductorOne/conductorone-sdk-go/pkg/models/shared"
	"net/http"
)

type C1APITaskV1TaskActionsServiceApproveRequest struct {
	C1APITaskV1TaskActionsServiceApproveRequestInput *shared.C1APITaskV1TaskActionsServiceApproveRequestInput `request:"mediaType=application/json"`
	TaskID                                           string                                                   `pathParam:"style=simple,explode=false,name=task_id"`
}

type C1APITaskV1TaskActionsServiceApproveResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful response
	C1APITaskV1TaskActionsServiceApproveResponse *shared.C1APITaskV1TaskActionsServiceApproveResponse
}
