// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PolicyRef - The PolicyRef message.
type PolicyRef struct {
	// The id field.
	ID *string `json:"id,omitempty"`
}

func (o *PolicyRef) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
