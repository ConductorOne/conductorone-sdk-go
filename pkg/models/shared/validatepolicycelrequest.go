// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ValidatePolicyCELRequest message.
type ValidatePolicyCELRequest struct {
	// The cel field.
	Cel *string `json:"cel,omitempty"`
}

func (o *ValidatePolicyCELRequest) GetCel() *string {
	if o == nil {
		return nil
	}
	return o.Cel
}
