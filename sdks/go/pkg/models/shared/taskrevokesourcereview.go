// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// TaskRevokeSourceReview - The TaskRevokeSourceReview message tracks which access review was the source of the specificed revoke ticket.
type TaskRevokeSourceReview struct {
	// The ID of the access review associated with the revoke task.
	AccessReviewID *string `json:"accessReviewId,omitempty"`
	// The ID of the certify ticket that was denied and created this revoke task.
	CertTicketID *string `json:"certTicketId,omitempty"`
}