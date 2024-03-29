// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UserRef - A reference to a user.
type UserRef struct {
	// The id of the user.
	ID *string `json:"id,omitempty"`
}

func (o *UserRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
