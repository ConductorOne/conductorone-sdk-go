// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// Status - The application user status field.
type Status string

const (
	StatusStatusUnspecified Status = "STATUS_UNSPECIFIED"
	StatusStatusEnabled     Status = "STATUS_ENABLED"
	StatusStatusDisabled    Status = "STATUS_DISABLED"
	StatusStatusDeleted     Status = "STATUS_DELETED"
)

func (e Status) ToPointer() *Status {
	return &e
}

// AppUserStatus - The satus of the applicaiton user.
type AppUserStatus struct {
	// The details of applicaiton user status.
	Details *string `json:"details,omitempty"`
	// The application user status field.
	Status *Status `json:"status,omitempty"`
}

func (o *AppUserStatus) GetDetails() *string {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *AppUserStatus) GetStatus() *Status {
	if o == nil {
		return nil
	}
	return o.Status
}
